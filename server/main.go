package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skirill430/Quick-Shop/server/utils"

	"io/ioutil"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello-world", helloWorld)
	r.HandleFunc("/walmart", walmart)

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

func walmart(w http.ResponseWriter, r *http.Request) {

	var search string

	fmt.Print("Search for: ")
	fmt.Scanf("%s", &search)

	//maybe modify the url?
	url := "https://walmart.p.rapidapi.com/products/v2/list?cat_id=0&sort=best_seller&page=1&query="

	url = url + search

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "813328aa2cmshbaf8f8dc041bb3ep1a7203jsnc647b6bcd1c7")
	req.Header.Add("X-RapidAPI-Host", "walmart.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
