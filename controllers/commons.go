package controllers

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Milkado/api-challenge-jornada-milhas/database"
	"github.com/Milkado/api-challenge-jornada-milhas/ent"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type (
	PictureUpdate struct {
		Picture string `json:"picture" xml:"picture" form:"picture" validate:"required"`
	}
)

var ctx = context.Background()

func withTx(c context.Context, client ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(c)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			if helpers.Env("DEBUG") == "false" {
				panic(v)
			}
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}

		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

func hashPass(pass string) string {
	hashed, hashErr := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if hashErr != nil {
		return hashErr.Error()
	}
	return string(hashed)
}

func checkHash(pass string, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	if err != nil {
		return nil
	}

	return err
}

func idToInt(c echo.Context, idParam string) int {
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return id
}

func checkIfDestinyExists(c echo.Context, id int) bool {
	client := database.ConnectDB()
	defer client.Close()

	_, err := client.Destinies.Get(ctx, id)
	return err == nil
}

func storeFormPicture(file *multipart.FileHeader, c echo.Context, path string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return err.Error(), err
	}
	defer src.Close()
	fileExt := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + time.Now().Format("2006-01-02") + fileExt
	fullPath := filepath.Join(".", path, fileName)

	dst, err := os.Create(fullPath)
	if err != nil {
		return "create error", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "copy error", err
	}

	return fileName, nil
}

func storeBase64Picture(base64File string, c echo.Context, path string) (string, error) {
	// Decode the base64 data
	data, decodeErr := base64.StdEncoding.DecodeString(base64File)
	if decodeErr != nil {
		return "", fmt.Errorf("base64 decode error: %v", decodeErr)
	}

	// Detect the file type
	fileType, err := detectFileType(data)
	if err != nil {
		return "", fmt.Errorf("file type detection error: %v", err)
	}

	// Generate a unique file name with the detected file extension
	fileName := uuid.New().String() + time.Now().Format("2006-01-02") + "." + fileType
	fullPath := filepath.Join(".", path, fileName)

	// Write the decoded file to disk
	if err := os.WriteFile(fullPath, data, 0644); err != nil {
		return "", fmt.Errorf("file write error: %v", err)
	}

	return fileName, nil
}

func detectFileType(data []byte) (string, error) {
	if len(data) < 4 {
		return "", fmt.Errorf("data too short to determine file type")
	}

	switch {
	case data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF:
		return "jpeg", nil
	case data[0] == 0x89 && data[1] == 'P' && data[2] == 'N' && data[3] == 'G':
		return "png", nil
	default:
		return "", fmt.Errorf("unknown file type or not supported")
	}
}
