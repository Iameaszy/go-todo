package TodoRoutes

import (
	"todo/modules/Todo/TodoHandlers.go"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	{
		r.GET("todo", TodoHandlers.GetTodos)
		r.POST("todo", TodoHandlers.CreateATodo)
		r.GET("todo/:id", TodoHandlers.GetATodo)
		r.PUT("todo/:id", TodoHandlers.UpdateATodo)
		r.DELETE("todo/:id", TodoHandlers.DeleteATodo)
	}
}
