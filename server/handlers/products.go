package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/skirill430/Quick-Shop/server/utils"
	"gorm.io/gorm"
)

func SaveProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product utils.UserProduct
	json.NewDecoder(r.Body).Decode(&product)

	username, status := utils.ExtractUsernameFromCookie(w, r)
	product.Username = username

	// if cookie unpacking failed, return appropriate error
	if status != 200 {
		w.WriteHeader(status)
		w.Write([]byte("Ensure that you are logged in and have a login cookie in your browser."))
		return
	}

	// fine if no image URL is provided
	if product.Username == "" || product.ProductName == "" || product.Price == "" || product.Rating == "" || product.SellerName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request Body Missing Fields"))
		return
	}

	// only allow users with accounts to save products
	var dbUser *utils.User
	err := utils.UsersDB.First(&dbUser, "username = ?", product.Username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("To gain permission to save a product, create an account with the given username."))
		return
	}

	// save product only if it hasn't already been
	res := utils.UserProductsDB.Where("product_name = ? AND seller_name = ? AND username = ?", product.ProductName, product.SellerName, product.Username).FirstOrCreate(&product)

	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Sorry, this product has already been saved to this user's list of products."))
		return
	}
}

func RemoveProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product utils.UserProduct
	json.NewDecoder(r.Body).Decode(&product)

	username, status := utils.ExtractUsernameFromCookie(w, r)
	product.Username = username

	// if cookie unpacking failed, return appropriate error
	if status != 200 {
		w.WriteHeader(status)
		w.Write([]byte("Ensure that you are logged in and have a login cookie in your browser."))
		return
	}

	// fine if no image URL is provided
	if product.Username == "" || product.ProductName == "" || product.Price == "" || product.Rating == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request Body Missing Fields"))
		return
	}

	// only allow users with accounts to delete products
	var dbUser *utils.User
	err := utils.UsersDB.First(&dbUser, "username = ?", product.Username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("To gain permission to delete a saved product, create an account with the given username."))
		return
	}

	// delete product only if ID and seller name match (IDs are specific to the seller i think)
	res := utils.UserProductsDB.Where("product_name = ? AND seller_name = ?", product.ProductName, product.SellerName).Delete(&product)

	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sorry, this product cannot be found in the user's saved products."))
		return
	}
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, status := utils.ExtractUsernameFromCookie(w, r)

	// if cookie unpacking failed, return appropriate error
	if status != 200 {
		w.WriteHeader(status)
		w.Write([]byte("Ensure that you are logged in and have a login cookie in your browser."))
		return
	}

	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request Body Missing Fields"))
		return
	}

	// check if username has account tied to it
	var dbUser *utils.User
	err := utils.UsersDB.First(&dbUser, "username = ?", username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("This username does not exist. Create an account."))
		return
	}

	var products []utils.UserProduct
	// return all fields except username (struct will omit returning username field as it is empty) to protect user data
	utils.UserProductsDB.Where("username = ?", username).Select("seller_name", "product_name", "price", "rating", "image_url").Find(&products)

	json.NewEncoder(w).Encode(products)
}
