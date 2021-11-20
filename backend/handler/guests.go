package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/models"
	"github.com/labstack/echo/v4"
)

type resultLists struct {
	Guests []models.Guest `json:"guests"`
}

type handler struct {
	GuestModel models.GuestModelImpl
}

type responseGuestModel struct {
	Title string `json:"title" db:"title"`
	Message string `json:"message" db:"message"`
}

func NewHandler(g models.GuestModelImpl) *handler {
	return &handler{g}
}

func (h *handler) GetIndex(c echo.Context) error {
	stringOffset := c.Param("offset")
	newOffset, err := strconv.Atoi(stringOffset)
	newOffset--
	lists := h.GuestModel.FindAll(newOffset)

	if err != nil {
		fmt.Println(err)
	}

	u := &resultLists{
		Guests: lists,
	}
	return c.JSON(http.StatusOK, u)
}

func (h *handler) GetDetail(c echo.Context) error {
	id := c.Param("id")
	u := h.GuestModel.FindByID(id)
	return c.JSON(http.StatusOK, u)
}

func (h *handler) PostCreate(c echo.Context) error{
	responseGuest := &responseGuestModel{}
	if err := c.Bind(responseGuest); err != nil {
		fmt.Println("Unmarshall")
		return err
	}

	h.GuestModel.CreateGuest(responseGuest.Title, responseGuest.Message)

	return c.JSON(http.StatusCreated, true)
}