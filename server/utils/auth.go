package utils

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// returns cookie, http status code
func GenerateUsernameCookie(username string) (*http.Cookie, int) {
	// Declare the expiration time of the token: 24 hours
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	cookies := Cookies{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cookies)
	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return nil, http.StatusInternalServerError
	}

	return &http.Cookie{
		Name:    "token",
		Domain:  "localhost",
		Path:    "/",
		Value:   tokenString,
		Expires: expirationTime,
	}, 200
}

// returns username, http status code
func ExtractUsernameFromCookie(w http.ResponseWriter, r *http.Request) (string, int) {
	// unpack the cookie in request
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return "", http.StatusUnauthorized
		}
		// For any other type of error, return a bad request status
		return "", http.StatusBadRequest
	}
	tknStr := c.Value
	claims := &Cookies{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", http.StatusUnauthorized
		}
		return "", http.StatusBadRequest
	}
	if !tkn.Valid {
		return "", http.StatusUnauthorized
	}

	return claims.Username, http.StatusOK
}
