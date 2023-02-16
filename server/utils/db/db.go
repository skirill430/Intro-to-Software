package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"username"`
	Password string `json:"password"`
	List     string `json:"list"`
}

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to database.")
	}

	DB = db
}
