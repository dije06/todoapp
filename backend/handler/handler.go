package handler

import (
	"net/http"
	"strconv"
	"to-do-app/models"
	"to-do-app/repository"

	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) error {
	todos, err := repository.GetTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, todos)
}

func AddTodo(c echo.Context) error {
	var title models.Todo
	if err := c.Bind(&title); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	created, err := repository.AddTodo(title.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, created)
}

func DeleteTodo(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = repository.DeleteTodo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
