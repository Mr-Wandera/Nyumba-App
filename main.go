package main

import (
	"fmt"
	"net/http"
	"nyumba/handlers"
	"os"
)

func main() {
	handlers.LoadData("houses.json", &handlers.Houses)
	handlers.SeedHouses()

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/explore", handlers.ExploreHandler)
	http.HandleFunc("/houses", handlers.GetHouses)
	http.HandleFunc("/add-house", handlers.AddHouseHandler)

	fs := http.FileServer(http.Dir("./uploads"))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port:", port)
	http.ListenAndServe(":"+port, nil)
}
