package services

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}
