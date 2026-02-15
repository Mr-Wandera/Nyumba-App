package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"nyumba/models"
	"nyumba/templates"
)

// Global Data Store
var users = []models.User{}
var houses = []models.House{}

const userFile = "users.json"
const houseFile = "houses.json"
const CookieName = "session_token"

// --- SAFARICOM CONFIG ---
const (
	consumerKey    = "COBGyH3dHvYrVjLKG0Znfh8RR1yAPeVbZ6hZitAwgvquIqhL"
	consumerSecret = "ovklACIWd4ZMihM4Vv28TAwgEBG8MywaI5FOnHahzIPXAG16CTCikL2RSSqT4cog"
	shortCode      = "174379"
	passkey        = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
	mpesaAuthURL   = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	mpesaPushURL   = "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	callbackURL    = "https://nyumba-app.onrender.com/callback"
)

// --- DATA HELPERS ---

func LoadData() {
	if _, err := os.Stat(userFile); err == nil {
		data, _ := os.ReadFile(userFile)
		json.Unmarshal(data, &users)
	}
	if _, err := os.Stat(houseFile); err == nil {
		data, _ := os.ReadFile(houseFile)
		json.Unmarshal(data, &houses)
	}
}

func saveData(filename string, data interface{}) {
	file, _ := json.MarshalIndent(data, "", " ")
	os.WriteFile(filename, file, 0644)
}

func formatPhoneNumber(phone string) string {
	phone = "" + phone
	if len(phone) > 0 && phone[0] == '0' {
		return "254" + phone[1:]
	}
	if len(phone) > 4 && phone[0] == '+' {
		return phone[1:]
	}
	return phone
}

func getCurrentUser(r *http.Request) *models.User {
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

// --- HANDLERS (PUBLIC) ---

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/manifest.json" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name": "Nyumba", "short_name": "Nyumba", "start_url": "/", "display": "standalone", "background_color": "#0f172a", "theme_color": "#0f172a", "icons": [{"src": "https://cdn-icons-png.flaticon.com/512/25/25694.png", "sizes": "192x192", "type": "image/png"}]}`)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	currentUser := getCurrentUser(r)
	isLoggedIn := "false"
	currentUsername := ""
	myHubButton := ""
	landlordPanelDisplay := "none"

	if currentUser != nil {
		isLoggedIn = "true"
		currentUsername = currentUser.Username
		if currentUser.Role == "landlord" {
			landlordPanelDisplay = "block"
		}
		myHubButton = `<button onclick="openDashboard()" class="w-full bg-slate-800 hover:bg-slate-700 text-white font-bold py-3 rounded-xl flex items-center justify-center gap-2 transition mb-4 border border-white/10">🔐 My Unlocked Contacts</button>`
	}

	fmt.Fprint(w, templates.GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		for _, u := range users {
			if u.Username == username {
				err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
				if err == nil {
					http.SetCookie(w, &http.Cookie{Name: CookieName, Value: username, Path: "/"})
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
			}
		}
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}
	fmt.Fprint(w, `<!DOCTYPE html><html><head><title>Login</title><meta name="viewport" content="width=device-width"><link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;600&display=swap" rel="stylesheet"><script src="https://cdn.tailwindcss.com"></script><style>body{font-family:'Outfit',sans-serif;background:#0b0f19;color:#fff}.glass{background:rgba(30,41,59,0.4);backdrop-filter:blur(16px);border:1px solid rgba(255,255,255,0.05)}</style></head><body class="h-screen flex items-center justify-center"><div class="glass p-10 rounded-3xl w-full max-w-sm"><h1 class="text-3xl font-bold mb-6 text-center">Login</h1><form method="POST"><input name="username" placeholder="Username" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><input name="password" type="password" placeholder="Password" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><button class="w-full bg-indigo-600 py-3 rounded-xl font-bold">Sign In</button></form><a href="/signup" class="block text-center mt-6 text-slate-400 text-sm">Create Account</a></div></body></html>`)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		phone := r.FormValue("phone")
		role := r.FormValue("role")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Server Error", 500)
			return
		}

		newUser := models.User{Username: username, Password: string(hashedPassword), Phone: phone, Role: role}
		users = append(users, newUser)
		saveData(userFile, users)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	fmt.Fprint(w, `<!DOCTYPE html><html><head><title>Join</title><meta name="viewport" content="width=device-width"><link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;600&display=swap" rel="stylesheet"><script src="https://cdn.tailwindcss.com"></script><style>body{font-family:'Outfit',sans-serif;background:#0b0f19;color:#fff}.glass{background:rgba(30,41,59,0.4);backdrop-filter:blur(16px);border:1px solid rgba(255,255,255,0.05)}</style></head><body class="h-screen flex items-center justify-center"><div class="glass p-10 rounded-3xl w-full max-w-sm"><h1 class="text-3xl font-bold mb-6 text-center">Join</h1><form method="POST"><input name="username" placeholder="Username" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><input name="phone" placeholder="Phone" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><input name="password" type="password" placeholder="Password" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><select name="role" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><option value="renter">Renter</option><option value="landlord">Landlord</option></select><button class="w-full bg-indigo-600 py-3 rounded-xl font-bold">Create Account</button></form><a href="/login" class="block text-center mt-6 text-slate-400 text-sm">Login Here</a></div></body></html>`)
}

