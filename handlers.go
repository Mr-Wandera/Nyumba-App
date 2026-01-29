package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// --- HANDLERS ---

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// 1. Get Form Data
		username := r.FormValue("username")
		password := r.FormValue("password")
		phone := r.FormValue("phone")
		role := r.FormValue("role")

		// 2. Simple Validation
		if username == "" || password == "" || phone == "" {
			http.Error(w, "All fields required", http.StatusBadRequest)
			return
		}

		// 3. Create User
		newUser := User{
			Username: username,
			Password: password,
			Phone:    phone,
			Role:     role,
		}

		// 4. Save to Database
		users = append(users, newUser)
		saveData(userFile, users)

		// 5. Redirect to Login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Join • Nyumba</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet">
		<script src="https://cdn.tailwindcss.com"></script>
		<style>
			body { font-family: 'Outfit', sans-serif; background: #0f172a; color: #f8fafc; }
			
			.glass-card {
				background: rgba(30, 41, 59, 0.7);
				border: 1px solid rgba(255, 255, 255, 0.1);
				box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
			}
		</style>
	</head>
	<body class="h-screen w-full flex items-center justify-center relative overflow-hidden bg-slate-900">
		
		<div class="absolute inset-0 bg-gradient-to-br from-slate-900 via-slate-900 to-indigo-900/20"></div>

		<div class="glass-card p-8 rounded-3xl w-full max-w-md mx-4 relative z-10">
			<div class="text-center mb-6">
				<h1 class="text-3xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300 mb-2">Join Nyumba.</h1>
				<p class="text-slate-400 text-sm font-medium">Start your hunt today.</p>
			</div>

			<form method="POST" action="/signup" class="space-y-4">
				<div>
					<label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block">Username</label>
					<input name="username" type="text" placeholder="Choose a name" required 
						class="w-full bg-slate-800/50 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-indigo-500/50 outline-none transition placeholder-slate-600">
				</div>
				
				<div>
					<label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block">Phone Number</label>
					<input name="phone" type="text" placeholder="2547..." required 
						class="w-full bg-slate-800/50 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-indigo-500/50 outline-none transition placeholder-slate-600">
				</div>

				<div>
					<label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block">Password</label>
					<input name="password" type="password" placeholder="Create password" required 
						class="w-full bg-slate-800/50 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-indigo-500/50 outline-none transition placeholder-slate-600">
				</div>

				<div>
					<label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block">I am a...</label>
					<div class="relative">
						<select name="role" class="w-full bg-slate-800/50 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-indigo-500/50 outline-none transition appearance-none cursor-pointer">
							<option value="renter">🏠 Renter (Looking for a house)</option>
							<option value="landlord">🔑 Landlord (Listing a house)</option>
						</select>
						<div class="absolute right-4 top-3.5 text-slate-500 pointer-events-none">▼</div>
					</div>
				</div>

				<button type="submit" class="w-full mt-4 bg-indigo-600 hover:bg-indigo-500 text-white font-bold py-3.5 rounded-xl shadow-lg shadow-indigo-500/20 transition transform hover:-translate-y-0.5">
					Create Account
				</button>
			</form>

			<div class="mt-6 text-center border-t border-white/5 pt-4">
				<p class="text-slate-500 text-xs mb-2">Already have an account?</p>
				<a href="/login" class="text-indigo-400 text-sm font-bold hover:text-indigo-300 transition">Login Here</a>
			</div>
		</div>

	</body>
	</html>`
	fmt.Fprint(w, html)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		for _, u := range users {
			if u.Username == username && u.Password == password {
				http.SetCookie(w, &http.Cookie{
					Name:  CookieName,
					Value: username,
					Path:  "/",
				})
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}
		http.Error(w, "Invalid Credentials. Try again.", http.StatusUnauthorized)
		return
	}

	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Login • Nyumba</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet">
		<script src="https://cdn.tailwindcss.com"></script>
		<style>
			body { font-family: 'Outfit', sans-serif; background: #0b0f19; color: #f8fafc; }
			
			/* Animations */
			@keyframes float { 0% { transform: translateY(0px); } 50% { transform: translateY(-20px); } 100% { transform: translateY(0px); } }
			.animate-float { animation: float 6s ease-in-out infinite; }
			
			.glass-card {
				background: rgba(30, 41, 59, 0.4);
				backdrop-filter: blur(16px);
				border: 1px solid rgba(255, 255, 255, 0.05);
				box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
			}
		</style>
	</head>
	<body class="h-screen w-full flex items-center justify-center relative overflow-hidden">
		
		<div class="absolute top-0 left-0 w-full h-full overflow-hidden pointer-events-none">
			<div class="absolute top-[-10%] left-[-10%] w-[500px] h-[500px] bg-indigo-600 rounded-full mix-blend-screen filter blur-[120px] opacity-30 animate-float"></div>
			<div class="absolute bottom-[-10%] right-[-10%] w-[500px] h-[500px] bg-purple-600 rounded-full mix-blend-screen filter blur-[120px] opacity-30 animate-float" style="animation-delay: 2s"></div>
		</div>

		<div class="glass-card p-10 rounded-3xl w-full max-w-sm mx-4 relative z-10 border-t border-white/10">
			<div class="text-center mb-8">
				<h1 class="text-4xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300 mb-2">Nyumba.</h1>
				<p class="text-slate-400 text-sm font-medium">Welcome back, Hunter.</p>
			</div>

			<form method="POST" action="/login" class="space-y-5">
				<div>
					<label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-2 block">Username</label>
					<input name="username" type="text" placeholder="username" required 
						class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-indigo-500/50 focus:border-transparent outline-none transition placeholder-slate-600">
				</div>
				
				<div>
					<label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-2 block">Password</label>
					<input name="password" type="password" placeholder="••••••••" required 
						class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-indigo-500/50 focus:border-transparent outline-none transition placeholder-slate-600">
				</div>

				<button type="submit" class="w-full mt-2 bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-500 hover:to-purple-500 text-white font-bold py-3.5 rounded-xl shadow-lg shadow-indigo-500/20 transition transform hover:-translate-y-0.5">
					Sign In
				</button>
			</form>

			<div class="mt-8 text-center border-t border-white/5 pt-6">
				<p class="text-slate-500 text-xs mb-2">New to Nyumba?</p>
				<a href="/signup" class="text-indigo-400 text-sm font-bold hover:text-indigo-300 transition">Create an Account</a>
			</div>
		</div>

	</body>
	</html>`
	fmt.Fprint(w, html)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: CookieName, Value: "", Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func payHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	phone := r.URL.Query().Get("phone")

	if phone == "" {
		http.Error(w, "Phone number required", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(idStr)
	for i, h := range houses {
		if h.ID == id {
			houses[i].IsBooked = true
			houses[i].TenantPhone = phone // 👈 NEW: Save the tenant's number!
			break
		}
	}
	saveData(houseFile, houses)

	// Trigger M-Pesa (1 KES for testing)
	response, err := initiateSTKPush(phone, "1")

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
	} else {
		fmt.Fprint(w, response)
	}
}

func getHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}

