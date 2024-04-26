package main

import (
	"log"
	"myproject/internal/api/routes"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Starting server on port %s...\n", port)
	http.Handle("/", routes.NewRouter())
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
