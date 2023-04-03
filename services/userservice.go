package services

import (
	"ginchat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary Get User List
// @Tags User
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary Create User
// @Tags User
// @Param name query string false "name"
// @Param password query string false "password"
// @Param rePassword query string false "rePassword"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	rePassword := c.Query("rePassword")
	if password != rePassword {
		c.JSON(200, gin.H{
			"message": "password not same",
		})
		return
	}
	user.Password = password
	models.CreateUser(user)

	c.JSON(200, gin.H{
		"message": "create success",
	})
}

// DeleteUser
// @Summary Delete User
// @Tags User
// @Param name query string false "id"
// @Success 200 {string} json{"code", "message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)

	c.JSON(200, gin.H{
		"message": "delete success",
	})
}

// UpdateUser
// @Summary Update User
// @Tags User
// @Param id formData string false "id"
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Success 200 {string} json{"code", "message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}

	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")

	models.UpdateUser(user)

	c.JSON(200, gin.H{
		"message": "update success",
	})
}
