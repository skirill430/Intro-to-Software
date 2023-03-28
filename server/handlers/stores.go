package handlers

import (
	"strings"
	//"encoding/json"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ItemData struct {
	price    string
	imgUrl   string
	rating   string
	name     string
	store_id string
}

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
}

func BothStores(w http.ResponseWriter, r *http.Request) {
	var search []byte
	var err error
	var append string

	var itemString string
	//var itemsJson []ItemData

	search, err = io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	append = bytes.NewBuffer(search).String()

	url := "https://target1.p.rapidapi.com/products/v2/list?store_id=911&category=5xtg6&keyword="
	url2 := "&count=20&offset=0&default_purchasability_filter=true&sort_by=relevance"

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

	//fmt.Printf("made it this far")

	NewJson := `[`
	i := 0
	itemString = string(body)
	//fmt.Printf(itemString)
	i2 := len(itemString) - 1
	i3 := 0
	for i != -1 {
		i = strings.Index(itemString, "parent")
		if i == -1 {
			break
		}
		i2 = len(itemString) - 1
		itemString = itemString[i:i2]
		//fmt.Printf("made it this far3")

		i = strings.Index(itemString, "primary_image_url")
		if i == -1 {
			break
		}
		//fmt.Printf("made it this far4")
		i2 = len(itemString) - 1
		fmt.Printf("i = %d, i3 = %d\n", i, i2)
		itemString = itemString[i:i2]
		//fmt.Printf("made it this far5")
		i3 = strings.Index(itemString, "}")
		//fmt.Printf("made it this far6")
		fmt.Printf("i = %d, i3 = %d\n", i, i3)
		NewJson = NewJson + "{\"image_url\":"
		NewJson = NewJson + itemString[19:i3-1] //contains the url
		NewJson = NewJson + "\","

		//fmt.Printf("made it this far4")

		i = strings.Index(itemString, "title")
		if i == -1 {
			break
		}
		i2 = len(itemString) - 1
		itemString = itemString[i:i2]
		i3 = strings.Index(itemString, ",")
		fmt.Printf("i = %d, i3 = %d\n", i, i3)
		NewJson = NewJson + "\"product_name\":"
		NewJson = NewJson + itemString[7:i3-1] //contains the name
		if strings.Contains(itemString[8:i3-1], "\"") {
			NewJson = NewJson + ","
		} else {
			NewJson = NewJson + "\","
		}

		//fmt.Printf("made it this far5")

		i = strings.Index(itemString, "_price")
		if i == -1 {
			break
		}
		i2 = len(itemString) - 1
		itemString = itemString[i:i2]
		i3 = strings.Index(itemString, ",")
		fmt.Printf("i = %d, i3 = %d\n", i, i3)
		NewJson = NewJson + "\"price\":"
		NewJson = NewJson + itemString[8:i3-1] //contains the price
		NewJson = NewJson + "\","

		//fmt.Printf("made it this far6")

		i = strings.Index(itemString, "average")
		if i == -1 {
			break
		}
		i2 = len(itemString) - 1
		itemString = itemString[i:i2]
		i3 = strings.Index(itemString, ",")
		fmt.Printf("i = %d, i3 = %d\n", i, i3)
		NewJson = NewJson + "\"rating\":"
		NewJson = NewJson + "\"" + itemString[9:i3-1] + "\"" //contains the rating
		NewJson = NewJson + ","

		NewJson = NewJson + "\"seller_name\":\"Target\"},"

		//fmt.Printf("made it this far7")
		//fmt.Printf(NewJson)
	}
	NewJson = NewJson[0 : len(NewJson)-2]
	NewJson = NewJson + "}]"
	fmt.Printf(NewJson)

	/*
		search, err = io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		append = bytes.NewBuffer(search).String()

		url = "https://walmart.p.rapidapi.com/products/v2/list?cat_id=0&sort=best_seller&page=1&query="

		url = url + append
		req, _ = http.NewRequest("GET", url, nil)

		req.Header.Add("X-RapidAPI-Key", "813328aa2cmshbaf8f8dc041bb3ep1a7203jsnc647b6bcd1c7")
		req.Header.Add("X-RapidAPI-Host", "walmart.p.rapidapi.com")

		res, err = http.DefaultClient.Do(req)

		if err != nil {
			fmt.Println("An error connecting to Walmart occurred.")
			return
		}

		defer res.Body.Close()
		body, _ = io.ReadAll(res.Body)

		i = 0
		itemString = string(body)
		i2 = len(itemString) - 1
		i3 = 0
		for i != -1 {
			i = strings.Index(itemString, "\"Product\"")
			itemString = itemString[i:i2]
			i = strings.Index(itemString, "\"name\"")
			i3 = strings.Index(itemString, ",")
			NewJson = NewJson + "{\"name\":"
			NewJson = NewJson + itemString[i+7:i3-1] //contains the name
			NewJson = NewJson + ","

			i = strings.Index(itemString, "thumbnailUrl")
			itemString = itemString[i:i2]
			i3 = strings.Index(itemString, "}")
			NewJson = NewJson + "\"imgUrl\":"
			NewJson = NewJson + itemString[i+14:i3-1] //contains the url
			NewJson = NewJson + ","

			i = strings.Index(itemString, "averageRating")
			itemString = itemString[i:i2]
			i3 = strings.Index(itemString, ",")
			NewJson = NewJson + "\"rating\":"
			NewJson = NewJson + "\"" + itemString[i+15:i3-1] + "\"" //contains the rating
			NewJson = NewJson + ","

			i = strings.Index(itemString, "$")
			itemString = itemString[i:i2]
			i3 = strings.Index(itemString, ",")
			NewJson = NewJson + "\"price\":"
			NewJson = NewJson + "\"" + itemString[i+1:i3-1] //contains the price
			NewJson = NewJson + ","

			NewJson = NewJson + "\"store_id\":\"Walmart\"},"

		}
	*/
	//json.Unmarshal([]byte(NewJson), &itemsJson)

	w.Write([]byte(NewJson))
}
