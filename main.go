package main

import (
	"fmt"
	"net/http"
	"os"

	"nyumba/handlers" // ✅ FIXED: Using your real module name 'nyumba'
)

func main() {
	// 1. Initialize Data
	handlers.LoadData()
	handlers.SeedHouses()

	// 2. Ensure uploads folder exists
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	// 3. Define Routes
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/callback", handlers.CallbackHandler)

	// API & Features
	http.HandleFunc("/houses", handlers.GetHouses)
	http.HandleFunc("/houses/upload", handlers.UploadHouse)
	http.HandleFunc("/houses/delete", handlers.DeleteHouseHandler)
	http.HandleFunc("/pay", handlers.PayHandler)
	http.HandleFunc("/uploads/", handlers.ServeMedia)

	// 4. Start Server
	fmt.Println("🚀 Nyumba App running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
