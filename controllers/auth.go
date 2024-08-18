package controllers

import (
	"net/http"
	"time"

	"github.com/Milkado/api-challenge-jornada-milhas/database"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/users"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	Credentials struct {
		Email    string `json:"email" xml:"email" form:"email" valdiate:"required"`
		Password string `json:"password" xml:"password" form:"password" valdiate:"required"`
	}
	SignedUser struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
		jwt.RegisteredClaims
	}
	Token struct {
		Token string `json:"token"`
	}
)

func Login(c echo.Context) error {
	client := database.ConnectDB()
	credentials := new(Credentials)
	if err := c.Bind(credentials); err != nil {
		defer client.Close()
		return c.JSON(http.StatusInternalServerError, err)
	}

	userGet, uErr := client.Users.Query().Where(users.Email(credentials.Email)).Only(ctx)
	if uErr != nil {
		defer client.Close()
		return c.JSON(http.StatusBadRequest, uErr)
	}

	if hashError := checkHash(credentials.Password+*userGet.RandSecurity, *userGet.Password); hashError != nil {
		return c.JSON(http.StatusBadRequest, "Email or password is incorrect")
	}

	token, tErr := generateJwt(userGet.Email, userGet.Name, userGet.ID)
	if tErr != nil {
		return c.JSON(http.StatusInternalServerError, tErr)
	}

	return c.JSON(http.StatusOK, &Token{token})
}

func generateJwt(email string, name string, id int) (string, error) {
	secret := []byte(helpers.Env("JWT_SECRET"))

	claims := &SignedUser{
		Email: email,
		Name:  name,
		ID:    id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(10))),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		return "Oops", err
	}

	return token, nil
}
