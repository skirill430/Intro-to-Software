package db

import (
	"fmt"

	"github.com/skirill430/Quick-Shop/server/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
	List     string `json:"list"`
}

var DB *gorm.DB

func ConnectDB(db_name string) {
	path := fmt.Sprintf("%s.db", db_name)
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to database.")
	} else {
		fmt.Printf("Connected to local database: '%s'.\n", path)
	}

	db.AutoMigrate(&User{})

	// create example user only upon first time creating users.db
	ex_user := &User{
		Username: "example_user",
		Password: utils.HashAndSalt([]byte("123456")),
		List:     "item1, item2",
	}
	// adds example user if database doesn't contain it already
	db.Where("username = ?", ex_user.Username).FirstOrCreate(&ex_user)

	DB = db
}

func DeleteUser(username string) {
	res := DB.Where("username = ?", username).Delete(&User{})
	// what if user couldn't be found?
	if res.RowsAffected == 0 {
		fmt.Printf("Could not delete '%s'.\n", username)
	}
}

func ClearDB() {
	res := DB.Exec("DELETE FROM users")
	if res.RowsAffected == 0 {
		fmt.Println("Could not clear database.")
	}
}
