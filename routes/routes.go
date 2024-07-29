package routes

import (
	"net/http"

	"github.com/Milkado/api-challenge-jornada-milhas/controllers"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HandleRequest() {
	r := echo.New()

	routes(r)

	r.Logger.Fatal(r.Start(helpers.Env("PORT")))
}

// TODO: if deployed, configure cors to accept only de front domain in .env file
func routes(echo *echo.Echo) {
	auth := echo.Group("api")
	auth.Use(echojwt.JWT([]byte(helpers.Env("JWT_SECRET"))))
	auth.Use(middleware.CORS())
	auth.GET("/me", controllers.Me)

	open := echo.Group("api")
	open.Use(middleware.CORS())

	//Auth
	open.POST("/login", controllers.Login)

	//Depoimentos
	open.GET("/depoimentos", controllers.IndexTestimonies)
	open.GET("/depoimentos/:id", controllers.ShowTestimony)
	open.POST("/depoimentos", controllers.StoreTestimony)
	open.PATCH("/depoimentos/:id", placeHolderEcho)  //Change to authenticated
	open.DELETE("/depoimentos/:id", placeHolderEcho) //Change to authenticated

	//User
	open.POST("/request-new-password", controllers.RequestNewPassword)
	open.PATCH("/create-password/:token", controllers.StorePassword)

}

func placeHolderEcho(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
