package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// --- CONFIGURATION ---
const CookieName = "nyumba_session"
const houseFile = "nyumba.json"
const userFile = "users.json"

// Global Variables
var houses = []House{}
var users = []User{}

// --- HELPERS ---
func saveData(filename string, data interface{}) {
	fileData, _ := json.MarshalIndent(data, "", "  ")
	ioutil.WriteFile(filename, fileData, 0644)
}

func loadData() {
	if _, err := os.Stat(houseFile); err == nil {
		data, _ := ioutil.ReadFile(houseFile)
		json.Unmarshal(data, &houses)
	}
	if _, err := os.Stat(userFile); err == nil {
		data, _ := ioutil.ReadFile(userFile)
		json.Unmarshal(data, &users)
	}
}

func getCurrentUser(r *http.Request) *User {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return nil
	}
	username := cookie.Value
	for _, u := range users {
		if u.Username == username {
			return &u
		}
	}
	return nil
}

func main() {
	// 1. Create uploads folder if it doesn't exist
	os.MkdirAll("uploads", os.ModePerm)

	// 2. Load data from JSON files
	loadData()

	// 3. Setup Routes (The logic is in handlers.go)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/houses", getHouses)
	http.HandleFunc("/houses/upload", uploadHouseHandler)
	http.HandleFunc("/houses/delete", deleteHouseHandler)

	// 4. Serve images from the uploads folder
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	// 5. START SERVER (Fixed for Render!)
	// Render gives us a PORT, or we use 8082 if testing on laptop
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	fmt.Println("Server running on port " + port + " ...")
	http.ListenAndServe(":"+port, nil)
}
