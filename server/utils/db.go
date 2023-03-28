package utils

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var UsersDB *gorm.DB
var ProductsDB *gorm.DB

func ConnectDB(db_name string) {
	path := fmt.Sprintf("%s.db", db_name)
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Printf("Could not connect to local database: '%s'.\n", path)
	}

	if db_name == "users" || db_name == "users_test" {
		db.AutoMigrate(&User{})

		// create example user only upon first time creating users.db
		ex_user := &User{
			Username: "example_user",
			Password: HashAndSalt([]byte("123456")),
		}
		db.Where("username = ?", ex_user.Username).FirstOrCreate(&ex_user)

		UsersDB = db

	} else if db_name == "products" {
		db.AutoMigrate(&UserProduct{})

		// create example product only upon first time creating products.db
		ex_product := &UserProduct{
			Username:    "example_user",
			SellerName:  "Target",
			ProductName: "2022 Apple MacBook Air Laptop with M2 chip",
			Price:       "$1,145.94",
			Rating:      "4.2",
			ImageURL:    "https://i5.walmartimages.com/asr/323a5b34-669e-4c8d-9a1f-c2ad73e3b15e.23b625e851179b54b0af5e7045347e79.jpeg?odnHeight=180&odnWidth=180&odnBg=FFFFFF",
		}

		db.Where("username = ?", ex_product.Username).FirstOrCreate(&ex_product)

		ProductsDB = db
	}
}

func DeleteUser(username string) {
	res := UsersDB.Where("username = ?", username).Delete(&User{})
	// what if user couldn't be found?
	if res.RowsAffected == 0 {
		fmt.Printf("Could not delete '%s' from users database.\n", username)
	}
}

func ClearUsersDB() {
	res := UsersDB.Exec("DELETE FROM users")
	if res.RowsAffected == 0 {
		fmt.Println("Could not clear users database.")
	}
}
