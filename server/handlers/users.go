package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"time"

	"github.com/skirill430/Quick-Shop/server/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var user utils.User
	json.NewDecoder(r.Body).Decode(&user)

	// check credentials are valid
	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request Body Missing Fields"))
		return
	}

	user.Password = utils.HashAndSalt([]byte(user.Password))
	// add user only if username doesn't exist in database already
	res := utils.UsersDB.Where("username = ?", user.Username).FirstOrCreate(&user)

	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Sorry, this username is already taken. Enter another username."))
		return
	}

	cookie, status := utils.GenerateUsernameCookie(user.Username)
	// if cookie creation failed, return appropriate error
	if status != 200 {
		w.WriteHeader(status)
		return
	}

	http.SetCookie(w, cookie)
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var inputCredentials Credentials
	json.NewDecoder(r.Body).Decode(&inputCredentials)

	// check credentials are valid
	if inputCredentials.Username == "" || inputCredentials.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request Body Missing Fields"))
		return
	}

	var dbUser *utils.User
	// cannot find username in database
	err := utils.UsersDB.First(&dbUser, "username = ?", inputCredentials.Username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("This username does not exist. Create an account."))
		return
	}

	// password does not match
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(inputCredentials.Password))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Incorrect password. Try again."))
		return
	}

	cookie, status := utils.GenerateUsernameCookie(dbUser.Username)
	// if cookie creation failed, return appropriate error
	if status != 200 {
		w.WriteHeader(status)
		return
	}

	http.SetCookie(w, cookie)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// clear the jwt authentication cookie from browser
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Domain:  "localhost",
		Path:    "/",
		Expires: time.Now(),
	})

	w.WriteHeader(http.StatusOK)
}
