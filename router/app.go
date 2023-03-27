package router

import (
	"ginchat/services"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", services.GetIndex)

	return r
}
