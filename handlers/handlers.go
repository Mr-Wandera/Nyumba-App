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
	// This links to your premium landing page in ui.go
	fmt.Fprint(w, templates.GetLandingHTML())
}

func ExploreHandler(w http.ResponseWriter, r *http.Request) {
	// These values simulate your session state
	isLoggedIn := "true"
	currentUsername := "Abdul"

	// Combines the structural UI with the house-loading scripts
	html := templates.GetHTML(isLoggedIn, currentUsername, "", "none") + templates.GetScripts(true, currentUsername)
	fmt.Fprint(w, html)
}
func GetHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprint(w, templates.GetSignupHTML())
		return
	}
	// Add your signup logic here (saving to users.json)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// You can create a GetLoginHTML similarly
		fmt.Fprint(w, templates.GetSignupHTML())
		return
	}
}
func LogoutHandler(w http.ResponseWriter, r *http.Request)      { http.Redirect(w, r, "/", 302) }
func UploadHouse(w http.ResponseWriter, r *http.Request)        { fmt.Fprint(w, "Upload Service") }
func PayHandler(w http.ResponseWriter, r *http.Request)         { fmt.Fprint(w, "Payment Processing") }
func CallbackHandler(w http.ResponseWriter, r *http.Request)    { fmt.Fprint(w, "M-Pesa Callback") }
func ServeMedia(w http.ResponseWriter, r *http.Request)         { fmt.Fprint(w, "Media Service") }
func DeleteHouseHandler(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Delete Logic") }
func SeedHouses() {
	if len(houses) > 0 {
		return
	}
	houses = append(houses, models.House{
		ID: 1, BuildingName: "Sample Sanctuary", Location: "Thika",
		Price: 1000, Details: "A beautiful home near MKU.",
		ImageURLs: []string{"https://images.unsplash.com/photo-1570129477492-45c003edd2be"},
	})
}
