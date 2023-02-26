package router

import (
	"github.com/gorilla/mux"
	"github.com/skirill430/Quick-Shop/server/handlers"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/walmart", handlers.Walmart)
	r.HandleFunc("/signup/{username}/{password}", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/signin/{username}/{password}", handlers.AuthenticateUser).Methods("POST")

	return r
}
