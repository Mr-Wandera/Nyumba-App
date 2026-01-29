package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 1. HOME PAGE
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	currentUser := getCurrentUser(r)
	isLoggedIn := "false"
	currentUsername := ""
	currentUserPhone := ""

	navLinks := `<a href="/login" class="text-sm font-medium text-slate-300 hover:text-white transition">Login</a>`
	landlordPanelDisplay := "none"

	if currentUser != nil {
		isLoggedIn = "true"
		currentUsername = currentUser.Username
		currentUserPhone = currentUser.Phone
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
			body { font-family: 'Outfit', sans-serif; background: #0f172a; color: #f8fafc; }
			::-webkit-scrollbar { width: 6px; }
			::-webkit-scrollbar-track { background: #0f172a; }
			::-webkit-scrollbar-thumb { background: #334155; border-radius: 3px; }
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
						<input id="searchLoc" onkeyup="fetchHouses()" type="text" placeholder="Try 'Kileleshwa'..." class="w-full bg-slate-800 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-indigo-500 outline-none">
					</div>
					<div>
						<label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block">Max Budget</label>
						<input id="searchPrice" onkeyup="fetchHouses()" type="number" placeholder="Any Price" class="w-full bg-slate-800 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-emerald-500 outline-none">
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
				fetch('/houses/upload', { method: 'POST', body: formData }).then(res => { fetchHouses(); showToast("Published Successfully"); });
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
				.then(data => { if(data.ResponseCode === "0") { showToast("Check your phone!"); fetchHouses(); } else { showToast("Connection Failed"); } });
			}
		</script>
	</body>
	</html>`
	fmt.Fprint(w, html)
}

// 2. LOGIN HANDLER
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		for _, u := range users {
			if u.Username == username && u.Password == password {
				http.SetCookie(w, &http.Cookie{Name: CookieName, Value: username, Path: "/"})
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}
	html := `<!DOCTYPE html><html><head><title>Login</title><meta name="viewport" content="width=device-width"><link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;600&display=swap" rel="stylesheet"><script src="https://cdn.tailwindcss.com"></script><style>body{font-family:'Outfit',sans-serif;background:#0b0f19;color:#fff}.glass{background:rgba(30,41,59,0.4);backdrop-filter:blur(16px);border:1px solid rgba(255,255,255,0.05)}</style></head><body class="h-screen flex items-center justify-center"><div class="glass p-10 rounded-3xl w-full max-w-sm"><h1 class="text-3xl font-bold mb-6 text-center">Login</h1><form method="POST"><input name="username" placeholder="Username" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><input name="password" type="password" placeholder="Password" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><button class="w-full bg-indigo-600 py-3 rounded-xl font-bold">Sign In</button></form><a href="/signup" class="block text-center mt-6 text-slate-400 text-sm">Create Account</a></div></body></html>`
	fmt.Fprint(w, html)
}

// 3. SIGNUP HANDLER
func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		phone := r.FormValue("phone")
		role := r.FormValue("role")
		newUser := User{Username: username, Password: password, Phone: phone, Role: role}
		users = append(users, newUser)
		saveData(userFile, users)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	html := `<!DOCTYPE html><html><head><title>Join</title><meta name="viewport" content="width=device-width"><link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;600&display=swap" rel="stylesheet"><script src="https://cdn.tailwindcss.com"></script><style>body{font-family:'Outfit',sans-serif;background:#0b0f19;color:#fff}.glass{background:rgba(30,41,59,0.4);backdrop-filter:blur(16px);border:1px solid rgba(255,255,255,0.05)}</style></head><body class="h-screen flex items-center justify-center"><div class="glass p-10 rounded-3xl w-full max-w-sm"><h1 class="text-3xl font-bold mb-6 text-center">Join</h1><form method="POST"><input name="username" placeholder="Username" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><input name="phone" placeholder="Phone" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><input name="password" type="password" placeholder="Password" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><select name="role" class="w-full bg-slate-900 border border-slate-700 rounded-xl p-3 mb-4"><option value="renter">Renter</option><option value="landlord">Landlord</option></select><button class="w-full bg-indigo-600 py-3 rounded-xl font-bold">Create Account</button></form><a href="/login" class="block text-center mt-6 text-slate-400 text-sm">Login Here</a></div></body></html>`
	fmt.Fprint(w, html)
}

// 4. UPLOAD LOGIC
func uploadHouse(w http.ResponseWriter, r *http.Request) {
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
	newHouse := House{
		ID: len(houses) + 1, Location: r.FormValue("location"), Type: r.FormValue("type"),
		Price: p, Utilities: u, Details: r.FormValue("details"), ImageURLs: imageURLs,
		Phone: currentUser.Phone, Owner: currentUser.Username, IsBooked: false, MapURL: r.FormValue("map_url"),
	}
	houses = append(houses, newHouse)
	saveData(houseFile, houses)
	w.WriteHeader(http.StatusCreated)
}

// 5. DELETE HANDLER
func deleteHouseHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	newHouses := []House{}
	for _, h := range houses {
		if h.ID != id {
			newHouses = append(newHouses, h)
		}
	}
	houses = newHouses
	saveData(houseFile, houses)
	w.WriteHeader(200)
}

// 6. PAY HANDLER
// 6. PAY HANDLER (Connected to Safaricom Sandbox)
func payHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Get the Phone Number & House ID
	idStr := r.URL.Query().Get("id")
	phone := r.URL.Query().Get("phone")
	id, _ := strconv.Atoi(idStr)

	// 2. Find the House
	var selectedHouse *House
	for i, h := range houses {
		if h.ID == id {
			selectedHouse = &houses[i]
			break
		}
	}

	if selectedHouse == nil {
		http.Error(w, "House not found", http.StatusNotFound)
		return
	}

	// 3. TRIGGER THE POP-UP (Call the function in mpesa.go)
	// We send "1" KES because Daraja Sandbox requires at least 1 shilling to trigger.
	err := initiateSTKPush(phone, "1")
	if err != nil {
		// If Safaricom fails (e.g., wrong phone number format)
		http.Error(w, "M-Pesa Failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Unlock the House (Optimistic Unlock)
	// In a real production app, we would wait for the "Callback" to confirm payment.
	// For this student project, we unlock it as soon as the Request is successful.
	selectedHouse.IsBooked = true
	selectedHouse.TenantPhone = phone
	saveData(houseFile, houses)

	// 5. Tell the Frontend it worked
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"ResponseCode": "0", "CustomerMessage": "STK Push Sent. Check phone."}`)
}
