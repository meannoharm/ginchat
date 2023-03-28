package router

import (
	"ginchat/services"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", services.GetIndex)
	r.GET("/user/getUserList", services.GetUserList)

	return r
}
