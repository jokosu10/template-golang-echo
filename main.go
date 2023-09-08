package main

import (
	"log"
	"net/http"
	"template-golang-echo/configs"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := configs.LoadConfig(".")

	// handle errors when read environment
	if err != nil {
		log.Fatalf("can't load environment app.env: %v", err)
	}

	e := echo.New()

	e.GET("/testing", func(c echo.Context) error {
		data := "Hello from index page"
		return c.JSON(http.StatusOK, data)
	})

	e.GET("/configs", func(c echo.Context) error {
		data := config
		return c.JSON(http.StatusOK, data)
	})

	// Start the server
	if err := e.Start(":" + config.AppPort); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
