package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nyumba/models"
	"nyumba/templates"
	"os"
)

var houses []models.House

func LoadData(filename string, target interface{}) {
	file, _ := os.ReadFile(filename)
	json.Unmarshal(file, target)
}

func SeedHouses() {
	if len(houses) > 0 {
		return
	}
	houses = append(houses, models.House{
		ID: 1, BuildingName: "Base Apartments", Location: "Thika",
		Price: 7500, ImageURLs: []string{"https://images.unsplash.com/photo-1570129477492-45c003edd2be"},
	})
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetLandingHTML())
}

func ExploreHandler(w http.ResponseWriter, r *http.Request) {
	html := templates.GetHTML("true", "Abdul", "", "none") + templates.GetScripts(true, "Abdul")
	fmt.Fprint(w, html)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, templates.GetSignupHTML())
		return
	}
	// Handles the POST request to prevent blank screen
	http.Redirect(w, r, "/explore", http.StatusSeeOther)
}

func GetHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}
