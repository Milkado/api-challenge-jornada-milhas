package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Milkado/api-challenge-jornada-milhas/controllers"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	testimonyJson = `{
	  "name": "Vivenna",
	  "testimony": "Haladren is strange. The gods creeps me out.",
	  "picture": "test"
	}`
	testimonyId = "1"
)

func TestCreateTestimony(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/depoimentos", strings.NewReader(testimonyJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.StoreTestimony(c)) {
		if assert.Equal(t, http.StatusCreated, rec.Code) {
			fmt.Println(helpers.Green + "TestCreateTestimony passed" + helpers.Reset)
		}
	}
}

func TestShowTestimony(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/depoimentos/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	type TestimonyResponse struct {
		ID string `json:"id"`
	}

	var response TestimonyResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response) // test later

	if assert.NoError(t, err) {
		if assert.NoError(t, controllers.ShowTestimony(c)) {
			if assert.Equal(t, http.StatusOK, rec.Code) {
				if assert.Equal(t, testimonyId, response.ID) {
					fmt.Println(helpers.Green + "TestShowTestimony passed" + helpers.Reset)
				}
			}
		}
	}
}
