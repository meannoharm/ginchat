package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:12345678@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "test"
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1)) // find product with integer primary key

	// Update - update product's price to 200
	db.Model(user).Update("Password", "1234")

}
