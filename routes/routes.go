package routes

import (
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
	auth.PATCH("/depoimentos/:id", controllers.ChangeTestimony)  
	auth.DELETE("/depoimentos/:id", controllers.DeleteTestimony)
	open.POST("/depoimentos/:id/store-picture", controllers.StoreTestimonyPicture)
	
	//User
	open.POST("/request-new-password", controllers.RequestNewPassword)
	open.PATCH("/create-password/:token", controllers.StorePassword)

	//Destinies
	open.GET("/destinies", controllers.IndexDestinies)
	open.GET("/destinies/:id", controllers.ShowDestinies)    //TODO: Function and test
	auth.POST("/destinies", controllers.StoreDestinies)      //TODO: Function and test
	auth.PATCH("/destinies/:id", controllers.ChangeDestiny)  //TODO: Function and test
	auth.DELETE("/destinies/:id", controllers.DeleteDestiny) //TODO: Function and test
	auth.POST("/destinies/:id/store-picture", controllers.StoreDestinyPicture)
	auth.POST("/destinies/:id/store-many-pictures", controllers.StoreManyDestinyPictures)

}