// --- UPDATED HANDLERS ---

func uploadHouseHandler(w http.ResponseWriter, r *http.Request) {
	user := getCurrentUser(r)
	if user == nil || user.Role != "landlord" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 1. Parse Form
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		return
	}

	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	utils, _ := strconv.ParseFloat(r.FormValue("utilities"), 64)
	loc := r.FormValue("location")
	houseType := r.FormValue("type") // 👈 NEW: Get the house type
	details := r.FormValue("details")
	var tags []string
	json.Unmarshal([]byte(r.FormValue("tags")), &tags)

	// 2. Handle Images
	var imagePaths []string
	files := r.MultipartForm.File["photos"]

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fileHeader.Filename)
		dstPath := filepath.Join("uploads", filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			continue
		}
		defer dst.Close()
		io.Copy(dst, file)
		imagePaths = append(imagePaths, "/uploads/"+filename)
	}

	// 3. Save House (With Type)
	newHouse := House{
		ID:        len(houses) + 1,
		Location:  loc,
		Price:     price,
		Type:      houseType, // 👈 Save it
		Utilities: utils,
		Details:   details,
		Tags:      tags,
		ImageURLs: imagePaths,
		Owner:     user.Username,
		Phone:     user.Phone,
		IsBooked:  false,
	}
	houses = append(houses, newHouse)
	saveData(houseFile, houses)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newHouse)
}

func deleteHouseHandler(w http.ResponseWriter, r *http.Request) {
	// (Keep this the same as before)
	user := getCurrentUser(r)
	if user == nil || user.Role != "landlord" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	newHouses := []House{}
	for _, h := range houses {
		if h.ID != id {
			newHouses = append(newHouses, h)
		}
	}
	houses = newHouses
	saveData(houseFile, houses)
	w.WriteHeader(http.StatusOK)
}

