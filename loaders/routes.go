package loaders

import (
	"todo/middlewares"
	"todo/modules/Auth/AuthRoutes"
	"todo/modules/Todo/TodoRoutes"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(route *gin.Engine) *gin.Engine {
	v1 := route.Group("/v1")
	AuthRoutes.Routes(v1)
	v1.Use(middlewares.Auth())
	{
		TodoRoutes.Routes(v1)
	}
	return route
}
