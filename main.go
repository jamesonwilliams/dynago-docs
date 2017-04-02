package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {

	router := NewRouter()
	corsOpts := handlers.AllowedOrigins([]string{"*"})
	corsHandler := handlers.CORS(corsOpts)(router)
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
