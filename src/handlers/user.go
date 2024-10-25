package handlers

import (
	"net/http"

	"demo-api/models"
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

func (h Handler) CreateUser(c echo.Context) error {
	repo, err := repos.NewRepo(h.connStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := models.User{}
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	newUser, err := repo.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newUser)
}

func (h Handler) UpdateUser(c echo.Context) error {
	repo, err := repos.NewRepo(h.connStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := models.User{}
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user.Id = id

	newUser, err := repo.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newUser)
}

func (h Handler) DeleteUser(c echo.Context) error {
	repo, err := repos.NewRepo(h.connStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = repo.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
