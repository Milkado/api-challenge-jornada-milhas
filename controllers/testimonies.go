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
	}
	TestimoniesUpdate struct {
		Testimony string `json:"testimony" xml:"testimony" form:"testimony" validate:"required,min=1"`
		Name      string `json:"name" xml:"name" form:"name" validate:"required,min=1"`
	}
)

func IndexTestimonies(c echo.Context) error {
	client := database.ConnectDB()
	testimonies, err := client.Testimonies.Query().All(ctx)
	defer client.Close()
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
	testimony, err := client.Testimonies.Query().Where(testimonies.ID(id)).Only(ctx)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, testimony)
}

func StoreTestimony(c echo.Context) error {

	client := database.ConnectDB()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		t := new(Testimonies)
		if err := c.Bind(t); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err := helpers.Validate(t, c); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		picture := uuid.New().String() + time.Now().Format("2006-01-02")

		newTestimony, err := tx.Testimonies.Create().SetTestimony(t.Testimony).SetName(t.Name).SetPicture(picture).Save(ctx)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, newTestimony)
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	defer client.Close()
	return nil
}

func ChangeTestimony(c echo.Context) error {
	client := database.ConnectDB()

	if err := withTx(ctx, client, func (tx *ent.Tx) error {
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

	defer client.Close()
	return nil
}

func DeleteTestimony(c echo.Context) error {
	client := database.ConnectDB()

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

	defer client.Close()
	return nil
}
