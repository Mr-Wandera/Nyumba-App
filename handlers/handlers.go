package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nyumba/models"
	"nyumba/templates"
	"os"
	"path/filepath"
	"time"
)

// M-Pesa API response structure
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

var Houses []models.House

func AddHouseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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

func GetMpesaToken(consumerKey, consumerSecret string) (string, error) {
	url := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret))
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("mpesa auth failed: %d", res.StatusCode)
	}

	var tokenRes TokenResponse
	if err := json.NewDecoder(res.Body).Decode(&tokenRes); err != nil {
		return "", err
	}

	return tokenRes.AccessToken, nil
}
