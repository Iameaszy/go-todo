package main

import (
	"todo/loaders"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	loaders.Load(r)
	r.Run()
}
