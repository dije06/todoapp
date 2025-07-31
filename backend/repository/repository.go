package repository

import (
	"to-do-app/database"
	"to-do-app/models"
)

func GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	result := database.DB.Find(&todos)
	return todos, result.Error
}

func AddTodo(title string) (models.Todo, error) {
	todo := models.Todo{
		Title: title,
	}
	result := database.DB.Create(&todo)
	return todo, result.Error
}

func DeleteTodo(id int) error {
	todo := models.Todo{
		ID: id,
	}
	result := database.DB.Delete(todo)
	return result.Error
}
