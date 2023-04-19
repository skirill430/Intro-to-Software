package handlers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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

	NewJson := `[`
	i := 0
	temp1 := 0
	temp2 := 0
	itemString = string(body)
	var i2 int
	i3 := 0
	count := 0
	temp1 = strings.Index(itemString, "enrichment\":{\"buy_url")
	temp2 = strings.Index(itemString, "\"Product\"")
	if temp1 == -1 && temp2 == -1 {
		NewJson = "{}"
	} else {
		for i != -1 {
			if count >= 20 {
				break
			}
			i = strings.Index(itemString, "enrichment\":{\"buy_url")
			if i == -1 {
				break
			}
			i = strings.Index(itemString, "primary_image_url")
			if i == -1 {
				break
			}
			i = strings.Index(itemString, "\"title\"")
			if i == -1 {
				break
			}
			i = strings.Index(itemString, "formatted_current_price\"")
			if i == -1 {
				break
			}
			i = strings.Index(itemString, "average")
			if i == -1 {
				break
			}

			i = strings.Index(itemString, "enrichment\":{\"buy_url")

			i2 = len(itemString) - 1
			itemString = itemString[i:i2]

			i = strings.Index(itemString, "primary_image_url")
			i2 = len(itemString) - 1
			itemString = itemString[i:i2]
			i3 = strings.Index(itemString, "}")
			NewJson = NewJson + "{\"image_url\":"
			NewJson = NewJson + itemString[19:i3-1] //contains the url
			NewJson = NewJson + "\","

			i = strings.Index(itemString, "\"title\"")
			i2 = len(itemString) - 1
			itemString = itemString[i:i2]
			i3 = strings.Index(itemString, "}")
			NewJson = NewJson + "\"product_name\""
			NewJson = NewJson + itemString[7:i3] //contains the name
			if strings.Contains(itemString[8:i3], "\"") {
				NewJson = NewJson + ","
			} else {
				NewJson = NewJson + "\","
			}

			i = strings.Index(itemString, "formatted_current_price\"")
			i2 = len(itemString) - 1
			itemString = itemString[i:i2]
			i3 = strings.Index(itemString, ",")
			NewJson = NewJson + "\"price\":"
			NewJson = NewJson + "\"" + itemString[26:i3] //contains the price
			if strings.Contains(itemString[27:i3], "\"") {
				if NewJson[len(NewJson)-1:] != "\"" {
					NewJson = NewJson + "\""
				}
				NewJson = NewJson + ","
			} else {
				NewJson = NewJson + "\","
			}

			i = strings.Index(itemString, "average")
			i2 = len(itemString) - 1
			itemString = itemString[i:i2]
			i3 = strings.Index(itemString, ",")
			temp := itemString[9 : i3-1]
			tempi := strings.Index(temp, "nul")
			if tempi != -1 {
				NewJson = NewJson + "\"rating\":"
				NewJson = NewJson + "\"" + "" + "\"" //contains the rating
				NewJson = NewJson + ","
			} else {
				NewJson = NewJson + "\"rating\":"
				NewJson = NewJson + "\"" + itemString[9:i3-1] + "\"" //contains the rating
				NewJson = NewJson + ","
			}
			NewJson = NewJson + "\"seller_name\":\"Target\"},"

			count++
		}
		NewJson = NewJson[0 : len(NewJson)-2]
		NewJson = NewJson + "},"

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
		i3 = 0
		count = 0
		for i != -1 {
			if count >= 20 {
				break
			}
			i = strings.Index(itemString, "\"Product\"")
			if i == -1 {
				break
			}
			i = strings.Index(itemString, "\"name\"")
			if i == -1 {
				break
			}
			i = strings.Index(itemString, "thumbnailUrl")
			if i == -1 {
				break
			}
			i = strings.Index(itemString, "averageRating")
			if i == -1 {
				break
			}
			i = strings.Index(itemString, "$")
			if i == -1 {
				break
			}

			i = strings.Index(itemString, "\"Product\"")

			i2 = len(itemString) - 1
			itemString = itemString[i:i2]
			i = strings.Index(itemString, "\"name\"")
			i2 = len(itemString) - 1
			itemString = itemString[i:i2]
			i3 = strings.Index(itemString, ",")
			NewJson = NewJson + "{\"product_name\":"
			NewJson = NewJson + itemString[7:i3-1] //contains the name
			NewJson = NewJson + "\","

			i = strings.Index(itemString, "thumbnailUrl")

			i2 = len(itemString) - 1
			itemString = itemString[i:i2]
			i3 = strings.Index(itemString, "}")
			NewJson = NewJson + "\"image_url\":"
			NewJson = NewJson + " \"" + itemString[15:i3-1] //contains the url
			if strings.Contains(itemString[16:i3-1], "\"") {
				NewJson = NewJson + ","
			} else {
				NewJson = NewJson + "\","
			}

			i = strings.Index(itemString, "averageRating")

			i2 = len(itemString) - 1
			itemString = itemString[i:i2]
			i3 = strings.Index(itemString, ",")
			temp := itemString[15:i3]
			tempi := strings.Index(temp, "nul")
			if tempi != -1 {
				NewJson = NewJson + "\"rating\":"
				NewJson = NewJson + " \"" + "" //contains the rating
				NewJson = NewJson + "\","
			} else {
				NewJson = NewJson + "\"rating\":"
				NewJson = NewJson + " \"" + itemString[15:i3] //contains the rating
				NewJson = NewJson + "\","
			}

			i = strings.Index(itemString, "$")
			i2 = len(itemString) - 1
			itemString = itemString[i:i2]
			i3 = strings.Index(itemString, ",")
			if strings.Contains(itemString[1:i3-1], "fals") {
				NewJson = NewJson + "\"price\":" + "\"fals\","
			} else {
				NewJson = NewJson + "\"price\":"
				NewJson = NewJson + "\"" + "$" + itemString[1:i3-1] //contains the price
				if strings.Contains(itemString[1:i3-1], "\"") {
					NewJson = NewJson + ","
				} else {
					NewJson = NewJson + "\","
				}
			}
			NewJson = NewJson + "\"seller_name\":\"Walmart\"},"

			count++
		}
		NewJson = NewJson[0 : len(NewJson)-2]
		NewJson = NewJson + "}]"
	}

	w.Write([]byte(NewJson))
}
