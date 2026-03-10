package main

import (
	"fmt"
	"net/http"
	"nyumba/handlers"
	"os"
)

func main() {
	// 1. Initialize data from houses.json
	handlers.LoadData("houses.json", &handlers.Houses)
	handlers.SeedHouses()

	// 2. Page Routes
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/explore", handlers.ExploreHandler)

	// 3. API & Data Routes
	http.HandleFunc("/houses", handlers.GetHouses)
	http.HandleFunc("/add-house", handlers.AddHouseHandler)

	// 4. M-Pesa Integration Routes (Critical for the Unlock button)
	http.HandleFunc("/trigger-payment", handlers.TriggerPaymentHandler)
	http.HandleFunc("/mpesa/callback", handlers.MpesaCallbackHandler)

	// 5. Static File Serving for Uploaded Photos
	fs := http.FileServer(http.Dir("./uploads"))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Nyumba Server successfully running on port:", port)
	http.ListenAndServe(":"+port, nil)
}
