package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skirill430/Quick-Shop/server/router"
	"github.com/skirill430/Quick-Shop/server/utils"

	"github.com/rs/cors"
)

var temp string

func main() {
	r := router.Router()
	utils.ConnectDB("users")
	utils.ConnectDB("products")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "http://localhost", "http://132.145.212.18", "http://132.145.212.18/home", "http://132.145.212.18/login", "http://132.145.212.18/cart", "http://quickshop.hopto.org", "https://quickshop.hopto.org"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	fmt.Println("Server running on Port 9000...")
	if temp == "1" {
		log.Fatal(http.ListenAndServeTLS(":9000", "/etc/letsencrypt/live/quickshop.hopto.org/cert.pem", "/etc/letsencrypt/live/quickshop.hopto.org/privkey.pem", corsHandler.Handler(r)))
	} else {
		log.Fatal(http.ListenAndServe(":9000", corsHandler.Handler(r)))
	}
}
