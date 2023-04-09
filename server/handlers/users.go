package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"fmt"
	"time"

	"github.com/skirill430/Quick-Shop/server/utils"

	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

func Welcome(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &utils.Cookies{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return utils.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
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
