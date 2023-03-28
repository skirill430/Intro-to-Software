package test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/skirill430/Quick-Shop/server/utils"
	"github.com/stretchr/testify/assert"
)

/* PRODUCT ROUTE TESTS */
func TestSaveProduct_OK(t *testing.T) {
	test_user_product := []byte(`{"username": "example_user", "seller_name": "Target",
	"product_name": "North Face Backpack", "price": "$120.00", "rating": "4.6", "image_url": ""}`)
	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(test_user_product))

	response := httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	a := assert.New(t)
	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	// delete newly created UserProduct in between tests
	utils.DeleteUserProduct("example_user", "North Face Backpack", "Target")
}

func TestSaveProduct_UnknownUsername(t *testing.T) {
	// "guest" will be unrecognized since it is not saved to users database (does not have account attached to username)
	test_user_product := []byte(`{"username": "guest", "seller_name": "Target",
	"product_name": "North Face Backpack", "price": "$120.00", "rating": "4.6", "image_url": ""}`)
	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(test_user_product))

	response := httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	a := assert.New(t)
	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusUnauthorized, response.Code, "HTTP request status code error")
}

func TestSaveProduct_AlreadySavedProduct(t *testing.T) {
	// save product first time
	test_user_product := []byte(`{"username": "example_user", "seller_name": "Target",
	"product_name": "North Face Backpack", "price": "$120.00", "rating": "4.6", "image_url": ""}`)
	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(test_user_product))

	response := httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	// attempt to save product again
	req, _ = http.NewRequest("POST", "/api/products", bytes.NewBuffer(test_user_product))

	response = httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	a := assert.New(t)
	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusConflict, response.Code, "HTTP request status code error")
}

func TestRemoveProduct_OK(t *testing.T) {
	test_user_product := []byte(`{"username": "example_user", "seller_name": "Target",
	"product_name": "2022 Apple MacBook Air Laptop with M2 chip", "price": "$1,145.94", 
	"rating": "4.2", "image_url": "https://i5.walmartimages.com/asr/323a5b34-669e-4c8d-9a1f-c2ad73e3b15e.23b625e851179b54b0af5e7045347e79.jpeg?odnHeight=180&odnWidth=180&odnBg=FFFFFF"}`)
	req, _ := http.NewRequest("DELETE", "/api/products", bytes.NewBuffer(test_user_product))

	response := httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	a := assert.New(t)
	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")
}

func TestRemoveProduct_UnknownUsername(t *testing.T) {
	test_user_product := []byte(`{"username": "guest", "seller_name": "Target",
	"product_name": "2022 Apple MacBook Air Laptop with M2 chip", "price": "$1,145.94", 
	"rating": "4.2", "image_url": "https://i5.walmartimages.com/asr/323a5b34-669e-4c8d-9a1f-c2ad73e3b15e.23b625e851179b54b0af5e7045347e79.jpeg?odnHeight=180&odnWidth=180&odnBg=FFFFFF"}`)
	req, _ := http.NewRequest("DELETE", "/api/products", bytes.NewBuffer(test_user_product))

	response := httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	a := assert.New(t)
	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusUnauthorized, response.Code, "HTTP request status code error")
}

func TestRemoveProduct_ProductNotSaved(t *testing.T) {
	test_user_product := []byte(`{"username": "example_user", "seller_name": "Target",
	"product_name": "Vagabond Vol. 3", "price": "$19.99", "rating": "5.0", "image_url": ""}`)
	req, _ := http.NewRequest("DELETE", "/api/products", bytes.NewBuffer(test_user_product))

	response := httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	a := assert.New(t)
	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusNotFound, response.Code, "HTTP request status code error")
}

func TestGetAllProducts_OK(t *testing.T) {
	body_username := []byte(`{"username": "test"}`)
	// first must create account
	test_credentials := []byte(`{"username": "test", 
	"password": "test-password"}`)
	req, _ := http.NewRequest("POST", "/api/user/signup", bytes.NewBuffer(test_credentials))
	response := httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	test_user_product1 := []byte(`{"username":"test","seller_name":"Target","product_name":"Vagabond Vol. 3","price":"$19.99","rating":"5.0","image_url":""}`)
	req, _ = http.NewRequest("POST", "/api/products", bytes.NewBuffer(test_user_product1))
	response = httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	test_user_product2 := []byte(`{"username":"test","seller_name":"Target","product_name":"North Face Backpack","price":"$120.00","rating":"4.6","image_url":""}`)
	req, _ = http.NewRequest("POST", "/api/products", bytes.NewBuffer(test_user_product2))
	response = httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	// two products added to username 'test'. Does GET return the correct information?
	req, _ = http.NewRequest("GET", "/api/products", bytes.NewBuffer(body_username))

	response = httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	a := assert.New(t)
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	product_list := fmt.Sprintf("[%s,%s]\n", test_user_product1, test_user_product2)
	a.Equal(product_list, response.Body.String(), "GET does not return the first product correctly")

	utils.DeleteUserProduct("test", "North Face Backpack", "Target")
	utils.DeleteUserProduct("test", "Vagabond Vol. 3", "Target")
	utils.DeleteUser("test")
}

func TestGetAllProducts_UnknownUsername(t *testing.T) {
	body_username := []byte(`{"username": "guest"}`)
	req, _ := http.NewRequest("GET", "/api/products", bytes.NewBuffer(body_username))

	response := httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	a := assert.New(t)
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusNotFound, response.Code, "HTTP request status code error")
}

func TestGetAllProducts_NoProductsSaved(t *testing.T) {
	body_username := []byte(`{"username": "test"}`)
	// create account, but do not save any products
	test_credentials := []byte(`{"username": "test", 
	"password": "test-password"}`)
	req, _ := http.NewRequest("POST", "/api/user/signup", bytes.NewBuffer(test_credentials))
	response := httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	// attempt to get saved products (none)
	req, _ = http.NewRequest("GET", "/api/products", bytes.NewBuffer(body_username))

	response = httptest.NewRecorder()
	Router.ServeHTTP(response, req)

	a := assert.New(t)
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")
	a.Equal("[]\n", response.Body.String(), "Returns an incorrect, nonempty list")
}
