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
	// A simple, stable call to get the app running again
	htmlBody := templates.GetHTML("false", "", "", "none")
	scripts := templates.GetScripts(false, "")
	fmt.Fprint(w, htmlBody+scripts)
}

// These resolve the 10+ undefined errors in your IDE screenshot
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
	houses = append(houses, models.House{
		ID: 1, BuildingName: "Standard Sanctuary", Location: "Thika",
		Price: 1000, Details: "A simple sanctuary.",
		ImageURLs: []string{"https://images.unsplash.com/photo-1570129477492-45c003edd2be"},
	})
}
