package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
	List     string `json:"list"`
}

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to database.")
	} else {
		fmt.Println("Connected to local database: 'users.db'")
	}

	db.AutoMigrate(&User{})

	// create example user only upon first time creating users.db
	ex_user := &User{
		Username: "admin",
		Password: "123456",
		List:     "",
	}
	// adds example user if database doesn't contain it already
	db.Where("username = ?", ex_user.Username).FirstOrCreate(&ex_user)

	DB = db
}
