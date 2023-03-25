package utils

type User struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
}

type UserProduct struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Username    string `json:"username"`
	SellerName  string `json:"seller_name"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	Rating      string `json:"rating"`
	ImageURL    string `json:"image_url"`
}

type Product struct {
	ID          string `json:"id"`
	SellerName  string `json:"seller_name"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	Rating      string `json:"rating"`
	ImageURL    string `json:"image_url"`
}
