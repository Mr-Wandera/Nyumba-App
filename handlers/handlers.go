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
	if r.Method == http.MethodPost {
		// 1. Parse multipart form (max 10MB)
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "File too large", http.StatusBadRequest)
			return
		}

		// 2. Handle Image Upload
		file, header, err := r.FormFile("property_photo")
		imagePath := "/uploads/default.jpg"
		if err == nil {
			defer file.Close()
			os.MkdirAll("./uploads", os.ModePerm)
			filename := filepath.Base(header.Filename)
			dstPath := filepath.Join("./uploads", filename)
			dst, _ := os.Create(dstPath)
			defer dst.Close()
			io.Copy(dst, file)
			imagePath = "/uploads/" + filename
		}

		// 3. Create House entry
		newHouse := models.House{
			ID:           len(Houses) + 1,
			BuildingName: r.FormValue("building_name"),
			Location:     r.FormValue("location"),
			MapLink:      r.FormValue("map_link"),
			Price:        7500.0,
			ImageURLs:    []string{imagePath},
		}
		Houses = append(Houses, newHouse)
		http.Redirect(w, r, "/explore", http.StatusSeeOther)
	}
}

func GetHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Houses)
}

func ExploreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetHTML("Abdul")+templates.GetScripts())
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetLandingHTML())
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
	Houses = append(Houses, models.House{ID: 1, BuildingName: "Base", Location: "Thika", Price: 7500})
}
