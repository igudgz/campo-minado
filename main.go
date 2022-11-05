package main

import (
	"net/http"

	"github.com/igudgz/campo-minado/delivery"
	"github.com/igudgz/campo-minado/repository"
	"github.com/igudgz/campo-minado/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	repository := repository.NewRepo()
	usecase := usecase.NewService(repository)
	handler := delivery.NewHTTPHandler(usecase)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/games/:id", handler.Get)
	e.POST("/games", handler.Create)
	e.PUT("/games/:id", handler.RevealCell)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
