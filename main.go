package main

import (
	"fmt"
	"net/http"
	"nyumba/handlers"
	"nyumba/models"
	"os"
)

// Declare these globally
var houses []models.House
var users []models.User

func main() {
	// 1. Load Data correctly
	handlers.LoadData("houses.json", &houses)
	handlers.SeedHouses()

	// 2. Define your routes
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/explore", handlers.ExploreHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/houses", handlers.GetHouses) // Matches handlers.GetHouses

	// 3. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("🚀 Nyumba Sanctuary live on port:", port)
	http.ListenAndServe(":"+port, nil)
}
