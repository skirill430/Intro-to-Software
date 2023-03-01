package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skirill430/Quick-Shop/server/router"
	"github.com/skirill430/Quick-Shop/server/utils/db"

	"github.com/rs/cors"
)

func main() {
	r := router.Router()
	db.ConnectDB("users")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})

	fmt.Println("Server running on Port 9000...")
	// "127.0.0.1" before port disables firewall popup when running dev environment
	log.Fatal(http.ListenAndServe("127.0.0.1:9000", corsHandler.Handler(r)))
}
