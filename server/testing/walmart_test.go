package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/skirill430/Quick-Shop/server/handlers"
)

var passed = 0

func TestWalmart(t *testing.T) {

	fmt.Println("Walmart Test Running...")
	var jsonStr = []byte("shoes")

	req := httptest.NewRequest(http.MethodGet, "/walmart", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()

	handlers.Walmart(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	//fmt.Println(res)
	//fmt.Println(string(data))
	if strings.Contains(string(data), "Results for") && strings.Contains(string(data), "\"query\":\"shoes\"") {
		fmt.Println("Walmart results for \"shoes\" ... Displayed!")
		passed = passed + 1
	}
	fmt.Println("Tests Passed: ", passed)
}

func TestTarget(t *testing.T) {

	fmt.Println("Target Test Running...")
	var jsonStr = []byte("shoes")

	req := httptest.NewRequest(http.MethodGet, "/Target", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()

	handlers.Target(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	//fmt.Println(res)
	//fmt.Println(string(data))
	if strings.Contains(string(data), "\"search_term\":\"shoes\"") {
		fmt.Println("Target results for \"shoes\" ... Displayed!")
		passed = passed + 1
	}
	fmt.Println("Tests Passed: ", passed)
}
