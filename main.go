package main

import (
	"log"
	"mvcsaga/configs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	config, err := configs.LoadConfig(".")

	// handle errors when read environment
	if err != nil {
		log.Fatalf("can't load environment app.env: %v", err)
	}

	e := echo.New()

	e.GET("/config", func(c echo.Context) error {
		data := "Hello from index page"
		return c.JSON(http.StatusOK, data)
	})

	// Start the server
	if err := e.Start(":" + config.AppPort); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}
