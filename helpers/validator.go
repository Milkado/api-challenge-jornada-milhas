package helpers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func Validate(i interface{}, c echo.Context) error {
	validate := validator.New()
	c.Echo().Validator = &CustomValidator{validator: validate}

	if err := c.Validate(i); err != nil {
		return err
	}

	return nil
}
