package handlers

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"nyumba/models"
	"nyumba/templates"
)

var users = []models.User{}
var houses = []models.House{}

const userFile = "users.json"
const houseFile = "houses.json"
const CookieName = "session_token"

const (
	shortCode    = "174379"
	mpesaAuthURL = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	mpesaPushURL = "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	callbackURL  = "https://nyumba-app.onrender.com/callback"
)

func LoadData() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ Notice: .env file not found (Normal on Render)")
	}
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

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/manifest.json" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name": "Nyumba", "short_name": "Nyumba", "start_url": "/", "display": "standalone", "background_color": "#0f172a", "theme_color": "#0f172a", "icons": [{"src": "https://cdn-icons-png.flaticon.com/512/25/25694.png", "sizes": "192x192", "type": "image/png"}]}`)
		return
	}

	// 🚦 TRAFFIC COP: If they visit the main domain ("/"), show the Landing Page
	if r.URL.Path == "/" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, templates.GetLandingHTML())
		return
	}

	// 🚦 TRAFFIC COP: If they visit anything else (like "/explore"), show the App
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	currentUser := getCurrentUser(r)
	isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay := "false", "", "", "none"

	if currentUser != nil {
		isLoggedIn, currentUsername = "true", currentUser.Username
		if currentUser.Role == "landlord" {
			landlordPanelDisplay = "block"
		}
		myHubButton = `<button onclick="openDashboard()" class="w-full bg-slate-800 hover:bg-slate-700 text-white font-bold py-3 rounded-xl flex items-center justify-center gap-2 transition mb-4 border border-white/10">🔐 My Unlocked Contacts</button>`
	}

	fmt.Fprint(w, templates.GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay))
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	errorMsg := ""

	if r.Method == http.MethodPost {
		username, password := r.FormValue("username"), r.FormValue("password")
		found := false
		for _, u := range users {
			if u.Username == username {
				found = true
				err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
				if err == nil {
					http.SetCookie(w, &http.Cookie{Name: CookieName, Value: username, Path: "/"})
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
			}
		}
		if !found || true {
			errorMsg = "Invalid Username or Password"
		}
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
					<label for="username" class="text-[10px] uppercase font-bold text-slate-500 tracking-wider ml-1">Username</label>
					<input id="username" name="username" type="text" placeholder="e.g. johndoe" required class="w-full bg-slate-900/50 border border-slate-700 focus:border-indigo-500 rounded-xl px-4 py-3 text-white outline-none transition">
				</div>
				<div>
					<label for="password" class="text-[10px] uppercase font-bold text-slate-500 tracking-wider ml-1">Password</label>
					<input id="password" name="password" type="password" placeholder="••••••••" required class="w-full bg-slate-900/50 border border-slate-700 focus:border-indigo-500 rounded-xl px-4 py-3 text-white outline-none transition">
				</div>
				<button class="w-full bg-gradient-to-r from-indigo-600 to-indigo-700 hover:from-indigo-500 hover:to-indigo-600 text-white font-bold py-4 rounded-xl shadow-lg shadow-indigo-500/30 transition transform active:scale-95 mt-4">Sign In</button>
			</form>
			
			<div class="mt-8 text-center border-t border-white/5 pt-6 flex flex-col gap-3">
				<p class="text-slate-400 text-sm">New here? <a href="/signup" class="text-indigo-400 font-bold hover:text-indigo-300 transition">Create Account</a></p>
				<a href="#" onclick="alert('Password reset link sent to your phone!')" class="text-slate-500 text-xs hover:text-slate-300 transition">Forgot Password?</a>
			</div>
		</div>
	</body></html>`, generateErrorHTML(errorMsg))

	fmt.Fprint(w, html)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username, password, phone, role := r.FormValue("username"), r.FormValue("password"), r.FormValue("phone"), r.FormValue("role")
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

	fmt.Fprint(w, `<!DOCTYPE html><html><head><title>Join • Nyumba</title><meta name="viewport" content="width=device-width"><link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;600;800&display=swap" rel="stylesheet"><script src="https://cdn.tailwindcss.com"></script>
	<style>body{font-family:'Outfit',sans-serif;background:#0f172a;color:#fff}.glass{background:rgba(30,41,59,0.7);backdrop-filter:blur(20px);border:1px solid rgba(255,255,255,0.1);box-shadow:0 25px 50px -12px rgba(0,0,0,0.5)}</style></head>
	<body class="h-screen flex items-center justify-center relative overflow-hidden">
		<div class="absolute top-0 left-0 w-full h-full overflow-hidden -z-10"><div class="absolute top-[-10%] right-[-10%] w-[40%] h-[40%] bg-purple-600/20 rounded-full blur-[100px]"></div></div>

		<div class="glass p-8 md:p-12 rounded-3xl w-full max-w-sm mx-4 relative">
			<div class="text-center mb-6">
				<h1 class="text-3xl font-bold text-white mb-2">Create Account</h1>
				<p class="text-xs text-slate-400">Join 500+ curated renters today.</p>
			</div>
			<form method="POST" class="space-y-3">
				<div>
					<label for="username" class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block ml-1">Username</label>
					<input id="username" name="username" type="text" placeholder="e.g. johndoe" required class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-3 text-white outline-none focus:border-indigo-500 transition">
				</div>
				<div>
					<label for="phone" class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block ml-1">M-Pesa Number</label>
					<input id="phone" name="phone" type="tel" placeholder="07XX..." required class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-3 text-white outline-none focus:border-indigo-500 transition">
				</div>
				<div>
					<label for="password" class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block ml-1">Password</label>
					<input id="password" name="password" type="password" placeholder="••••••••" required class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-3 text-white outline-none focus:border-indigo-500 transition">
				</div>
				<div class="relative pt-2">
					<select name="role" class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-3 text-white outline-none appearance-none cursor-pointer"><option value="renter">I want to Rent</option><option value="landlord">I am a Landlord</option></select><div class="absolute right-4 top-1/2 translate-y-[-20%] pointer-events-none text-slate-500">▼</div>
				</div>
				<button class="w-full bg-indigo-600 hover:bg-indigo-500 text-white font-bold py-4 rounded-xl shadow-lg mt-4 transition">Start Journey</button>
			</form>
			<div class="mt-6 text-center border-t border-white/5 pt-4">
				<a href="/login" class="text-slate-400 text-sm hover:text-white transition">Already have an account? Login</a>
			</div>
		</div>
	</body></html>`)
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

	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	var imageURLs []string
	for _, fileHeader := range r.MultipartForm.File["photos"] {
		file, _ := fileHeader.Open()
		ctx := context.Background()
		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "nyumba_app_houses"})
		file.Close()
		if err == nil {
			imageURLs = append(imageURLs, resp.SecureURL)
		}
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

func ServeMedia(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "."+r.URL.Path) }
func GetHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: CookieName, Value: "", Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func initiateSTKPush(phoneNumber, amount string) (string, error) {
	token, err := getAccessToken()
	if err != nil {
		return "", err
	}

	passkey := os.Getenv("MPESA_PASSKEY")
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
		return "", fmt.Errorf("failed")
	}
	return res["CheckoutRequestID"].(string), nil
}

func getAccessToken() (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", mpesaAuthURL, nil)

	key, secret := os.Getenv("MPESA_KEY"), os.Getenv("MPESA_SECRET")
	auth := base64.StdEncoding.EncodeToString([]byte(key + ":" + secret))
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

func generateErrorHTML(msg string) string {
	if msg == "" {
		return ""
	}
	return fmt.Sprintf(`<div class="bg-red-500/10 border border-red-500/50 text-red-200 px-4 py-3 rounded-xl mb-6 text-xs font-bold flex items-center gap-3 animate-pulse"><span class="text-lg">⚠️</span> %s</div>`, msg)
}
