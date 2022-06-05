package loaders

import (
	"todo/models/db"

	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {
	db.Start()
	LoadRoutes(r)
}
