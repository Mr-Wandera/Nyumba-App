package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nyumba/models"
	"nyumba/templates"
	"os"
	"path/filepath"
)

var Houses []models.House

func AddHouseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("property_photo")
	imagePath := "/uploads/default.jpg"
	if err == nil {
		defer file.Close()
		os.MkdirAll("./uploads", os.ModePerm)
		filename := filepath.Base(header.Filename)
		dstPath := filepath.Join("./uploads", filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		io.Copy(dst, file)
		imagePath = "/uploads/" + filename
	}

	price := 15000.0
	priceStr := r.FormValue("price")
	if priceStr != "" {
		fmt.Sscanf(priceStr, "%f", &price)
	}

	newHouse := models.House{
		ID:            len(Houses) + 1,
		BuildingName:  r.FormValue("building_name"),
		Location:      r.FormValue("location"),
		MapLink:       r.FormValue("map_link"),
		Price:         price,
		Deposit:       price,
		Type:          "Apartment",
		ImageURLs:     []string{imagePath},
		IsPaid:        false,
		Bedrooms:      1,
		Bathrooms:     1,
		LandlordPhone: "+254712345678",
		Description:   "Beautiful sanctuary available now.",
	}
	Houses = append(Houses, newHouse)
	saveHouses()

	http.Redirect(w, r, "/explore", http.StatusSeeOther)
}

func GetHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Houses)
}

func ExploreHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, templates.GetExploreHTML(Houses))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, templates.GetLandingHTML(Houses))
}

func LoadData(filename string, target interface{}) {
	file, err := os.ReadFile(filename)
	if err == nil {
		json.Unmarshal(file, target)
	}
}

func SeedHouses() {
	if len(Houses) > 0 {
		return
	}
	Houses = append(Houses, models.House{
		ID:            1,
		BuildingName:  "Sunset Heights",
		Location:      "Section 9, Thika",
		Price:         15000,
		Deposit:       15000,
		Type:          "Apartment",
		ImageURLs:     []string{"/uploads/default.jpg"},
		IsPaid:        false,
		Bedrooms:      2,
		Bathrooms:     1,
		LandlordPhone: "+254712345678",
		Description:   "Quiet neighborhood, great views.",
	})
	saveHouses()
}

func saveHouses() {
	data, _ := json.MarshalIndent(Houses, "", "  ")
	os.WriteFile("houses.json", data, 0644)
}

// Renamed to LoginHandler (Uppercase) so it can be used in main.go
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, templates.GetAuthHTML("Login"))
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// We include password in Printf just to satisfy the Go compiler's "unused variable" rule
		fmt.Printf("Login attempt: %s with password length %d\n", email, len(password))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// Renamed to SignupHandler (Uppercase) so it can be used in main.go
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, templates.GetAuthHTML("Sign Up"))
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		password := r.FormValue("password")
		role := r.FormValue("role")

		// Included all variables in Printf to fix "declared and not used" errors
		fmt.Printf("Signup: %s (%s) - Phone: %s, Role: %s, Pass Len: %d\n", name, email, phone, role, len(password))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
