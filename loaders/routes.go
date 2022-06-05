package loaders

import (
	"todo/modules/Todo/TodoRoutes"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(route *gin.Engine) *gin.Engine {
	v1 := route.Group("/v1")
	TodoRoutes.Routes(v1)
	return route
}
