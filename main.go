package main

import (
	"fmt"
	"net/http"
	"nyumba/handlers"
	"nyumba/models"
	"os"
)

var houses []models.House

func main() {
	handlers.LoadData("houses.json", &houses)
	handlers.SeedHouses()

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/explore", handlers.ExploreHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/houses", handlers.GetHouses)
	http.HandleFunc("/add-house", handlers.AddHouseHandler) //

	// Fix Port for Render
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Sanctuary live on port:", port)
	http.ListenAndServe(":"+port, nil)
}
