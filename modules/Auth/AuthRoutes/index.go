package AuthRoutes

import (
	"todo/modules/Auth/AuthHandlers.go"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	{
		r.POST("signin", AuthHandlers.Signin)
		r.POST("signup", AuthHandlers.Signup)
	}
}
