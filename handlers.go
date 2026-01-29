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
	if r.Method == "GET" {
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Sign Up</title>
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<style>
				body { font-family: sans-serif; background: #f3f4f6; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; }
				.card { background: white; padding: 2rem; border-radius: 16px; width: 100%; max-width: 400px; text-align: center; }
				input, select { width: 100%; padding: 10px; margin-bottom: 10px; border: 1px solid #ccc; border-radius: 8px; box-sizing: border-box; }
				button { width: 100%; padding: 10px; background: #4f46e5; color: white; border: none; border-radius: 8px; cursor: pointer; }
			</style>
		</head>
		<body>
			<div class="card">
				<h2>✨ Join Nyumba</h2>
				<form method="POST" action="/signup">
					<input type="text" name="username" placeholder="Username" required>
					<input type="password" name="password" placeholder="Password" required>
					<input type="text" name="phone" placeholder="Phone (e.g. 2547...)" required>
					<select name="role">
						<option value="renter">👤 Renter</option>
						<option value="landlord">🏠 Landlord</option>
					</select>
					<button>Create Account</button>
				</form>
				<p><a href="/login">Login</a></p>
			</div>
		</body>
		</html>`
		fmt.Fprint(w, html)
		return
	}
	username := r.FormValue("username")
	for _, u := range users {
		if u.Username == username {
			http.Error(w, "User exists!", http.StatusBadRequest)
			return
		}
	}
	newUser := User{Username: username, Password: r.FormValue("password"), Role: r.FormValue("role"), Phone: r.FormValue("phone")}
	users = append(users, newUser)
	saveData(userFile, users)
	http.SetCookie(w, &http.Cookie{Name: CookieName, Value: username, Path: "/"})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Login</title>
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<style>
				body { font-family: sans-serif; background: #f3f4f6; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; }
				.card { background: white; padding: 2rem; border-radius: 16px; width: 100%; max-width: 350px; text-align: center; }
				input { width: 100%; padding: 10px; margin-bottom: 10px; border: 1px solid #ccc; border-radius: 8px; box-sizing: border-box; }
				button { width: 100%; padding: 10px; background: #4f46e5; color: white; border: none; border-radius: 8px; cursor: pointer; }
			</style>
		</head>
		<body>
			<div class="card">
				<h2>🔐 Login</h2>
				<form method="POST" action="/login">
					<input type="text" name="username" placeholder="Username" required>
					<input type="password" name="password" placeholder="Password" required>
					<button>Sign In</button>
				</form>
				<p><a href="/signup">Create Account</a></p>
			</div>
		</body>
		</html>`
		fmt.Fprint(w, html)
		return
	}
	user := r.FormValue("username")
	pass := r.FormValue("password")
	for _, u := range users {
		if u.Username == user && u.Password == pass {
			http.SetCookie(w, &http.Cookie{Name: CookieName, Value: user, Path: "/"})
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
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

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	currentUser := getCurrentUser(r)
	isLoggedIn := "false"
	currentUsername := ""

	welcomeMsg := "Welcome"
	navLinks := `<a href="/login" class="text-sm font-medium text-slate-300 hover:text-white transition">Login</a>`
	landlordPanelDisplay := "none"

	if currentUser != nil {
		isLoggedIn = "true"
		currentUsername = currentUser.Username
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
			body { font-family: 'Outfit', sans-serif; background: #0b0f19; color: #f8fafc; }
			
			/* Custom Scrollbar */
			::-webkit-scrollbar { width: 6px; }
			::-webkit-scrollbar-track { background: #0b0f19; }
			::-webkit-scrollbar-thumb { background: #334155; border-radius: 3px; }

			/* Glassmorphism Classes */
			.glass {
				background: rgba(30, 41, 59, 0.4);
				backdrop-filter: blur(16px);
				-webkit-backdrop-filter: blur(16px);
				border: 1px solid rgba(255, 255, 255, 0.05);
			}
			.glass-strong {
				background: rgba(15, 23, 42, 0.8);
				backdrop-filter: blur(20px);
				border-right: 1px solid rgba(255, 255, 255, 0.05);
			}

			/* Animations */
			@keyframes float { 0% { transform: translateY(0px); } 50% { transform: translateY(-10px); } 100% { transform: translateY(0px); } }
			.animate-float { animation: float 6s ease-in-out infinite; }
		</style>
	</head>
	<body class="h-screen flex overflow-hidden selection:bg-indigo-500 selection:text-white">
		
		<aside class="w-80 flex-shrink-0 glass-strong flex flex-col h-full relative z-20">
			<div class="p-8 pb-4">
				<h1 class="text-4xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300 cursor-pointer">
					Nyumba.
				</h1>
				<p class="text-xs text-slate-500 font-medium tracking-widest uppercase mt-2">Curated Living</p>
			</div>

			<div class="px-6 py-4 space-y-6 flex-1 overflow-y-auto">
				
				<div style="display: ` + landlordPanelDisplay + `;" class="glass rounded-2xl p-5 mb-8 border border-indigo-500/20 shadow-lg shadow-indigo-900/20">
					<h3 class="text-xs font-bold text-indigo-400 uppercase tracking-wider mb-4 flex items-center gap-2">
						<span class="w-2 h-2 rounded-full bg-indigo-400 animate-pulse"></span> Landlord Mode
					</h3>
					<div class="space-y-3">
						<input id="loc" type="text" placeholder="Location" class="w-full bg-slate-900/50 border border-slate-700/50 rounded-lg px-3 py-2 text-sm focus:border-indigo-500 transition outline-none">
						<select id="type" class="w-full bg-slate-900/50 border border-slate-700/50 rounded-lg px-3 py-2 text-sm outline-none">
							<option>Bedsitter</option><option>One Bedroom</option><option>Two Bedroom</option><option>Studio</option>
						</select>
						<div class="grid grid-cols-2 gap-2">
							<input id="price" type="number" placeholder="Rent" class="w-full bg-slate-900/50 border border-slate-700/50 rounded-lg px-3 py-2 text-sm outline-none">
							<input id="utils" type="number" placeholder="Bills" class="w-full bg-slate-900/50 border border-slate-700/50 rounded-lg px-3 py-2 text-sm outline-none">
						</div>
						<input id="photos" type="file" multiple class="text-xs text-slate-500 file:mr-2 file:py-1 file:px-2 file:rounded-md file:border-0 file:text-xs file:font-semibold file:bg-indigo-500/10 file:text-indigo-400">
						<textarea id="details" placeholder="Description..." class="w-full bg-slate-900/50 border border-slate-700/50 rounded-lg px-3 py-2 text-sm h-16 outline-none resize-none"></textarea>
						<button onclick="uploadHouse()" class="w-full bg-indigo-600 hover:bg-indigo-500 text-white font-bold py-2 rounded-lg text-sm transition shadow-lg shadow-indigo-600/20">Post Listing</button>
					</div>
				</div>

				<div class="space-y-4">
					<div class="relative group">
						<label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block">Where to?</label>
						<input id="searchLoc" onkeyup="fetchHouses()" type="text" placeholder="Try 'Kileleshwa'..." 
							class="w-full bg-slate-800/50 border border-slate-700 rounded-xl px-4 py-3 text-white placeholder-slate-600 focus:ring-2 focus:ring-indigo-500/50 focus:border-transparent transition outline-none text-lg font-medium">
						<div class="absolute right-3 top-8 text-slate-600">🔍</div>
					</div>

					<div>
						<label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block">Max Budget</label>
						<input id="searchPrice" onkeyup="fetchHouses()" type="number" placeholder="Any Price" 
							class="w-full bg-slate-800/50 border border-slate-700 rounded-xl px-4 py-3 text-white placeholder-slate-600 focus:ring-2 focus:ring-emerald-500/50 transition outline-none text-lg font-medium">
					</div>
					
					<div class="grid grid-cols-2 gap-3 mt-6">
						<div class="glass p-3 rounded-xl text-center">
							<div class="text-2xl font-bold text-white">12</div>
							<div class="text-[10px] text-slate-400 uppercase tracking-wider">New Today</div>
						</div>
						<div class="glass p-3 rounded-xl text-center">
							<div class="text-2xl font-bold text-emerald-400">85%</div>
							<div class="text-[10px] text-slate-400 uppercase tracking-wider">Response</div>
						</div>
					</div>
				</div>
			</div>

			<div class="p-6 border-t border-white/5 flex items-center justify-between">
				<div class="flex items-center gap-3">
					<div class="w-8 h-8 rounded-full bg-gradient-to-tr from-indigo-500 to-purple-500 flex items-center justify-center text-xs font-bold">👤</div>
					<div class="text-sm">
						<div class="font-bold text-white leading-none">` + currentUsername + `</div>
						<div class="text-xs text-slate-500 mt-1">` + welcomeMsg + `</div>
					</div>
				</div>
				` + navLinks + `
			</div>
		</aside>

		<main class="flex-1 h-full overflow-y-auto relative bg-[#0b0f19]">
			
			<div class="fixed top-0 right-0 w-3/4 h-full pointer-events-none opacity-20">
				<div class="absolute top-1/4 right-1/4 w-96 h-96 bg-indigo-600 rounded-full mix-blend-screen filter blur-[120px] animate-float"></div>
				<div class="absolute bottom-1/4 right-10 w-64 h-64 bg-emerald-600 rounded-full mix-blend-screen filter blur-[100px] animate-float" style="animation-delay: 2s"></div>
			</div>

			<div class="p-8 max-w-[1600px] mx-auto">
				<header class="flex justify-between items-end mb-8 relative z-10">
					<div>
						<h2 class="text-3xl font-light text-white">Discover <span class="font-bold text-indigo-400">Sanctuary</span></h2>
						<p class="text-slate-400 mt-1">Curated homes for your next chapter.</p>
					</div>
					<button class="bg-white text-black px-4 py-2 rounded-full text-sm font-bold hover:bg-slate-200 transition">View Map 🗺️</button>
				</header>

				<div id="results-area" class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6 auto-rows-[minmax(180px,auto)] pb-20 relative z-10">
					</div>
			</div>
		</main>

		<div id="toast" class="fixed top-6 left-1/2 -translate-x-1/2 glass px-6 py-3 rounded-full text-sm font-bold text-white shadow-2xl translate-y-[-150%] transition-transform duration-500 z-50 flex items-center gap-2">
			<span class="text-lg">✨</span> <span id="toast-msg">Notification</span>
		</div>

		<script>
			const isLoggedIn = ` + isLoggedIn + `;
			const currentUsername = "` + currentUsername + `";

			document.addEventListener("DOMContentLoaded", () => fetchHouses());

			function showToast(msg) {
				const t = document.getElementById("toast");
				document.getElementById("toast-msg").innerText = msg;
				t.classList.remove("translate-y-[-150%]");
				setTimeout(() => t.classList.add("translate-y-[-150%]"), 3000);
			}

			function fetchHouses() {
				const sLoc = document.getElementById('searchLoc').value.toLowerCase();
				const sPrice = document.getElementById('searchPrice').value;

				fetch('/houses').then(res => res.json()).then(data => {
					const container = document.getElementById('results-area');
					container.innerHTML = "";
					
					// Filter
					let filtered = data.filter(h => {
						if(sLoc && !h.location.toLowerCase().includes(sLoc)) return false;
						if(sPrice && h.price > parseFloat(sPrice)) return false;
						return true;
					});

					if (filtered.length === 0) {
						container.innerHTML = "<div class='col-span-full text-center text-slate-500 py-20'>No sanctuaries found. Try adjusting your filters.</div>";
						return;
					}

					// Inject "Magazine Modules" to break the monotony
					// We'll insert a "Trending" card at index 2
					const trendingCard = { isPromo: true, type: 'trending' };
					if(filtered.length >= 2) filtered.splice(2, 0, trendingCard);

					filtered.forEach((h, index) => {
						// 1. RENDER PROMO CARDS (The "Tall Rectangle")
						if (h.isPromo) {
							const promoHtml = 
							'<div class="glass p-6 rounded-3xl flex flex-col justify-between row-span-2 border border-indigo-500/30 relative overflow-hidden group">' +
								'<div class="absolute inset-0 bg-gradient-to-b from-indigo-900/50 to-transparent opacity-50"></div>' +
								'<div class="relative z-10">' +
									'<h3 class="text-xl font-bold text-white mb-4">Trending 🔥</h3>' +
									'<ul class="space-y-3 text-sm font-medium text-slate-300">' +
										'<li class="flex justify-between border-b border-white/10 pb-2"><span>Kileleshwa</span> <span class="text-emerald-400">↑ 14%</span></li>' +
										'<li class="flex justify-between border-b border-white/10 pb-2"><span>Westlands</span> <span class="text-emerald-400">↑ 8%</span></li>' +
										'<li class="flex justify-between border-b border-white/10 pb-2"><span>Kilimani</span> <span class="text-emerald-400">↑ 5%</span></li>' +
										'<li class="flex justify-between border-b border-white/10 pb-2"><span>Juja</span> <span class="text-slate-500">- 2%</span></li>' +
									'</ul>' +
								'</div>' +
								'<button class="relative z-10 w-full mt-4 bg-white/10 hover:bg-white/20 text-white text-xs font-bold py-3 rounded-xl transition">View Heatmap</button>' +
							'</div>';
							container.innerHTML += promoHtml;
							return;
						}

						// 2. RENDER HOUSE CARDS
						const isOwner = (h.owner === currentUsername);
						
						// Layout Logic: First item is HERO (Wide)
						let gridClass = (index === 0) ? "md:col-span-2 row-span-2" : "";
						
						// Styling Logic
						let statusBadge, opacityClass, actionBtn;
						let imageSrc = (h.image_urls && h.image_urls.length > 0) ? h.image_urls[0] : 'https://via.placeholder.com/600x400?text=No+Image';

						if (h.is_booked) {
							if (isOwner) {
								statusBadge = '<span class="absolute top-4 right-4 bg-indigo-600 text-white text-[10px] font-bold px-3 py-1 rounded-full z-20 shadow-lg shadow-indigo-500/50">Booked by: ' + h.tenant_phone + '</span>';
								opacityClass = "border-indigo-500";
								actionBtn = '<button onclick="deleteHouse(' + h.id + ')" class="mt-4 w-full py-3 rounded-xl bg-slate-800 text-red-400 text-xs font-bold hover:bg-slate-700 transition">Delete Listing</button>';
							} else {
								statusBadge = '<span class="absolute top-4 right-4 bg-slate-900/90 text-slate-400 text-[10px] font-bold px-3 py-1 rounded-full z-20 backdrop-blur">TAKEN</span>';
								opacityClass = "opacity-50 grayscale";
								actionBtn = '<button disabled class="mt-4 w-full py-3 rounded-xl bg-slate-800/50 text-slate-500 text-xs font-bold cursor-not-allowed">Unavailable</button>';
							}
						} else {
							// Available
							statusBadge = '<span class="absolute top-4 right-4 bg-white text-black text-[10px] font-bold px-3 py-1 rounded-full z-20 shadow-xl">AVAILABLE</span>';
							opacityClass = "";
							
							if (isOwner) {
								actionBtn = '<button onclick="deleteHouse(' + h.id + ')" class="mt-4 w-full py-3 rounded-xl border border-red-500/30 text-red-400 text-xs font-bold hover:bg-red-500/10 transition">Remove Listing</button>';
							} else if (isLoggedIn) {
								let waLink = "https://wa.me/" + h.phone + "?text=Hi, I found your " + h.type + " on Nyumba.";
								actionBtn = '<div class="grid grid-cols-2 gap-2 mt-4">' +
									'<a href="' + waLink + '" target="_blank" class="flex items-center justify-center bg-emerald-500 hover:bg-emerald-400 text-white text-xs font-bold py-3 rounded-xl transition shadow-lg shadow-emerald-500/20">Chat</a>' +
									'<button onclick="payWithMpesa(' + h.id + ')" class="bg-indigo-600 hover:bg-indigo-500 text-white text-xs font-bold py-3 rounded-xl transition shadow-lg shadow-indigo-500/30">Book Now</button>' +
								'</div>';
							} else {
								actionBtn = '<a href="/login" class="block mt-4 w-full py-3 rounded-xl bg-slate-800 hover:bg-slate-700 text-white text-center text-xs font-bold transition">Login to Secure</a>';
							}
						}

						const html = 
						'<div class="glass rounded-3xl p-4 flex flex-col relative group transition hover:-translate-y-1 hover:shadow-2xl hover:shadow-indigo-500/10 ' + gridClass + ' ' + opacityClass + '">' +
							statusBadge +
							// Image Area
							'<div class="w-full h-48 ' + (index===0 ? 'h-64' : '') + ' bg-slate-800 rounded-2xl overflow-hidden relative mb-4">' +
								'<img src="' + imageSrc + '" class="w-full h-full object-cover group-hover:scale-105 transition duration-700 ease-out">' +
								'<div class="absolute inset-0 bg-gradient-to-t from-slate-900/90 via-transparent to-transparent"></div>' +
								'<div class="absolute bottom-4 left-4">' +
									'<p class="text-xs font-bold text-indigo-300 uppercase tracking-widest mb-1">' + h.type + '</p>' +
									'<h3 class="text-2xl font-bold text-white leading-none">' + h.location + '</h3>' +
								'</div>' +
							'</div>' +
							// Info Area
							'<div class="flex-1">' +
								'<p class="text-slate-400 text-sm line-clamp-2 leading-relaxed">' + h.details + '</p>' +
							'</div>' +
							// Price Area
							'<div class="mt-4 pt-4 border-t border-white/5 flex items-end justify-between">' +
								'<div>' +
									'<p class="text-[10px] text-slate-500 uppercase font-bold">Monthly Rent</p>' +
									'<p class="text-xl font-bold text-white">KES ' + h.price.toLocaleString() + '</p>' +
								'</div>' +
								'<div class="text-right">' +
									'<p class="text-[10px] text-slate-500 uppercase font-bold">Bills</p>' +
									'<p class="text-sm font-medium text-slate-300">~' + h.utilities.toLocaleString() + '</p>' +
								'</div>' +
							'</div>' +
							actionBtn +
						'</div>';

						container.innerHTML += html;
					});
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
			function uploadHouse() {
				const formData = new FormData();
				formData.append("location", document.getElementById('loc').value);
				formData.append("type", document.getElementById('type').value);
				formData.append("price", document.getElementById('price').value);
				formData.append("utilities", document.getElementById('utils').value);
				formData.append("details", document.getElementById('details').value);
				formData.append("tags", JSON.stringify([]));
				const fileInput = document.getElementById('photos');
				for (let i = 0; i < fileInput.files.length; i++) { formData.append("photos", fileInput.files[i]); }
				fetch('/houses/upload', { method: 'POST', body: formData }).then(res => { 
					fetchHouses(); showToast("Published Successfully");
					// Clear inputs
					document.getElementById('loc').value = "";
					document.getElementById('price').value = "";
				});
			}
		</script>
	</body>
	</html>`
	fmt.Fprint(w, html)
}
