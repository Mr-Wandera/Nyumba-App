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
)

var houses []models.House
var users []models.User

const houseFile = "houses.json"
const userFile = "users.json"

// SeedHouses populates the app with initial professional listings
func SeedHouses() {
	if len(houses) > 0 {
		return
	}
	seedData := []models.House{
		{
			ID: 1, BuildingName: "Lavington Heights", Location: "Lavington",
			Type: "Two Bedroom", Price: 65000, Details: "Modern apartment with a panoramic view. High-speed lift and 24/7 security.",
			ImageURLs: []string{"https://res.cloudinary.com/dqkqyhou9/image/upload/v1740244597/lavington_main.jpg"},
			Phone:     "0712345678",
		},
		{
			ID: 2, BuildingName: "Thika Greens Villa", Location: "Thika",
			Type: "Three Bedroom", Price: 85000, Details: "Spacious family home near MKU. Quiet neighborhood with a private garden.",
			ImageURLs: []string{"https://res.cloudinary.com/dqkqyhou9/image/upload/v1740244600/thika_greens.jpg"},
			Phone:     "0722000111",
		},
	}
	houses = append(houses, seedData...)
	// saveData(houseFile, houses) // Ensure your saveData function is available
}

func getAccessToken() (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials", nil)
	key, secret := os.Getenv("MPESA_KEY"), os.Getenv("MPESA_SECRET")
	auth := base64.StdEncoding.EncodeToString([]byte(key + ":" + secret))
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("🚨 SAFARICOM ERROR:", string(body))
		return "", fmt.Errorf("failed to get token")
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["access_token"].(string), nil
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	errorMsg := ""
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		for _, u := range users {
			if u.Username == username && u.Password == password {
				// Set session logic here
				http.Redirect(w, r, "/explore", http.StatusSeeOther)
				return
			}
		}
		errorMsg = "Invalid username or password"
	}

	html := fmt.Sprintf(`<!DOCTYPE html><html><head><title>Login • Nyumba</title><meta name="viewport" content="width=device-width"><link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;600;800&display=swap" rel="stylesheet"><script src="https://cdn.tailwindcss.com"></script>
    <style>body{font-family:'Outfit',sans-serif;background:#0f172a;color:#fff}.glass{background:rgba(30,41,59,0.7);backdrop-filter:blur(20px);border:1px solid rgba(255,255,255,0.1);box-shadow:0 25px 50px -12px rgba(0,0,0,0.5)}</style></head>
    <body class="h-screen flex items-center justify-center relative overflow-hidden">
        <div class="absolute top-0 left-0 w-full h-full overflow-hidden -z-10"><div class="absolute top-[-10%%] left-[-10%%] w-[40%%] h-[40%%] bg-indigo-600/20 rounded-full blur-[100px]"></div><div class="absolute bottom-[-10%%] right-[-10%%] w-[40%%] h-[40%%] bg-emerald-500/10 rounded-full blur-[100px]"></div></div>
        <div class="glass p-8 md:p-12 rounded-3xl w-full max-w-sm mx-4 relative transform transition hover:scale-[1.01] duration-500">
            <div class="text-center mb-8">
                <h1 class="text-4xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300 mb-2">Nyumba.</h1>
                <p class="text-xs text-slate-400 font-medium tracking-widest uppercase">Welcome Back</p>
            </div>
            %s 
            <form method="POST" class="space-y-4">
                <div>
                    <label class="text-[10px] uppercase font-bold text-slate-500 ml-1">Username</label>
                    <input name="username" type="text" placeholder="e.g. Abdul" required class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-3 text-white outline-none focus:border-indigo-500 transition">
                </div>
                <div>
                    <label class="text-[10px] uppercase font-bold text-slate-500 ml-1">Password</label>
                    <input name="password" type="password" placeholder="••••••••" required class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-3 text-white outline-none focus:border-indigo-500 transition">
                </div>
                <button class="w-full bg-gradient-to-r from-indigo-600 to-indigo-700 text-white font-bold py-4 rounded-xl shadow-lg mt-4">Sign In</button>
            </form>
        </div>
    </body></html>`, generateErrorHTML(errorMsg))
	fmt.Fprint(w, html)
}

func generateErrorHTML(msg string) string {
	if msg == "" {
		return ""
	}
	return fmt.Sprintf(`<div class="bg-red-500/10 border border-red-500/50 text-red-400 text-xs py-3 px-4 rounded-xl mb-6 font-bold text-center">⚠️ %s</div>`, msg)
}
func ExploreHandler(w http.ResponseWriter, r *http.Request) {
	// Assume session check happens here
	isLoggedIn := "true"
	currentUsername := "Abdul" // cite: User Summary
	myHubButton := `<button onclick="openDashboard()" class="w-full bg-indigo-600/20 text-indigo-400 border border-indigo-500/30 py-3 rounded-xl text-sm font-bold mb-4">My Unlocked Sanctuaries</button>`
	landlordPanelDisplay := "block"

	// COMBINE THE TWO TEMPLATES TO FIX SYNTAX ERRORS
	htmlBody := templates.GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay)
	scripts := templates.GetScripts(isLoggedIn == "true", currentUsername)

	fmt.Fprint(w, htmlBody+scripts)
}

func GetHousesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}
