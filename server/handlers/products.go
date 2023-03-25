package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/skirill430/Quick-Shop/server/utils/db"
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

	// delete product only if ID and seller name match (IDs are specific to the seller i think)
	res := db.ProductsDB.Where("id = ? AND seller_name = ?", product.ID, product.SellerName).Delete(&product)

	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sorry, this product cannot be found in the user's saved products."))
		return
	}

}
