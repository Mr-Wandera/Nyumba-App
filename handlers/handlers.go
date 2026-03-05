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

// LoadData restores the 2-argument version required by main.go
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
	// Using hardcoded context for Abdul
	isLoggedIn := "true"
	currentUsername := "Abdul"

	// Combines the structural HTML with the Dynamic Scripts
	htmlBody := templates.GetHTML(isLoggedIn, currentUsername, "", "none")
	scripts := templates.GetScripts(isLoggedIn == "true", currentUsername)

	fmt.Fprint(w, htmlBody+scripts)
}

// These resolve all 'undefined' errors in your main.go
func LoginHandler(w http.ResponseWriter, r *http.Request)  { fmt.Fprint(w, "Login Page") }
func SignupHandler(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Signup Page") }
func LogoutHandler(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/", 302) }
func GetHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}
func UploadHouse(w http.ResponseWriter, r *http.Request)        { fmt.Fprint(w, "Upload Service") }
func PayHandler(w http.ResponseWriter, r *http.Request)         { fmt.Fprint(w, "Payment Processing") }
func CallbackHandler(w http.ResponseWriter, r *http.Request)    { fmt.Fprint(w, "Callback Service") }
func ServeMedia(w http.ResponseWriter, r *http.Request)         { fmt.Fprint(w, "Media Service") }
func DeleteHouseHandler(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Delete Logic") }

func SeedHouses() {
	if len(houses) > 0 {
		return
	}
	// Bringing back your original House Data
	houses = append(houses, models.House{
		ID: 1, BuildingName: "Sample Sanctuary", Location: "Thika",
		Details:   "A beautiful home near MKU.",
		ImageURLs: []string{"https://images.unsplash.com/photo-1570129477492-45c003edd2be"},
	})
}
