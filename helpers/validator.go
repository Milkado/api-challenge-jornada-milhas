package helpers

import (
	"encoding/base64"
	"net/http"
	"reflect"

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

func base64Validation(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() != reflect.Slice && field.Kind() != reflect.Array {
		return false
	}

	for i := 0; i < field.Len(); i++ {
		element := field.Index(i)
		if element.Kind() != reflect.String {
			return false
		}
		str := element.String()
		if _, err := base64.StdEncoding.DecodeString(str); err != nil {
			return false
		}
	}
	return true
}

func Validate(i interface{}, c echo.Context) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("base64_array", base64Validation)
	c.Echo().Validator = &CustomValidator{validator: validate}

	if err := c.Validate(i); err != nil {
		return err
	}

	return nil
}