func UploadHouse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}
	r.ParseMultipartForm(20 << 20)
	currentUser := getCurrentUser(r)
	if currentUser == nil || currentUser.Role != "landlord" {
		http.Error(w, "Unauthorized", 401)
		return
	}

	var imageURLs []string
	for _, fileHeader := range r.MultipartForm.File["photos"] {
		file, _ := fileHeader.Open()
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fileHeader.Filename)
		dst, _ := os.Create("uploads/" + filename)
		io.Copy(dst, file)
		dst.Close()
		file.Close()
		imageURLs = append(imageURLs, "/uploads/"+filename)
	}

	p, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	u, _ := strconv.ParseFloat(r.FormValue("utilities"), 64)
	newHouse := models.House{
		ID: len(houses) + 1, BuildingName: r.FormValue("building_name"), Location: r.FormValue("location"),
		Type: r.FormValue("type"), Price: p, Utilities: u, Details: r.FormValue("details"),
		ImageURLs: imageURLs, Phone: currentUser.Phone, Owner: currentUser.Username, IsBooked: false, MapURL: r.FormValue("map_url"),
	}
	houses = append(houses, newHouse)
	saveData(houseFile, houses)
	w.WriteHeader(http.StatusCreated)
}

func DeleteHouseHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	newHouses := []models.House{}
	for _, h := range houses {
		if h.ID != id {
			newHouses = append(newHouses, h)
		}
	}
	houses = newHouses
	saveData(houseFile, houses)
	w.WriteHeader(200)
}

// PAY HANDLER (Updates to request payment & Save ID)
func PayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	rawPhone := r.URL.Query().Get("phone")
	phone := formatPhoneNumber(rawPhone)

	var selectedHouse *models.House
	for i, h := range houses {
		if h.ID == id {
			selectedHouse = &houses[i]
			break
		}
	}

	if selectedHouse == nil {
		w.WriteHeader(404)
		fmt.Fprint(w, `{"ResponseCode": "1", "CustomerMessage": "House Not Found"}`)
		return
	}

	checkoutID, err := initiateSTKPush(phone, "1")
	if err != nil {
		fmt.Fprintf(w, `{"ResponseCode": "1", "CustomerMessage": "M-Pesa Error: %s"}`, err.Error())
		return
	}

	selectedHouse.CheckoutRequestID = checkoutID
	saveData(houseFile, houses)
	fmt.Fprint(w, `{"ResponseCode": "0", "CustomerMessage": "Request Sent"}`)
}

// CALLBACK HANDLER (Listens for Success from Safaricom)
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	var callback models.MpesaCallback
	json.NewDecoder(r.Body).Decode(&callback)

	if callback.Body.StkCallback.ResultCode == 0 {
		targetID := callback.Body.StkCallback.CheckoutRequestID
		for i, h := range houses {
			if h.CheckoutRequestID == targetID {
				houses[i].IsBooked = true
				houses[i].CheckoutRequestID = ""
				saveData(houseFile, houses)
				break
			}
		}
	}
	w.WriteHeader(200)
}

// ServeMedia (Restored!)
func ServeMedia(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "."+r.URL.Path) }

func GetHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: CookieName, Value: "", Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// --- MPESA HELPERS ---

func initiateSTKPush(phoneNumber, amount string) (string, error) {
	token, err := getAccessToken()
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Format("20060102150405")
	password := base64.StdEncoding.EncodeToString([]byte(shortCode + passkey + timestamp))
	headers := map[string]string{"Authorization": "Bearer " + token, "Content-Type": "application/json"}
	payload := map[string]string{
		"BusinessShortCode": shortCode, "Password": password, "Timestamp": timestamp,
		"TransactionType": "CustomerPayBillOnline", "Amount": amount, "PartyA": phoneNumber,
		"PartyB": shortCode, "PhoneNumber": phoneNumber, "CallBackURL": callbackURL, "AccountReference": "NyumbaApp", "TransactionDesc": "Viewing Fee",
	}
	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", mpesaPushURL, bytes.NewBuffer(jsonPayload))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	if res["ResponseCode"] != "0" {
		return "", fmt.Errorf("failed") // Simplified error for brevity
	}
	return res["CheckoutRequestID"].(string), nil
}

func getAccessToken() (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", mpesaAuthURL, nil)
	auth := base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret))
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to get token")
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["access_token"].(string), nil
}
