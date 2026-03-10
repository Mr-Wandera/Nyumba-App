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

// HomePage serves the landing page
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, templates.GetLandingHTML())
}

// ExploreHandler serves the main search interface
func ExploreHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, templates.GetHTML("User"))
}

// GetHouses returns the list of houses as JSON for the frontend
func GetHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Houses)
}

// TriggerPaymentHandler connects the "Unlock" button to M-Pesa
func TriggerPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		HouseID int `json:"house_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// This is where you call your Safaricom STK Push function
	fmt.Printf("M-Pesa STK Push initiated for House ID: %d\n", req.HouseID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Check your phone for the M-Pesa PIN prompt",
	})
}

// MpesaCallbackHandler handles the result sent back by Safaricom
func MpesaCallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing M-Pesa Callback...")
	w.WriteHeader(http.StatusOK)
}

// AddHouseHandler handles the 'Publish Listing' form and file uploads
func AddHouseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB limit
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("property_photo")
	imagePath := "/uploads/default.jpg"
	if err == nil {
		defer file.Close()
		os.MkdirAll("./uploads", os.ModePerm)
		filename := fmt.Sprintf("%d_%s", len(Houses)+1, filepath.Base(header.Filename))
		dstPath := filepath.Join("./uploads", filename)
		dst, _ := os.Create(dstPath)
		defer dst.Close()
		io.Copy(dst, file)
		imagePath = "/uploads/" + filename
	}

	price := 15000.0
	fmt.Sscanf(r.FormValue("price"), "%f", &price)

	newHouse := models.House{
		ID:            len(Houses) + 1,
		BuildingName:  r.FormValue("building_name"),
		Location:      r.FormValue("location"),
		Price:         price,
		ImageURLs:     []string{imagePath},
		IsPaid:        false,
		LandlordPhone: "+254712345678",
	}

	Houses = append(Houses, newHouse)
	saveHouses()

	http.Redirect(w, r, "/explore", http.StatusSeeOther)
}

// LoadData reads the JSON file into memory
func LoadData(filename string, target interface{}) {
	file, err := os.ReadFile(filename)
	if err == nil {
		json.Unmarshal(file, target)
	}
}

// saveHouses persists current memory state back to the JSON file
func saveHouses() {
	data, _ := json.MarshalIndent(Houses, "", "  ")
	os.WriteFile("houses.json", data, 0644)
}

// SeedHouses adds a default listing if the file is empty
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
		LandlordPhone: "+254712345678",
	})
	saveHouses()
}
