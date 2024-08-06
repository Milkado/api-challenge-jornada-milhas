package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Milkado/api-challenge-jornada-milhas/database"
	"github.com/Milkado/api-challenge-jornada-milhas/ent"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/destinies"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	Destinies struct {
		Name    string  `json:"name" xml:"name" form:"name" validate:"required,min=1"`
		Price   float64 `json:"price" xml:"price" form:"price" validate:"required,numeric"`
		Picture string  `json:"picture" xml:"picture" form:"picture" validate:"required"`
	}
	DestiniesUpdate struct {
		Name    string  `json:"name" xml:"name" form:"name" validate:"required,min=1"`
		Price   float64 `json:"price" xml:"price" form:"price" validate:"required,numeric"`
	}
	PictureUpdate struct {
		Picture string `json:"picture" xml:"picture" form:"picture" validate:"required"`
	}
)

func IndexDestinies(c echo.Context) error {
	client := database.ConnectDB()
	destinies, err := client.Destinies.Query().All(ctx)
	defer client.Close()
	if err != nil {
		c.JSON(http.StatusAccepted, err)
	}

	return c.JSON(http.StatusOK, destinies)
}

func ShowDestinies(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	client := database.ConnectDB()
	destiny, sErr := client.Destinies.Query().Where(destinies.ID(id)).Only(ctx)

	if sErr != nil {
		if ent.IsNotFound(sErr) {
			return c.JSON(http.StatusNotFound, sErr.Error())
		}
		c.JSON(http.StatusBadRequest, sErr.Error())
	}

	return c.JSON(http.StatusOK, destiny)
}

func StoreDestinies(c echo.Context) error {
	client := database.ConnectDB()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		d := new(Destinies)
		if bindErr := c.Bind(d); bindErr != nil {
			return c.JSON(http.StatusBadRequest, bindErr)
		}
		if validateErr := helpers.Validate(d, c); validateErr != nil {
			return c.JSON(http.StatusBadRequest, validateErr)
		}

		picture := uuid.New().String() + time.Now().Format("2006-01-02")

		newDestiny, storeErr := client.Destinies.Create().SetName(d.Name).SetPicture(picture).SetPrice(d.Price).Save(ctx)

		if storeErr != nil {
			return c.JSON(http.StatusBadRequest, storeErr)
		}

		return c.JSON(http.StatusCreated, newDestiny)
	}); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	defer client.Close()
	return nil
}

func ChangeDestiny(c echo.Context) error {
	client := database.ConnectDB()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		id := idToInt(c, c.Param("id"))
		d := new(DestiniesUpdate)
		if err := c.Bind(d); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		
		destiny ,err := client.Destinies.UpdateOneID(id).SetName(d.Name).SetPrice(d.Price).Save(ctx)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, destiny)
	}); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	defer client.Close()
	return nil
}

func DeleteDestiny(c echo.Context) error {
	client := database.ConnectDB()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		deleteErr := client.Destinies.DeleteOneID(id).Exec(ctx)
		if deleteErr != nil {
			return c.JSON(http.StatusBadRequest, deleteErr)
		}

		return c.JSON(http.StatusOK, "Destiny deleted")
	}); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	defer client.Close()
	return nil
}
