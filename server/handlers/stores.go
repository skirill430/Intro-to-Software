package handlers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Walmart(w http.ResponseWriter, r *http.Request) {

	var search []byte
	var err error
	var append string

	//https://www.informit.com/articles/article.aspx?p=2861456&seqNum=7 got the following switch statement form here
	search, err = io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	append = bytes.NewBuffer(search).String()

	url := "https://walmart.p.rapidapi.com/products/v2/list?cat_id=0&sort=best_seller&page=1&query="

	fmt.Println(append)
	url = url + append
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "813328aa2cmshbaf8f8dc041bb3ep1a7203jsnc647b6bcd1c7")
	req.Header.Add("X-RapidAPI-Host", "walmart.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("An error connecting to Walmart occurred.")
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	w.Write(body)
	//fmt.Println(res)
	//fmt.Println(string(body))

}

func Target(w http.ResponseWriter, r *http.Request) {

	var search []byte
	var err error
	var append string

	search, err = io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	append = bytes.NewBuffer(search).String()

	url := "https://target1.p.rapidapi.com/products/v2/list?store_id=911&category=5xtg6&keyword="
	url2 := "&count=20&offset=0&default_purchasability_filter=true&sort_by=relevance"

	fmt.Println(append)
	url = url + append + url2

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "813328aa2cmshbaf8f8dc041bb3ep1a7203jsnc647b6bcd1c7")
	req.Header.Add("X-RapidAPI-Host", "target1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("An error connecting to Target occurred.")
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	w.Write(body)
	//fmt.Println(res)
	//fmt.Println(string(body))

}
