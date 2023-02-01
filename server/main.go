package main

import (
	"example/go-server/server/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello-world", helloWorld)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})

	fmt.Println("Server running on Port 9000...")
	log.Fatal(http.ListenAndServe(":9000", corsHandler.Handler(r)))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO")
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Golang + Angular Starter Kit",
	}

	jsonBytes, err := utils.StructToJson(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
