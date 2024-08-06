package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Milkado/api-challenge-jornada-milhas/ent"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
