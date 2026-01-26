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
	os.MkdirAll("uploads", os.ModePerm)
	loadData()

	// These functions are found in handlers.go
	http.HandleFunc("/", homePage)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/houses", getHouses)
	http.HandleFunc("/houses/upload", uploadHouseHandler)
	http.HandleFunc("/houses/delete", deleteHouseHandler)

	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	fmt.Println("Server running on http://localhost:8082 ...")
	http.ListenAndServe(":8082", nil)
}
