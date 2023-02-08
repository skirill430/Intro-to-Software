package main

import (
	"bytes"
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

	var search []byte
	var err error
	var append string

	//https://www.informit.com/articles/article.aspx?p=2861456&seqNum=7 got the following switch statement form here
	search, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	append = bytes.NewBuffer(search).String()

	url := "https://walmart.p.rapidapi.com/products/v2/list?cat_id=0&sort=best_seller&page=1&query="

	fmt.Printf(append)
	url = url + append
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "813328aa2cmshbaf8f8dc041bb3ep1a7203jsnc647b6bcd1c7")
	req.Header.Add("X-RapidAPI-Host", "walmart.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Write(body)
	//fmt.Println(res)
	//fmt.Println(string(body))

}
