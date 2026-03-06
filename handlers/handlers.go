package handlers

import (
	"encoding/json"
	"fmt"
	"io"
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

func AddHouseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// 1. Parse the multipart form
		r.ParseMultipartForm(10 << 20) // 10MB limit

		// 2. Handle the Image Upload
		file, handler, err := r.FormFile("property_photo")
		imagePath := "/uploads/default.jpg" // Default if no image is uploaded

		if err == nil {
			defer file.Close()
			// Create a unique filename and save to uploads folder
			imagePath = "/uploads/" + handler.Filename
			f, _ := os.OpenFile("."+imagePath, os.O_WRONLY|os.O_CREATE, 0666)
			defer f.Close()
			io.Copy(f, file)
		}

		// 3. Save the new House with the real image path
		newHouse := models.House{
			ID:           len(houses) + 1,
			BuildingName: r.FormValue("building_name"),
			Location:     r.FormValue("location"),
			Price:        7500,
			ImageURLs:    []string{imagePath}, // Use the uploaded image
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