// --- PASTE THIS INSIDE handlers.go (Replacing the old homePage) ---

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	currentUser := getCurrentUser(r)
	isLoggedIn := "false"
	currentUsername := ""
	currentUserPhone := ""

	welcomeMsg := "Welcome"
	navLinks := `<a href="/login" class="text-sm font-medium text-slate-300 hover:text-white transition">Login</a>`
	landlordPanelDisplay := "none"

	if currentUser != nil {
		isLoggedIn = "true"
		currentUsername = currentUser.Username
		currentUserPhone = currentUser.Phone
		welcomeMsg = "Hi, " + currentUser.Username
		navLinks = `<a href="/logout" class="text-sm font-bold text-red-400 border border-red-500/30 px-3 py-1 rounded-full hover:bg-red-500/10 transition">Logout</a>`
		if currentUser.Role == "landlord" {
			landlordPanelDisplay = "block"
		}
	}

	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Nyumba Discovery</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet">
		<script src="https://cdn.tailwindcss.com"></script>
		<style>
			/* Changed background to Slate-900 (Lighter than before) */
			body { font-family: 'Outfit', sans-serif; background: #0f172a; color: #f8fafc; }
			::-webkit-scrollbar { width: 6px; }
			::-webkit-scrollbar-track { background: #0f172a; }
			::-webkit-scrollbar-thumb { background: #334155; border-radius: 3px; }
			/* Simplified Glass - No Strong Blur to prevent layering issues */
			.glass { background: rgba(30, 41, 59, 0.7); border: 1px solid rgba(255, 255, 255, 0.1); }
			.glass-sidebar { background: #1e293b; border-right: 1px solid rgba(255, 255, 255, 0.05); }
		</style>
	</head>
	<body class="h-screen flex overflow-hidden">
		
		<aside class="w-80 flex-shrink-0 glass-sidebar flex flex-col h-full relative z-20">
			<div class="p-8 pb-4">
				<h1 class="text-4xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300">Nyumba.</h1>
				<p class="text-xs text-slate-500 font-medium tracking-widest uppercase mt-2">Curated Living</p>
			</div>

			<div class="px-6 py-4 space-y-6 flex-1 overflow-y-auto">
				<div style="display: ` + landlordPanelDisplay + `;" class="glass rounded-2xl p-5 mb-8">
					<h3 class="text-xs font-bold text-indigo-400 uppercase tracking-wider mb-4">Landlord Mode</h3>
					<div class="space-y-3">
						<input id="loc" type="text" placeholder="Location (e.g. Juja)" class="w-full bg-slate-900 border border-slate-700 rounded-lg px-3 py-2 text-sm text-white outline-none">
						<input id="map_url" type="text" placeholder="📍 Google Maps Link" class="w-full bg-slate-900 border border-indigo-500/30 rounded-lg px-3 py-2 text-sm text-indigo-300 outline-none">
						<select id="type" class="w-full bg-slate-900 border border-slate-700 rounded-lg px-3 py-2 text-sm outline-none">
							<option>Bedsitter</option><option>One Bedroom</option><option>Two Bedroom</option><option>Studio</option>
						</select>
						<div class="grid grid-cols-2 gap-2">
							<input id="price" type="number" placeholder="Rent" class="w-full bg-slate-900 border border-slate-700 rounded-lg px-3 py-2 text-sm outline-none">
							<input id="utils" type="number" placeholder="Bills" class="w-full bg-slate-900 border border-slate-700 rounded-lg px-3 py-2 text-sm outline-none">
						</div>
						<input id="photos" type="file" multiple class="text-xs text-slate-500">
						<textarea id="details" placeholder="Description..." class="w-full bg-slate-900 border border-slate-700 rounded-lg px-3 py-2 text-sm h-16 resize-none outline-none"></textarea>
						<button onclick="uploadHouse()" class="w-full bg-indigo-600 hover:bg-indigo-500 text-white font-bold py-2 rounded-lg text-sm transition">Post Listing</button>
					</div>
				</div>

				<div class="space-y-4">
					<div>
						<label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block">Where to?</label>
						<input id="searchLoc" onkeyup="fetchHouses()" type="text" placeholder="Try 'Kileleshwa'..." 
							class="w-full bg-slate-800 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-indigo-500 outline-none">
					</div>
					<div>
						<label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block">Max Budget</label>
						<input id="searchPrice" onkeyup="fetchHouses()" type="number" placeholder="Any Price" 
							class="w-full bg-slate-800 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-emerald-500 outline-none">
					</div>
				</div>
			</div>

			<div class="p-6 border-t border-white/5 flex items-center justify-between bg-[#1e293b]">
				<div class="flex items-center gap-3">
					<div class="w-8 h-8 rounded-full bg-indigo-500 flex items-center justify-center text-xs font-bold">👤</div>
					<div class="text-sm"><div class="font-bold text-white">` + currentUsername + `</div></div>
				</div>
				` + navLinks + `
			</div>
		</aside>

		<main class="flex-1 h-full overflow-y-auto bg-slate-900 relative z-10">
			<div class="p-8 max-w-[1600px] mx-auto">
				<header class="flex justify-between items-end mb-8">
					<div>
						<h2 class="text-3xl font-light text-white">Discover <span class="font-bold text-indigo-400">Sanctuary</span></h2>
						<p class="text-slate-400 mt-1">Pay the service fee to unlock locations instantly.</p>
					</div>
					<button onclick="alert('Full Map View coming in v2.0!')" class="bg-white text-black px-4 py-2 rounded-full text-sm font-bold hover:bg-slate-200 transition">View Map 🗺️</button>
				</header>
				<div id="results-area" class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6 auto-rows-[minmax(180px,auto)] pb-20"></div>
			</div>
		</main>

		<div id="toast" class="fixed top-6 left-1/2 -translate-x-1/2 bg-indigo-600 px-6 py-3 rounded-full text-sm font-bold text-white shadow-2xl translate-y-[-150%] transition-transform duration-500 z-50 flex items-center gap-2">
			<span class="text-lg">✨</span> <span id="toast-msg">Notification</span>
		</div>

		<script>
			const isLoggedIn = ` + isLoggedIn + `;
			const currentUsername = "` + currentUsername + `";
			const currentUserPhone = "` + currentUserPhone + `";

			document.addEventListener("DOMContentLoaded", () => fetchHouses());

			function showToast(msg) {
				const t = document.getElementById("toast"); document.getElementById("toast-msg").innerText = msg;
				t.classList.remove("translate-y-[-150%]"); setTimeout(() => t.classList.add("translate-y-[-150%]"), 3000);
			}

			function fetchHouses() {
				const sLoc = document.getElementById('searchLoc').value.toLowerCase();
				const sPrice = document.getElementById('searchPrice').value;

				fetch('/houses').then(res => res.json()).then(data => {
					const container = document.getElementById('results-area');
					container.innerHTML = "";
					
					let filtered = data.filter(h => {
						if(sLoc && !h.location.toLowerCase().includes(sLoc)) return false;
						if(sPrice && h.price > parseFloat(sPrice)) return false;
						return true;
					});

					if (filtered.length === 0) { container.innerHTML = "<div class='col-span-full text-center text-slate-500 py-20'>No sanctuaries found.</div>"; return; }

					filtered.forEach((h, index) => {
						const isOwner = (h.owner === currentUsername);
						const didIPay = (h.tenant_phone === currentUserPhone && currentUserPhone !== "");
						
						let gridClass = (index === 0) ? "md:col-span-2 row-span-2" : "";
						let statusBadge, opacityClass, actionBtn;
						let imageSrc = (h.image_urls && h.image_urls.length > 0) ? h.image_urls[0] : 'https://via.placeholder.com/600x400?text=No+Image';

						if (h.is_booked) {
							if (isOwner) {
								statusBadge = '<span class="absolute top-4 right-4 bg-indigo-600 text-white text-[10px] font-bold px-3 py-1 rounded-full z-20">Paid by: ' + h.tenant_phone + '</span>';
								opacityClass = "border-2 border-indigo-500";
								actionBtn = '<button onclick="deleteHouse(' + h.id + ')" class="mt-4 w-full py-3 rounded-xl bg-slate-800 text-red-400 text-xs font-bold">Delete Listing</button>';
							} else if (didIPay) {
								statusBadge = '<span class="absolute top-4 right-4 bg-emerald-500 text-white text-[10px] font-bold px-3 py-1 rounded-full z-20 shadow-xl">UNLOCKED ✅</span>';
								opacityClass = "border-2 border-emerald-500 shadow-xl";
								let mapLink = h.map_url ? h.map_url : "https://www.google.com/maps/search/?api=1&query=" + h.location; 
								actionBtn = '<a href="' + mapLink + '" target="_blank" class="block mt-4 w-full py-3 rounded-xl bg-blue-600 hover:bg-blue-500 text-white text-center text-sm font-bold transition">📍 Get Directions (Paid)</a>';
							} else {
								statusBadge = '<span class="absolute top-4 right-4 bg-slate-900/90 text-slate-400 text-[10px] font-bold px-3 py-1 rounded-full z-20">TAKEN</span>';
								opacityClass = "opacity-50 grayscale";
								actionBtn = '<button disabled class="mt-4 w-full py-3 rounded-xl bg-slate-800/50 text-slate-500 text-xs font-bold cursor-not-allowed">Unavailable</button>';
							}
						} else {
							statusBadge = '<span class="absolute top-4 right-4 bg-white text-black text-[10px] font-bold px-3 py-1 rounded-full z-20 shadow-xl">AVAILABLE</span>';
							opacityClass = "";
							if (isOwner) {
								actionBtn = '<button onclick="deleteHouse(' + h.id + ')" class="mt-4 w-full py-3 rounded-xl border border-red-500/30 text-red-400 text-xs font-bold">Remove Listing</button>';
							} else if (isLoggedIn) {
								let waLink = "https://wa.me/" + h.phone + "?text=Hi, I found your " + h.type + " on Nyumba.";
								actionBtn = '<div class="grid grid-cols-2 gap-2 mt-4">' +
									'<a href="' + waLink + '" target="_blank" class="flex items-center justify-center bg-emerald-500 hover:bg-emerald-400 text-white text-xs font-bold py-3 rounded-xl transition">Chat</a>' +
									'<button onclick="payWithMpesa(' + h.id + ')" class="bg-indigo-600 hover:bg-indigo-500 text-white text-xs font-bold py-3 rounded-xl transition">💳 Pay Fee (1k)</button>' +
								'</div>';
							} else {
								actionBtn = '<a href="/login" class="block mt-4 w-full py-3 rounded-xl bg-slate-800 hover:bg-slate-700 text-white text-center text-xs font-bold transition">Login to Unlock</a>';
							}
						}

						const html = 
						'<div class="glass rounded-3xl p-4 flex flex-col relative group transition hover:-translate-y-1 hover:shadow-2xl ' + gridClass + ' ' + opacityClass + '">' +
							statusBadge +
							'<div class="w-full h-48 ' + (index===0 ? 'h-64' : '') + ' bg-slate-800 rounded-2xl overflow-hidden relative mb-4">' +
								'<img src="' + imageSrc + '" class="w-full h-full object-cover group-hover:scale-105 transition duration-700 ease-out">' +
								'<div class="absolute inset-0 bg-gradient-to-t from-slate-900/90 via-transparent to-transparent"></div>' +
								'<div class="absolute bottom-4 left-4">' +
									'<p class="text-xs font-bold text-indigo-300 uppercase tracking-widest mb-1">' + h.type + '</p>' +
									'<h3 class="text-2xl font-bold text-white leading-none">' + h.location + '</h3>' +
								'</div>' +
							'</div>' +
							'<div class="flex-1"><p class="text-slate-400 text-sm line-clamp-2 leading-relaxed">' + h.details + '</p></div>' +
							'<div class="mt-4 pt-4 border-t border-white/5 flex items-end justify-between">' +
								'<div><p class="text-[10px] text-slate-500 uppercase font-bold">Monthly Rent</p><p class="text-xl font-bold text-white">KES ' + h.price.toLocaleString() + '</p></div>' +
								'<div class="text-right"><p class="text-[10px] text-slate-500 uppercase font-bold">Bills</p><p class="text-sm font-medium text-slate-300">~' + h.utilities.toLocaleString() + '</p></div>' +
							'</div>' +
							actionBtn +
						'</div>';
						container.innerHTML += html;
					});
				});
			}

			function uploadHouse() {
				const formData = new FormData();
				formData.append("location", document.getElementById('loc').value);
				formData.append("type", document.getElementById('type').value);
				formData.append("price", document.getElementById('price').value);
				formData.append("utilities", document.getElementById('utils').value);
				formData.append("details", document.getElementById('details').value);
				formData.append("map_url", document.getElementById('map_url').value);
				formData.append("tags", JSON.stringify([]));
				const fileInput = document.getElementById('photos');
				for (let i = 0; i < fileInput.files.length; i++) { formData.append("photos", fileInput.files[i]); }
				fetch('/houses/upload', { method: 'POST', body: formData }).then(res => { 
					fetchHouses(); showToast("Published Successfully");
					document.getElementById('loc').value = ""; document.getElementById('price').value = "";
				});
			}

			function deleteHouse(id) {
				if(!confirm("Are you sure?")) return;
				fetch('/houses/delete?id=' + id, {method: 'POST'}).then(() => { showToast("Listing Deleted"); fetchHouses(); });
			}
			function payWithMpesa(id) {
				let phone = prompt("M-Pesa Number:");
				if (!phone) return;
				showToast("Requesting M-Pesa...");
				fetch('/pay?id=' + id + '&phone=' + phone, {method: 'POST'})
				.then(res => res.json())
				.then(data => { 
					if(data.ResponseCode === "0") { showToast("Check your phone!"); fetchHouses(); }
					else { showToast("Connection Failed"); }
				});
			}
		</script>
	</body>
	</html>`
	fmt.Fprint(w, html)
}
