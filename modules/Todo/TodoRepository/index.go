package TodoRepository

import (
	"todo/models"
	"todo/models/db"
)

//fetch all todos at once
func GetAllTodos(todo *[]models.Todo) (err error) {
	if err = db.DB.Find(&todo).Error; err != nil {
		return err
	}
	return nil
}

//insert a todo
func CreateATodo(todo *models.Todo) (err error) {
	if err = db.DB.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

//fetch one todo
func GetATodo(todo *models.Todo, id string) (err error) {
	if err := db.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return err
	}
	return nil
}

//update a todo
func UpdateATodo(todo *models.Todo, id string) (err error) {
	db.DB.Save(&todo)
	return nil
}

//delete a todo
func DeleteATodo(todo *models.Todo, id string) (err error) {
	db.DB.Where("id = ?", id).Delete(&todo)
	return nil
}
