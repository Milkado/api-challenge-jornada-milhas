package controllers

import (
	"context"
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

// type Testimonies struct {
// 	ID        uint   `json:"id"`
// 	Testimony string `json:"testmony"`
// 	Name      string `json:"name"`
// 	Picture   string `josn:"picture"`
// 	CreatedAt string `json:"created_at"`
// }
//

type (
	Testimonies struct {
		Testimony string `json:"testimony" xml:"testimony" form:"testimony" validate:"required,min=1"`
		Name      string `json:"name" xml:"name" form:"name" validate:"required,min=1"`
		Picture   string `json:"picture" xml:"picture" form:"picture" validate:"required,base64"`
	}
)

var ctx = context.Background()

func IndexTestimonies(c echo.Context) error {
	client := database.ConnectDB()
	testimonies, err := client.Testimonies.Query().All(ctx)
	defer client.Close()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, testimonies)
}

func StoreTestimony(c echo.Context) error {

	client := database.ConnectDB()

	if err := helpers.WithTx(ctx, client, func(tx *ent.Tx) error {
		t := new(Testimonies)
		if err := c.Bind(t); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		if err := helpers.Validate(t, c); err != nil {
			return err
		}

		picture := uuid.New().String() + time.Now().Format("2006-01-02")

		newTestimony, err := tx.Testimonies.Create().SetTestimony(t.Testimony).SetName(t.Name).SetPicture(picture).Save(ctx)
		if err != nil {
			defer client.Close()
			return c.String(http.StatusBadRequest, err.Error())
		}

		defer client.Close()
		return c.JSON(http.StatusCreated, newTestimony)
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	defer client.Close()
	return nil
}

func ShowTestimony(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	client := database.ConnectDB()
	testimony, err := client.Testimonies.Query().Where(testimonies.ID(id)).Only(ctx)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, testimony)
}
