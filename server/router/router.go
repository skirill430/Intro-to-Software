package router

import (
	"github.com/gorilla/mux"
	"github.com/skirill430/Quick-Shop/server/handlers"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/walmart", handlers.Walmart)
	r.HandleFunc("/target", handlers.Target)
	r.HandleFunc("/api/user/signup", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/api/user/signin", handlers.AuthenticateUser).Methods("POST")
	r.HandleFunc("/api/products/save", handlers.SaveProduct).Methods("POST")
	r.HandleFunc("/api/products/remove", handlers.RemoveProduct).Methods("DELETE")

	return r
}
