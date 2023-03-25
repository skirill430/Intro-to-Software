package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/skirill430/Quick-Shop/server/utils/db"
	"gorm.io/gorm"
)

/*
	{
		"id": "1113706",
		"username": "user1",
		"seller_name": "Target",
		"product_name": "North Face Backpack",
		"price": "$120.00",
		"rating": "4.6",
		"image_url": ""
	}
*/
func SaveProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product db.ProductInfo
	json.NewDecoder(r.Body).Decode(&product)

	// fine if no image URL is provided
	if product.Username == "" || product.ProductName == "" || product.Price == "" || product.Rating == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request Body Missing Fields"))
		return
	}

	// only allow users with accounts to save products
	var dbUser *db.User
	err := db.UsersDB.First(&dbUser, "username = ?", product.Username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("This username does not exist. Before saving a product, create an account."))
		return
	}

	// save product only if it hasn't already been
	res := db.ProductsDB.Where("id = ? AND seller_name = ?", product.ID, product.SellerName).Create(&product)

	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Sorry, this product has already been saved to this user's list of products."))
		return
	}
}

func RemoveProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product db.ProductInfo
	json.NewDecoder(r.Body).Decode(&product)

	// fine if no image URL is provided
	if product.Username == "" || product.ProductName == "" || product.Price == "" || product.Rating == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request Body Missing Fields"))
		return
	}

	// only allow users with accounts to delete products
	var dbUser *db.User
	err := db.UsersDB.First(&dbUser, "username = ?", product.Username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("This username does not exist. Before saving a product, create an account."))
		return
	}

	// delete product only if ID and seller name match (IDs are specific to the seller i think)
	res := db.ProductsDB.Where("id = ? AND seller_name = ?", product.ID, product.SellerName).Delete(&product)

	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sorry, this product cannot be found in the user's saved products."))
		return
	}

}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request Body Missing Fields"))
		return
	}

	// check if username has account tied to it
	var dbUser *db.User
	err := db.UsersDB.First(&dbUser, "username = ?", username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("This username does not exist. Create an account."))
		return
	}

	var products []db.ProductInfo
	db.ProductsDB.Where("username = ?", username).Find(&products)

	json.NewEncoder(w).Encode(products)
}
