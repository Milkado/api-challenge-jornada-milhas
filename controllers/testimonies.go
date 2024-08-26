package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Milkado/api-challenge-jornada-milhas/database"
	"github.com/Milkado/api-challenge-jornada-milhas/ent"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/testimonies"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	Testimonies struct {
		Testimony string `json:"testimony" xml:"testimony" form:"testimony" validate:"required,min=1"`
		Name      string `json:"name" xml:"name" form:"name" validate:"required,min=1"`
		Picture   string `json:"picture" xml:"picture" form:"picture" validate:"required"`
		DestinyID int    `json:"destiny_id" xml:"destiny_id" form:"destiny_id" validate:"required"`
	}
	TestimoniesUpdate struct {
		Testimony string `json:"testimony" xml:"testimony" form:"testimony" validate:"required,min=1"`
		Name      string `json:"name" xml:"name" form:"name" validate:"required,min=1"`
		DestinyID int    `json:"destiny_id" xml:"destiny_id" form:"destiny_id" validate:"required"`
	}
)

var (
	testimonyPath = "/public/pictures/testimonies/"
)

func IndexTestimonies(c echo.Context) error {
	client := database.ConnectDB()
	defer client.Close()

	testimonies, err := client.Testimonies.Query().WithDestinies().All(ctx)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, testimonies)
}

func ShowTestimony(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	client := database.ConnectDB()
	defer client.Close()
	testimony, err := client.Testimonies.Query().Where(testimonies.ID(id)).WithDestinies().Only(ctx)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, testimony)
}

func StoreTestimony(c echo.Context) error {

	client := database.ConnectDB()
	defer client.Close()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		t := new(Testimonies)
		if err := c.Bind(t); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err := helpers.Validate(t, c); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if !checkIfDestinyExists(c, t.DestinyID) {
			return c.JSON(http.StatusBadRequest, "Destiny does not exist")
		}

		picture := uuid.New().String() + time.Now().Format("2006-01-02")

		newTestimony, err := tx.Testimonies.Create().SetTestimony(t.Testimony).SetName(t.Name).SetPicture(picture).SetDestinyID(t.DestinyID).Save(ctx)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, newTestimony)
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func ChangeTestimony(c echo.Context) error {
	client := database.ConnectDB()
	defer client.Close()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		id, cErr := strconv.Atoi(c.Param("id"))
		if cErr != nil {
			return c.JSON(http.StatusInternalServerError, cErr)
		}
		t := new(TestimoniesUpdate)
		if err := c.Bind(t); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		testimony, uErr := client.Testimonies.UpdateOneID(id).SetName(t.Name).SetTestimony(t.Testimony).Save(ctx)
		if uErr != nil {
			return c.JSON(http.StatusBadRequest, uErr)
		}

		return c.JSON(http.StatusOK, testimony)
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func DeleteTestimony(c echo.Context) error {
	client := database.ConnectDB()
	defer client.Close()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		dErr := client.Testimonies.DeleteOneID(id).Exec(ctx)

		if dErr != nil {
			return c.JSON(http.StatusBadRequest, dErr.Error())
		}

		return c.JSON(http.StatusOK, "Deleted")
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

func StoreTestimonyPicture(c echo.Context) error {
	client := database.ConnectDB()
	defer client.Close()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		p := new(Picture)
		if bindErr := c.Bind(p); bindErr != nil {
			return c.JSON(http.StatusBadRequest, bindErr)
		}

		id := idToInt(c, c.Param("id"))

		var fileName string
		var err error
		if p.Picture != "" {
			if validateErr := helpers.Validate(p, c); validateErr != nil {
				return c.JSON(http.StatusBadRequest, validateErr)
			}

			base64File := p.Picture
			fileName, err = storeBase64Picture(base64File, c, testimonyPath)
			if err != nil {
				return c.JSON(http.StatusBadRequest, []string{err.Error(), fileName})
			}
		} else {
			file, fErr := c.FormFile("picture")
			if fErr != nil {
				return c.JSON(http.StatusBadRequest, []string{fErr.Error(), "no file"})
			}
			fileName, err = storeFormPicture(file, c, testimonyPath)
			if err != nil {
				return c.JSON(http.StatusBadRequest, []string{err.Error(), fileName})
			}
		}

		updateTestimony, err := client.Testimonies.UpdateOneID(id).SetPicture(fileName).Save(ctx)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, updateTestimony)
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil
}
