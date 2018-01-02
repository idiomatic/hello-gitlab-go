package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoot(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Error("OK response expected")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("reading body:", err)
	}
	if strings.TrimSpace(string(body)) != `<p>hello, world.</p>` {
		t.Error("Unexpected body")
	}
}

func TestHelloJSON(t *testing.T) {
	request, _ := http.NewRequest("GET", "/hello.json", nil)
	response := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Error("OK response expected")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("reading body:", err)
	}
	if strings.TrimSpace(string(body)) != `{"message":"aloha"}` {
		t.Error("Unexpected body")
	}
}

func Test404(t *testing.T) {
	request, _ := http.NewRequest("GET", "/404", nil)
	response := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(response, request)
	if response.Code != http.StatusNotFound {
		t.Error("Not Found response expected")
	}
}
