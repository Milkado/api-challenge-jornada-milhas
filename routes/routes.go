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
	open.GET("/depoimentos/:id", controllers.ShowTestimony)
	open.POST("/depoimentos", controllers.StoreTestimony)
	open.PATCH("/depoimentos/:id", placeHolderEcho)  //Change to authenticated
	open.DELETE("/depoimentos/:id", placeHolderEcho) //Change to authenticated
}

func placeHolderEcho(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
