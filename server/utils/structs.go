package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("secret_key")

type User struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
}

type Cookies struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type UserProduct struct {
	Username    string `json:"username"`
	SellerName  string `json:"seller_name"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	Rating      string `json:"rating"`
	ImageURL    string `json:"image_url"`
}
