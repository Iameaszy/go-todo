package TodoService

import (
	"todo/models"
	"todo/modules/Todo/TodoRepository"
)

func GetTodos(todo *[]models.Todo) (err error) {
	return TodoRepository.GetAllTodos(todo)
}

func CreateATodo(todo *models.Todo) (err error) {
	return TodoRepository.CreateATodo(todo)
}

func GetATodo(id string, todo *models.Todo) (err error) {
	return TodoRepository.GetATodo(todo, id)
}

func UpdateATodo(id string, todo *models.Todo) (err error) {
	return TodoRepository.UpdateATodo(todo, id)
}

func DeleteATodo(id string, todo *models.Todo) (err error) {
	return TodoRepository.DeleteATodo(todo, id)
}
