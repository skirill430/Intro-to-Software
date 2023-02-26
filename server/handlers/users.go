package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/skirill430/Quick-Shop/server/utils/db"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	// check credentials are valid
	if vars["username"] == "" || vars["password"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request Body Missing Fields"))
		return
	}

	user := &db.User{
		Username: vars["username"],
		Password: vars["password"],
		List:     "",
	}

	// add user only if username doesn't exist in database already
	res := db.DB.Where("username = ?", user.Username).FirstOrCreate(&user)

	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sorry, this username is already taken. Enter another username."))
		return
	}
}
