package services

import "github.com/gin-gonic/gin"

// GetIndex
// @Tags Index
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome !!",
	})
}
