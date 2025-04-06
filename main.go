package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yashK-2703/go-url-shortener/handlers"
	"github.com/yashK-2703/go-url-shortener/storage"

	"github.com/gorilla/mux"
)

func main() {
	// Init Redis
	storage.InitializeRedis()

	// Setup Router
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/shorten", handlers.ShortenHandler).Methods("POST")
	r.HandleFunc("/r/{id}", handlers.RedirectHandler).Methods("GET")

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
