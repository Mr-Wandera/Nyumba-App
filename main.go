package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"nyumba/models"
	"nyumba/templates"
)

var houses []models.House

func loadHouses() {
	file, err := os.Open("houses.json")
	if err != nil {
		log.Println("Error opening houses.json:", err)
		return
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	json.Unmarshal(byteValue, &houses)
}

func saveHouses() {
	byteValue, _ := json.MarshalIndent(houses, "", "  ")
	os.WriteFile("houses.json", byteValue, 0644)
}

func main() {
	loadHouses()

	// Static files
	fs := http.FileServer(http.Dir("uploads"))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	// Routes
	http.HandleFunc("/", landingHandler)
	http.HandleFunc("/explore", exploreHandler)
	http.HandleFunc("/landlord", landlordHandler)
	http.HandleFunc("/houses", housesAPIHandler)
	http.HandleFunc("/add-house", addHouseHandler)
	http.HandleFunc("/trigger-payment", paymentHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server running on http://0.0.0.0:%s\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func landingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Pass featured houses to landing
	featured := houses
	if len(featured) > 3 {
		featured = featured[:3]
	}
	// Corrected to pass models.House slice
	fmt.Fprint(w, templates.GetLandingHTML(featured))
}

func exploreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetExploreHTML())
}

func landlordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetLandlordHTML())
}

func housesAPIHandler(w http.ResponseWriter, r *http.Request) {
	neighborhood := r.URL.Query().Get("neighborhood")
	search := r.URL.Query().Get("search")

	filtered := houses
	if neighborhood != "" {
		// Changed to models.House to match the global slice
		var temp []models.House
		for _, h := range houses {
			if strings.Contains(strings.ToLower(h.Location), strings.ToLower(neighborhood)) {
				temp = append(temp, h)
			}
		}
		filtered = temp
	}

	if search != "" {
		// Changed to models.House to match the global slice
		var temp []models.House
		for _, h := range filtered {
			if strings.Contains(strings.ToLower(h.BuildingName), strings.ToLower(search)) ||
				strings.Contains(strings.ToLower(h.Location), strings.ToLower(search)) {
				temp = append(temp, h)
			}
		}
		filtered = temp
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func addHouseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/landlord", http.StatusSeeOther)
		return
	}

	// Handle both JSON and Form
	// Changed to models.House for type consistency
	var newHouse models.House
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		json.NewDecoder(r.Body).Decode(&newHouse)
	} else {
		r.ParseForm()
		newHouse.BuildingName = r.FormValue("building_name")
		newHouse.Location = r.FormValue("location")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		newHouse.Price = price
		beds, _ := strconv.Atoi(r.FormValue("bedrooms"))
		newHouse.Bedrooms = beds
		baths, _ := strconv.Atoi(r.FormValue("bathrooms"))
		newHouse.Bathrooms = baths
	}

	newHouse.ID = len(houses) + 1
	newHouse.ImageURLs = []string{fmt.Sprintf("https://picsum.photos/seed/%d/800/600", time.Now().UnixNano())}
	newHouse.IsPaid = false

	houses = append(houses, newHouse)
	saveHouses()

	if strings.Contains(r.Header.Get("Accept"), "application/json") {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": "Sanctuary published successfully!",
			"house":   newHouse,
		})
		return
	}

	http.Redirect(w, r, "/explore", http.StatusSeeOther)
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	// Simulate M-Pesa
	time.Sleep(1500 * time.Millisecond)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "M-Pesa STK Push sent! Please check your phone to complete the payment of KES 1,000.",
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetAuthHTML("Login"))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetAuthHTML("Sign Up"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetStaticHTML("About Us", "Nyumba is Kenya's premier sanctuary discovery platform. We eliminate the middleman to bring you closer to your next home."))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetStaticHTML("Contact Support", "Need help? Reach out to our 24/7 support team at support@nyumba.co.ke or call +254 700 000 000."))
}
