package router

import (
	"ginchat/docs"
	"ginchat/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/index", services.GetIndex)
	r.GET("/user/getUserList", services.GetUserList)
	r.GET("/user/createUser", services.CreateUser)
	r.GET("/user/deleteUser", services.DeleteUser)
	r.POST("/user/updateUser", services.UpdateUser)

	return r
}
