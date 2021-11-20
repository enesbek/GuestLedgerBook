package main

import (
	"backend/db"
	"backend/handler"
	"backend/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	d := db.DBConnect()
	h := handler.NewHandler(models.NewGuestModel(d))

	e.GET("/:offset", h.GetIndex)
	e.GET("/guest/:id", h.GetDetail)
	e.POST("/create", h.PostCreate)

	e.Logger.Fatal(e.Start(":3000"))
}
