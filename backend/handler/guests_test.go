package handler

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GuestsModelStub struct {}

func (g *GuestsModelStub) FindByID(id string) models.Guest {
	return models.Guest{
		Title: "sample title",
		Message: "sample message",
	}
}

func (g *GuestsModelStub) FindAll(offset int) []models.Guest {
	guests := []models.Guest{}
	guests = append(guests, models.Guest{
		ID: 100,
		Title: "sample title",
		Message: "sample message",
	})
	return guests
}

func (g *GuestsModelStub) CreateGuest(title string, message string) bool {
	newGuest := models.Guest{Title: title, Message: message}
	if newGuest.Title != title {
		return false
	}
	return true
}

func TestGetDetail(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/guest/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	g := &GuestsModelStub{}
	h := NewHandler(g)

	var guestJSON = models.Guest{Title: "sample title", Message: "sample message"}

	if assert.NoError(t, h.GetDetail(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body, err := ioutil.ReadAll(rec.Body)
		if err != nil {
			fmt.Println(err)
		}
		actual := models.Guest{}
		err = json.Unmarshal(body, &actual)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equal(t, guestJSON.Title, actual.Title)
		assert.Equal(t, guestJSON.Message, actual.Message)
	}
}

func TestGetIndex(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	g := &GuestsModelStub{}
	h := NewHandler(g)

	var guestJSON = []models.Guest{models.Guest{ID: 1, Title: "sample title", Message: "sample message"}}

	if assert.NoError(t, h.GetDetail(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body, err := ioutil.ReadAll(rec.Body)
		if err != nil {
			fmt.Println(err)
		}
		actual := models.Guest{}
		err = json.Unmarshal(body, &actual)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equal(t, guestJSON[0].ID, actual.ID)
		assert.Equal(t, guestJSON[0].Title, actual.Title)
		assert.Equal(t, guestJSON[0].Message, actual.Message)
	}
}

func TestPostCreate(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/create")

	g := &GuestsModelStub{}
	h := NewHandler(g)

	var guestJSON = models.Guest{ID: 1, Title: "sample title", Message: "sample message"}

	if assert.NoError(t, h.GetDetail(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body, err := ioutil.ReadAll(rec.Body)
		if err != nil {
			fmt.Println(err)
		}
		actual := models.Guest{}
		err = json.Unmarshal(body, &actual)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equal(t, guestJSON.ID, actual.ID)
		assert.Equal(t, guestJSON.Title, actual.Title)
		assert.Equal(t, guestJSON.Message, actual.Message)
	}
}
