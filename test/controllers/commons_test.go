package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"github.com/labstack/echo/v4"
)

var (
	loginJson = `{
		"email":` + helpers.Env("USER_FOR_TEST") + `,
		"password":` + helpers.Env("PASSWORD_FOR_TEST") + `
	}`
)

func GetAuthToken() string {
	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(loginJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	type Bearer struct {
		Token string `json:"token"`
	}

	var response Bearer
	err := json.Unmarshal(rec.Body.Bytes(), &response) // test later

	if err != nil {
		fmt.Println(helpers.Red + err.Error() + helpers.Reset)
	}

	return "Bearer " + response.Token
}

func ExtractCreatedId(rec *httptest.ResponseRecorder) string {
	type Id struct {
		ID int64 `json:"id"`
	}

	var response Id
	err := json.Unmarshal(rec.Body.Bytes(), &response)

	if err != nil {
		fmt.Println(helpers.Red + err.Error() + helpers.Reset)
		return ""
	}

	return fmt.Sprintf("%d", response.ID)
}
