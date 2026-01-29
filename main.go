package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 1. Create uploads folder
	os.MkdirAll("uploads", os.ModePerm)

	// 2. Load data
	loadData()

	// 3. Setup Routes
	http.HandleFunc("/", homePage)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/houses", getHouses)
	http.HandleFunc("/houses/upload", uploadHouseHandler)
	http.HandleFunc("/houses/delete", deleteHouseHandler)

	// 4. Serve images
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	// 5. START SERVER
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	fmt.Println("Server running on port " + port + " ...")
	http.ListenAndServe(":"+port, nil)
}

// --- HELPERS ---

func getCurrentUser(r *http.Request) *User {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return nil
	}
	for _, u := range users {
		if u.Username == cookie.Value {
			return &u
		}
	}
	return nil
}

func saveData(filename string, data interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	file, _ := json.MarshalIndent(data, "", " ")
	os.WriteFile(filename, file, 0644)
}

func loadData() {
	if _, err := os.Stat(userFile); err == nil {
		file, _ := os.Open(userFile)
		byteValue, _ := io.ReadAll(file)
		json.Unmarshal(byteValue, &users)
		file.Close()
	}
	if _, err := os.Stat(houseFile); err == nil {
		file, _ := os.Open(houseFile)
		byteValue, _ := io.ReadAll(file)
		json.Unmarshal(byteValue, &houses)
		file.Close()
	}
}
