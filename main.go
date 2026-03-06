package main

import (
	"fmt"
	"log"
	"net/http"
	"nyumba/handlers"
	"os"
)

func main() {
	// 1. Register Handlers from the handlers package
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/explore", handlers.ExploreHandler)
	http.HandleFunc("/houses", handlers.GetHouses)
	http.HandleFunc("/add-house", handlers.AddHouseHandler)

	// 2. Serve Static Uploads
	// This allows <img src="/uploads/photo.jpg"> to work
	fs := http.FileServer(http.Dir("./uploads"))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	// 3. Environment Setup for Render
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🏠 Nyumba Sanctuary live on port: %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
