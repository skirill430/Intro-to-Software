package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Walmart(w http.ResponseWriter, r *http.Request) {

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
