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
var users []models.User

// LoadData fixes the 'WrongArgCount' error in main.go
func LoadData(filename string, target interface{}) {
	file, err := os.ReadFile(filename)
	if err == nil {
		json.Unmarshal(file, target)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetLandingHTML())
}

func ExploreHandler(w http.ResponseWriter, r *http.Request) {
	// Combines the structural UI with the scripts to load houses
	html := templates.GetHTML("true", "Abdul", "", "none") + templates.GetScripts(true, "Abdul")
	fmt.Fprint(w, html)
}

func GetHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, templates.GetSignupHTML())
		return
	}
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/explore", http.StatusSeeOther)
	}
}

// Minimal handlers to resolve main.go undefined errors
func LoginHandler(w http.ResponseWriter, r *http.Request)    { fmt.Fprint(w, "Login Page") }
func LogoutHandler(w http.ResponseWriter, r *http.Request)   { http.Redirect(w, r, "/", 302) }
func UploadHouse(w http.ResponseWriter, r *http.Request)     { fmt.Fprint(w, "Upload Service") }
func CallbackHandler(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Callback Service") }

func SeedHouses() {
	if len(houses) > 0 {
		return
	}
	houses = append(houses, models.House{
		ID: 1, BuildingName: "Thika Sanctuary", Location: "Thika",
		Price: 7500, Details: "A premium unit near MKU.",
		ImageURLs: []string{"https://images.unsplash.com/photo-1570129477492-45c003edd2be"},
	})
}
