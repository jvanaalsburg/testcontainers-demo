package handlers

import (
	"net/http"

	"demo-api/repos"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h Handler) GetAllUsers(c echo.Context) error {
	repo, err := repos.NewRepo(h.connStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	users, err := repo.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h Handler) GetUser(c echo.Context) error {
	repo, err := repos.NewRepo(h.connStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user, err := repo.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
