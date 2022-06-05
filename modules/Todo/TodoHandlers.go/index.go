package TodoHandlers

import (
	"log"
	"net/http"
	"todo/models"
	"todo/modules/Todo/TodoService"

	"github.com/gin-gonic/gin"
)

//List all todos
func GetTodos(c *gin.Context) {
	var todo []models.Todo
	err := TodoService.GetTodos(todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
	log.Println(todo)
}

//Create a Todo
func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := TodoService.CreateATodo(todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Get a particular Todo with id
func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	err := TodoService.GetATodo(id, todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Update an existing Todo
func UpdateATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	c.BindJSON(&todo)
	err := TodoService.UpdateATodo(id, todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Delete a Todo
func DeleteATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := TodoService.DeleteATodo(id, todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
