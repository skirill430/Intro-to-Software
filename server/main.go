package main

import (
	"example/go-server/server/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello-world", helloWorld)

	fmt.Println("Server running on Port 4200...")
	log.Fatal(http.ListenAndServe(":4200", r))
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
