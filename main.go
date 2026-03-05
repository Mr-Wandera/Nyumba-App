package main

import (
	"fmt"
	"net/http"
	"nyumba/handlers"
	"nyumba/models"
	"os"
)

// Declare these globally so they are shared across the app
var houses []models.House
var users []models.User

func main() {
	// 1. Load your existing data
	handlers.LoadData("houses.json", &houses)
	handlers.SeedHouses()

	// 2. Define your routes for every button
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/explore", handlers.ExploreHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/houses", handlers.GetHouses) // Matches GetScripts fetch

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("🚀 Nyumba Sanctuary live on port:", port)
	http.ListenAndServe(":"+port, nil)
}
