package main

import (
	"fmt"
	"net/http"
	"os"

	"nyumba/handlers"
	"nyumba/models"
)

// Declare these at the package level so all functions can see them
var houses []models.House
var users []models.User

func main() {
	// 1. Initialize Data
	// This fixes the "WrongArgCount" and "undefined: houses" errors
	handlers.LoadData("houses.json", &houses)
	handlers.LoadData("users.json", &users)
	handlers.SeedHouses()

	// 2. Ensure uploads folder exists
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	// 3. Define Routes
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/explore", handlers.ExploreHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/callback", handlers.CallbackHandler)

	// API & Features
	http.HandleFunc("/houses", handlers.GetHouses)
	http.HandleFunc("/houses/upload", handlers.UploadHouse)
	http.HandleFunc("/pay", handlers.PayHandler)

	// Static Media
	http.HandleFunc("/uploads/", handlers.ServeMedia)

	// 4. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Nyumba Sanctuary is live on port:", port)
	http.ListenAndServe(":"+port, nil)
}
