package main

import (
	"log"
	"myproject/internal/api/routes"
	"net/http"
)

func main() {
	log.Println("Starting server on port 8080...")
	http.Handle("/", routes.NewRouter())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
