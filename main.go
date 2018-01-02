package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{"aloha"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func init() {
	http.HandleFunc("/hello.json", helloHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))
}

func main() {
	var addr = ":8000"
	if port, ok := os.LookupEnv("PORT"); ok {
		addr = ":" + port
	}
	log.Println("listening to", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
