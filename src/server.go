package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const API_VERSION = "0.1.0"

func getVersion(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"version": API_VERSION})
}

func main() {
	e := echo.New()
	e.GET("/version", getVersion)
	e.Logger.Fatal(e.Start(":1323"))
}
