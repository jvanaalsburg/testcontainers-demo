package main

import (
	"fmt"
	"net/http"

	"demo-api/handlers"

	"github.com/labstack/echo/v4"
)

const API_VERSION = "0.1.0"

func getVersion(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"version": API_VERSION})
}

func initHandler() *handlers.Handler {
	db_host := "db"
	db_port := 5432
	db_user := "postgres"
	db_pass := "postgres"
	db_name := "demo"

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		db_user,
		db_pass,
		db_host,
		db_port,
		db_name,
	)

	return handlers.NewHandler(connStr)
}

func main() {
	e := echo.New()
	e.GET("/version", getVersion)

	handler := initHandler()

	e.GET("/users", handler.GetAllUsers)
	e.GET("/users/:id", handler.GetUser)

	e.Logger.Fatal(e.Start(":1323"))
}
