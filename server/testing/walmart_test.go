package test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/skirill430/Quick-Shop/server/handlers"
)

func TestWalmart(t *testing.T) {
	var jsonStr = []byte("shoes")

	req := httptest.NewRequest(http.MethodGet, "/walmart", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()

	handlers.Walmart(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if !(strings.Contains(string(data), "Results for") && strings.Contains(string(data), "\"query\":\"shoes\"")) {
		t.Fail()
	}
}

func TestTarget(t *testing.T) {
	var jsonStr = []byte("shoes")

	req := httptest.NewRequest(http.MethodGet, "/Target", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()

	handlers.Target(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if !(strings.Contains(string(data), "\"search_term\":\"shoes\"")) {
		t.Fail()
	}
}
