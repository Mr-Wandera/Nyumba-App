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

// Fixes the blank screen by handling the POST request from 'Start Journey'
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, templates.GetSignupHTML())
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		// Success! Redirecting to Explore page
		http.Redirect(w, r, "/explore", http.StatusSeeOther)
	}
}

func ExploreHandler(w http.ResponseWriter, r *http.Request) {
	currentUsername := "Abdul" // cite: User Summary
	html := templates.GetHTML("true", currentUsername, "", "none") + templates.GetScripts(true, currentUsername)
	fmt.Fprint(w, html)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetLandingHTML())
}

func GetHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}

func LoadData(filename string, target interface{}) {
	file, err := os.ReadFile(filename)
	if err == nil {
		json.Unmarshal(file, target)
	}
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

func LoginHandler(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, templates.GetSignupHTML()) }
