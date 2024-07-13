package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Milkado/api-challenge-jornada-milhas/database"
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

type Testimonies struct {
	Testimony string `json:"testimony" xml:"testimony" form:"testimony"`
	Name      string `json:"name" xml:"name" form:"name"`
	Picture   string `json:"picture" xml:"picture" form:"picture"`
}

func IndexTestimonies(c echo.Context) error {
	client := database.ConnectDB()
	ctx := context.Background()
	testimonies, err := client.Testimonies.Query().All(ctx)
	defer client.Close()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, testimonies)
}

func StoreTestimony(c echo.Context) error {

	client := database.ConnectDB()
	ctx := context.Background()
	t := new(Testimonies)
	if err := c.Bind(t); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	now := time.Now()
	newTestimony, err := client.Testimonies.Create().SetTestimony(t.Testimony).SetName(t.Name).SetPicture(t.Picture).SetCreatedAt(now).SetUpdatedAt(now).Save(ctx)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, newTestimony)
}
