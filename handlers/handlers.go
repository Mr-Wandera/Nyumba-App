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
		ImageURLs:     []string{imagePath},
		IsPaid:        false,
		Bedrooms:      1,
		Bathrooms:     1,
		LandlordPhone: "+254712345678",
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
	fmt.Fprint(w, templates.GetExploreHTML())
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
		ImageURLs:     []string{"/uploads/default.jpg"},
		IsPaid:        false,
		Bedrooms:      2,
		Bathrooms:     1,
		LandlordPhone: "+254712345678",
	})
	saveHouses()
}

func saveHouses() {
	data, _ := json.MarshalIndent(Houses, "", "  ")
	os.WriteFile("houses.json", data, 0644)
}
