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

// LoadData reads the JSON database
func LoadData(filename string, target interface{}) {
	file, err := os.ReadFile(filename)
	if err == nil {
		json.Unmarshal(file, target)
	}
}

// GetHouses supports filtering by location (e.g., /houses?location=Section9)
func GetHouses(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	var filtered []models.House

	for _, h := range houses {
		if location == "" || h.Location == location {
			filtered = append(filtered, h)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

// AddHouseHandler processes the sidebar form and redirects to explore
func AddHouseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		newHouse := models.House{
			ID:           len(houses) + 1,
			BuildingName: r.FormValue("building_name"),
			Location:     r.FormValue("location"),
			Price:        7500, // Default price placeholder
			ImageURLs:    []string{"https://images.unsplash.com/photo-1570129477492-45c003edd2be"},
		}
		houses = append(houses, newHouse)
		http.Redirect(w, r, "/explore", http.StatusSeeOther)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetLandingHTML())
}

func ExploreHandler(w http.ResponseWriter, r *http.Request) {
	currentUsername := "Abdul" // cite: User Summary
	html := templates.GetHTML("true", currentUsername, "", "none") + templates.GetScripts(true, currentUsername)
	fmt.Fprint(w, html)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, templates.GetSignupHTML())
		return
	}
	http.Redirect(w, r, "/explore", http.StatusSeeOther)
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
