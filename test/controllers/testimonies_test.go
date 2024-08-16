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
	testimonyUpdate = `{
		"name": "Siri",
		"testimony": "I'm a little teapot"
	}`
)

// var token = GetAuthToken()
var id string

func TestCreateTestimony(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/depoimentos", strings.NewReader(testimonyJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.StoreTestimony(c)) {
		if assert.Equal(t, http.StatusCreated, rec.Code) {
			id = ExtractCreatedId(rec)
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
	c.SetParamValues(id)

	type TestimonyResponse struct {
		ID int64 `json:"id"`
	}

	if assert.NoError(t, controllers.ShowTestimony(c)) {
		if assert.Equal(t, http.StatusOK, rec.Code) {
			var response TestimonyResponse
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			idConv := fmt.Sprintf("%d", response.ID)
			if err != nil {
				fmt.Println(helpers.Red + err.Error() + helpers.Reset)
			}
			if assert.Equal(t, id, idConv) {
				fmt.Println(helpers.Green + "TestShowTestimony passed" + helpers.Reset)
			}
		}
	}
}

func TestUpdateTestimony(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/api/depoimentos/:id", strings.NewReader(testimonyUpdate))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	type TestimonyResponse struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		Testimony string `json:"testimony"`
	}

	type TestimonyUpdate struct {
		Name      string `json:"name"`
		Testimony string `json:"testimony"`
	}

	expected := TestimonyUpdate{
		Name:      "Siri",
		Testimony: "I'm a little teapot",
	}

	if assert.NoError(t, controllers.ChangeTestimony(c)) {
		if assert.Equal(t, http.StatusOK, rec.Code) {
			var response TestimonyResponse
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			idConv := fmt.Sprintf("%d", response.ID)
			if err != nil {
				fmt.Println(helpers.Red + err.Error() + helpers.Reset)
			}
			if assert.Equal(t, id, idConv) && assert.Equal(t, expected.Name, response.Name) && assert.Equal(t, expected.Testimony, response.Testimony) {
				fmt.Println(helpers.Green + "TestUpdateTestimony passed" + helpers.Reset)
			}
		}
	}
}

func TestDeleteTestimony(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/depoimentos/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, controllers.DeleteTestimony(c)) {
		var response string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		if err != nil {
			fmt.Println(helpers.Red + err.Error() + helpers.Reset)
		}
		if assert.Equal(t, http.StatusOK, rec.Code) && assert.Equal(t, "Deleted", response) {
			fmt.Println(helpers.Green + "TestDeleteTestimony passed" + helpers.Reset)
		}
	}
}
