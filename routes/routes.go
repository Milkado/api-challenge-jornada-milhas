package routes

import (
	"net/http"

	"github.com/Milkado/api-challenge-jornada-milhas/controllers"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"github.com/labstack/echo/v4"
)

func HandleRequest() {
	r := echo.New()

	routes(r)

	r.Logger.Fatal(r.Start(helpers.Env("PORT")))
}

func routes(echo *echo.Echo) {
	// auth := echo.Group("api")
	// auth.Use()

	open := echo.Group("api")

	//Depoimentos
	open.GET("/depoimentos", controllers.IndexTestimonies)
	echo.GET("/depoimentos/:id", placeHolderEcho)
	echo.POST("/depoimentos", controllers.StoreTestimony)
	echo.PATCH("/depoimentos/:id", placeHolderEcho)  //Change to authenticated
	echo.DELETE("/depoimentos/:id", placeHolderEcho) //Change to authenticated
}

func placeHolderEcho(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
