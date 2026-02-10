package main

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
)

// --- SAFARICOM CONFIG ---
const (
	// Your Sandbox Keys
	consumerKey    = "COBGyH3dHvYrVjLKG0Znfh8RR1yAPeVbZ6hZitAwgvquIqhL"
	consumerSecret = "ovklACIWd4ZMihM4Vv28TAwgEBG8MywaI5FOnHahzIPXAG16CTCikL2RSSqT4cog"

	shortCode    = "174379"
	passkey      = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
	mpesaAuthURL = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	mpesaPushURL = "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	callbackURL  = "https://nyumba-app.onrender.com/callback"
)

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

// 1. HOME PAGE
func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/manifest.json" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"name": "Nyumba Discovery", "short_name": "Nyumba", "start_url": "/", "display": "standalone",
			"background_color": "#0f172a", "theme_color": "#0f172a",
			"icons": [{ "src": "https://cdn-icons-png.flaticon.com/512/25/25694.png", "sizes": "192x192", "type": "image/png" }]
		}`)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	currentUser := getCurrentUser(r)
	isLoggedIn := "false"
	currentUsername := ""
	navLinks := `<a href="/login" class="text-sm font-medium text-slate-300 hover:text-white transition">Login</a>`
	landlordPanelDisplay := "none"
	myHubButton := ""

	if currentUser != nil {
		isLoggedIn = "true"
		currentUsername = currentUser.Username
		navLinks = `<a href="/logout" class="text-sm font-bold text-red-400 border border-red-500/30 px-3 py-1 rounded-full hover:bg-red-500/10 transition">Logout</a>`
		if currentUser.Role == "landlord" {
			landlordPanelDisplay = "block"
		}
		myHubButton = `<button onclick="openDashboard()" class="w-full bg-slate-800 hover:bg-slate-700 text-white font-bold py-3 rounded-xl flex items-center justify-center gap-2 transition mb-4 border border-white/10">🔐 My Unlocked Contacts</button>`
	}

	// START OF HTML
	html := `<!DOCTYPE html>
	<html>
	<head>
		<title>Nyumba Discovery</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="manifest" href="/manifest.json">
		<meta name="theme-color" content="#0f172a">
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet">
		<script src="https://cdn.tailwindcss.com"></script>
		<style>
			body { font-family: 'Outfit', sans-serif; background: #0f172a; color: #f8fafc; }
			.glass-card { background: rgba(30, 41, 59, 0.7); border: 1px solid rgba(255, 255, 255, 0.1); backdrop-filter: blur(10px); transform-style: preserve-3d; transition: transform 0.1s ease-out; }
			.glass-sidebar { background: #1e293b; border-right: 1px solid rgba(255, 255, 255, 0.05); }
			.nav-arrow { background: rgba(0,0,0,0.8); color: white; border-radius: 50%; width: 32px; height: 32px; display: flex; align-items: center; justify-content: center; cursor: pointer; border: 1px solid rgba(255,255,255,0.3); z-index: 40; transition: 0.2s; }
			.gallery-btn { background: rgba(255,255,255,0.1); backdrop-filter: blur(5px); border: 1px solid rgba(255,255,255,0.2); padding: 8px 16px; border-radius: 99px; font-size: 12px; font-weight: bold; color: white; cursor: pointer; }
			/* 3D PARALLAX CSS */
			.glass-card > .absolute, .glass-card > div.mt-4 { transform: translateZ(40px); box-shadow: 0 10px 20px rgba(0,0,0,0.3); }
		</style>
	</head>
	<body class="h-screen flex flex-col md:flex-row overflow-hidden">
		<div class="md:hidden flex items-center justify-between p-4 bg-slate-900 border-b border-white/5 z-40">
			<h1 class="text-xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300">Nyumba.</h1>
			<button onclick="toggleMenu()" class="text-white text-2xl px-2">☰</button>
		</div>
		<div id="backdrop" onclick="toggleMenu()" class="fixed inset-0 bg-black/80 z-40 hidden md:hidden transition-opacity"></div>

		<aside id="sidebar" class="fixed inset-y-0 left-0 z-50 w-80 bg-[#1e293b] md:bg-transparent md:static md:flex flex-col h-full transform -translate-x-full md:translate-x-0 transition-transform duration-300 glass-sidebar">
			<div class="p-8 pb-4 flex justify-between items-center">
				<div><h1 class="text-4xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300">Nyumba.</h1><p class="text-xs text-slate-500 font-medium tracking-widest uppercase mt-2">Curated Living</p></div>
				<button onclick="toggleMenu()" class="md:hidden text-white text-3xl">&times;</button>
			</div>
			<div class="px-6 py-4 space-y-6 flex-1 overflow-y-auto">
				` + myHubButton + `
				<div style="display: ` + landlordPanelDisplay + `;" class="glass-card rounded-2xl p-5 mb-8">
					<h3 class="text-xs font-bold text-indigo-400 uppercase tracking-wider mb-4">Landlord Mode</h3>
					<div class="space-y-3">
						<input id="building" type="text" placeholder="Apartment Name" class="w-full bg-slate-900 border border-slate-700 rounded-lg px-3 py-2 text-sm text-white outline-none">
						<input id="loc" type="text" placeholder="Location (e.g. Juja)" class="w-full bg-slate-900 border border-slate-700 rounded-lg px-3 py-2 text-sm text-white outline-none">
						<input id="map_url" type="text" placeholder="📍 Google Maps Link" class="w-full bg-slate-900 border border-indigo-500/30 rounded-lg px-3 py-2 text-sm text-indigo-300 outline-none">
						<select id="type" class="w-full bg-slate-900 border border-slate-700 rounded-lg px-3 py-2 text-sm outline-none"><option>Bedsitter</option><option>One Bedroom</option><option>Two Bedroom</option><option>Studio</option></select>
						<div class="grid grid-cols-2 gap-2"><input id="price" type="number" placeholder="Rent" class="w-full bg-slate-900 border border-slate-700 rounded-lg px-3 py-2 text-sm outline-none"><input id="utils" type="number" placeholder="Bills" class="w-full bg-slate-900 border border-slate-700 rounded-lg px-3 py-2 text-sm outline-none"></div>
						<input id="photos" type="file" multiple class="text-xs text-slate-500">
						<textarea id="details" placeholder="Description..." class="w-full bg-slate-900 border border-slate-700 rounded-lg px-3 py-2 text-sm h-16 resize-none outline-none"></textarea>
						<button onclick="uploadHouse()" class="w-full bg-indigo-600 hover:bg-indigo-500 text-white font-bold py-2 rounded-lg text-sm transition">Post Listing</button>
					</div>
				</div>
				<div class="space-y-4">
					<div><label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block">Where to?</label><input id="searchLoc" onkeyup="fetchHouses()" type="text" placeholder="Try 'Kileleshwa'..." class="w-full bg-slate-800 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-indigo-500 outline-none"></div>
					<div><label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 block">Max Budget</label><input id="searchPrice" onkeyup="fetchHouses()" type="number" placeholder="Any Price" class="w-full bg-slate-800 border border-slate-700 rounded-xl px-4 py-3 text-white focus:ring-2 focus:ring-emerald-500 outline-none"></div>
				</div>
			</div>
			<div class="p-6 border-t border-white/5 flex items-center justify-between bg-[#1e293b]">
				<div class="flex items-center gap-3"><div class="w-8 h-8 rounded-full bg-indigo-500 flex items-center justify-center text-xs font-bold">👤</div><div class="text-sm"><div class="font-bold text-white">` + currentUsername + `</div></div></div>
				` + navLinks + `
			</div>
		</aside>
		<main class="flex-1 h-full overflow-y-auto bg-slate-900 relative z-10">
			<div class="p-4 md:p-8 max-w-[1600px] mx-auto">
				<header class="flex justify-between items-end mb-8 mt-4 md:mt-0">
					<div><h2 class="text-2xl md:text-3xl font-light text-white">Discover <span class="font-bold text-indigo-400">Sanctuary</span></h2><p class="text-slate-400 mt-1 text-sm">Pay the viewing fee to unlock landlord contacts instantly.</p></div>
					<div id="offline-badge" class="hidden bg-amber-500/20 text-amber-500 border border-amber-500/50 px-4 py-2 rounded-lg text-xs font-bold animate-pulse">⚠️ OFFLINE MODE</div>
				</header>
				<div id="results-area" class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6 pb-20"></div>
			</div>
		</main>

		<div id="dashboard-modal" class="fixed inset-0 z-[100] bg-black/95 hidden flex flex-col items-center justify-center p-4">
			<div class="bg-slate-900 w-full max-w-md rounded-3xl p-6 border border-white/10 relative">
				<button onclick="closeDashboard()" class="absolute top-4 right-4 text-white text-2xl">&times;</button>
				<h2 class="text-xl font-bold text-white mb-4">My Unlocked Contacts</h2>
				<p class="text-xs text-slate-400 mb-6">These are the landlords you have paid to connect with. Contact them directly to view.</p>
				<div id="dashboard-list" class="space-y-3 max-h-[60vh] overflow-y-auto"></div>
			</div>
		</div>

		<div id="gallery-modal" class="fixed inset-0 z-[100] bg-black/95 hidden flex flex-col items-center justify-center p-4">
			<button onclick="closeGallery()" class="absolute top-6 right-6 text-white text-4xl">&times;</button>
			<img id="gallery-img" src="" class="max-h-[80vh] max-w-full rounded-lg shadow-2xl object-contain mb-4">
			<div class="flex items-center gap-6"><button onclick="navGallery(-1)" class="text-white text-3xl">❮</button><p id="gallery-counter" class="text-slate-400 font-medium">1 / 1</p><button onclick="navGallery(1)" class="text-white text-3xl">❯</button></div>
		</div>
		<div id="toast" class="fixed top-6 left-1/2 -translate-x-1/2 bg-indigo-600 px-6 py-3 rounded-full text-sm font-bold text-white shadow-2xl translate-y-[-200%] transition-transform duration-500 z-[60] flex items-center gap-2"><span class="text-lg">✨</span> <span id="toast-msg">Notification</span></div>

		<script>
			const isLoggedIn = ` + isLoggedIn + `;
			const currentUsername = "` + currentUsername + `";
			let houseImages = {}; let currentGalleryID = 0; let galleryIndex = 0; let autoScrollInterval;
			let allHousesData = []; // Store raw data for the dashboard

			// --- SKELETON LOADER ---
			const skeletonHTML = '<div class="glass-card rounded-3xl p-4 flex flex-col relative animate-pulse border border-white/5"><div class="w-full h-48 bg-slate-800 rounded-2xl mb-4"></div><div class="h-4 bg-slate-800 rounded w-1/3 mb-2"></div><div class="h-6 bg-slate-800 rounded w-3/4 mb-4"></div><div class="grid grid-cols-2 gap-2 mt-4"><div class="h-10 bg-slate-800 rounded-xl"></div><div class="h-10 bg-slate-800 rounded-xl"></div></div></div>';

			document.addEventListener("DOMContentLoaded", () => { fetchHouses(); startAutoScroll(); });

			// --- DASHBOARD LOGIC ---
			function openDashboard() {
				const container = document.getElementById('dashboard-list');
				container.innerHTML = "";
				// Filter houses where I AM the tenant (I paid the fee)
				const myHouses = allHousesData.filter(h => h.is_booked === true); 
				if(myHouses.length === 0) {
					container.innerHTML = "<div class='text-center text-slate-500 py-10'>No contacts unlocked yet.</div>";
				} else {
					myHouses.forEach(h => {
						let item = '<div class="bg-slate-800 p-4 rounded-xl flex justify-between items-center">' +
							'<div><div class="font-bold text-white">' + h.building_name + '</div><div class="text-xs text-slate-400">' + h.location + '</div></div>' +
							'<a href="tel:' + h.phone + '" class="bg-emerald-500 text-white px-4 py-2 rounded-lg text-xs font-bold">📞 ' + h.phone + '</a>' +
						'</div>';
						container.innerHTML += item;
					});
				}
				document.getElementById('dashboard-modal').classList.remove('hidden');
			}
			function closeDashboard() { document.getElementById('dashboard-modal').classList.add('hidden'); }

			function add3DEffect(card) {
				card.addEventListener('mousemove', (e) => {
					const rect = card.getBoundingClientRect();
					const x = e.clientX - rect.left; const y = e.clientY - rect.top;
					const centerX = rect.width / 2; const centerY = rect.height / 2;
					const rotateX = ((y - centerY) / centerY) * -10; const rotateY = ((x - centerX) / centerX) * 10;
					card.style.transform = "perspective(1000px) rotateX(" + rotateX + "deg) rotateY(" + rotateY + "deg) scale(1.02)";
				});
				card.addEventListener('mouseleave', () => { card.style.transform = "perspective(1000px) rotateX(0) rotateY(0) scale(1)"; });
			}

			function toggleMenu() { document.getElementById('sidebar').classList.toggle('-translate-x-full'); document.getElementById('backdrop').classList.toggle('hidden'); }
			function showToast(msg) { const t = document.getElementById("toast"); document.getElementById("toast-msg").innerText = msg; t.classList.remove("translate-y-[-200%]"); setTimeout(() => t.classList.add("translate-y-[-200%]"), 3000); }
			function openGallery(id) { const images = houseImages[id]; if(!images || images.length === 0) return; currentGalleryID = id; galleryIndex = 0; updateGalleryView(); document.getElementById('gallery-modal').classList.remove('hidden'); }
			function closeGallery() { document.getElementById('gallery-modal').classList.add('hidden'); }
			function navGallery(step) { const images = houseImages[currentGalleryID]; galleryIndex += step; if(galleryIndex >= images.length) galleryIndex = 0; if(galleryIndex < 0) galleryIndex = images.length - 1; updateGalleryView(); }
			function updateGalleryView() { const images = houseImages[currentGalleryID]; document.getElementById('gallery-img').src = images[galleryIndex]; document.getElementById('gallery-counter').innerText = (galleryIndex + 1) + " / " + images.length; }
			function startAutoScroll() { if (autoScrollInterval) clearInterval(autoScrollInterval); autoScrollInterval = setInterval(() => { document.querySelectorAll('[id^="img-"]').forEach(img => { let id = img.id.split('-')[1]; if (document.getElementById('gallery-modal').classList.contains('hidden')) { changeSlide(id, 1); } }); }, 3500); }
			function changeSlide(id, step) { const images = houseImages[id]; if (!images || images.length <= 1) return; let imgEl = document.getElementById('img-' + id); let current = parseInt(imgEl.dataset.index || 0); let next = current + step; if (next >= images.length) next = 0; if (next < 0) next = images.length - 1; imgEl.dataset.index = next; imgEl.src = images[next]; }

			function fetchHouses() {
				const container = document.getElementById('results-area');
				container.innerHTML = skeletonHTML + skeletonHTML + skeletonHTML + skeletonHTML;
				fetch('/houses').then(res => res.json()).then(data => {
					allHousesData = data; // Save for dashboard
					localStorage.setItem('nyumba_cache', JSON.stringify(data));
					renderList(data, true);
				}).catch(err => {
					const cached = localStorage.getItem('nyumba_cache');
					if(cached) { allHousesData = JSON.parse(cached); renderList(allHousesData, false); showToast("Offline Mode"); } 
					else { container.innerHTML = "<div class='col-span-full text-center text-slate-500'>Connection Failed</div>"; }
				});
			}

			function renderList(data, isOnline) {
				const container = document.getElementById('results-area'); container.innerHTML = "";
				const sLoc = document.getElementById('searchLoc').value.toLowerCase();
				const sPrice = document.getElementById('searchPrice').value;
				const badge = document.getElementById('offline-badge');
				if(isOnline) { badge.classList.add('hidden'); } else { badge.classList.remove('hidden'); }

				let filtered = data.filter(h => {
					if(sLoc && !h.location.toLowerCase().includes(sLoc)) return false;
					if(sPrice && h.price > parseFloat(sPrice)) return false;
					return true;
				});

				if (filtered.length === 0) { container.innerHTML = "<div class='col-span-full text-center text-slate-500 py-20'>No sanctuaries found.</div>"; return; }
				
				filtered.forEach((h) => {
					houseImages[h.id] = h.image_urls;
					const isOwner = (h.owner === currentUsername);
					let imageSrc = (h.image_urls && h.image_urls.length > 0) ? h.image_urls[0] : 'https://via.placeholder.com/600x400?text=No+Image';
					
					let actionBtn;
					if (h.is_booked) {
						// For the lead-gen model, if it's booked (unlocked), we show a button telling them to check the dashboard
						if (isLoggedIn) {
							actionBtn = '<button onclick="openDashboard()" class="mt-4 w-full py-3 rounded-xl bg-emerald-500/20 text-emerald-400 border border-emerald-500/50 text-xs font-bold">✨ Contact Unlocked (Check Hub)</button>';
						} else {
							actionBtn = '<button disabled class="mt-4 w-full py-3 rounded-xl bg-slate-800/50 text-slate-500 text-xs font-bold">Taken</button>';
						}
					} else if (isLoggedIn) {
						actionBtn = '<div class="grid grid-cols-2 gap-2 mt-4"><a href="https://wa.me/' + h.phone + '" class="flex items-center justify-center bg-emerald-500 hover:bg-emerald-400 text-white text-xs font-bold py-3 rounded-xl transition">Chat</a><button onclick="payWithMpesa(' + h.id + ')" class="bg-indigo-600 hover:bg-indigo-500 text-white text-xs font-bold py-3 rounded-xl transition">💳 Unlock (1k)</button></div>';
					} else {
						actionBtn = '<a href="/login" class="block mt-4 w-full py-3 rounded-xl bg-slate-800 hover:bg-slate-700 text-white text-center text-xs font-bold transition">Login to Unlock</a>';
					}

					const card = document.createElement('div');
					card.className = "glass-card rounded-3xl p-4 flex flex-col relative group transition hover:-translate-y-1 hover:shadow-2xl";
					card.innerHTML = 
						'<div class="w-full h-48 bg-slate-800 rounded-2xl overflow-hidden relative mb-4"><img id="img-' + h.id + '" src="' + imageSrc + '" class="w-full h-full object-cover transition duration-700 ease-out"></div>' +
						'<div class="flex-1"><h3 class="text-xl font-bold text-white">' + h.building_name + '</h3><p class="text-xs text-slate-400 mb-2">📍 ' + h.location + '</p><p class="text-slate-400 text-sm line-clamp-2">' + h.details + '</p></div>' +
						'<div class="mt-4 pt-4 border-t border-white/5 flex items-end justify-between"><div><p class="text-[10px] text-slate-500 uppercase font-bold">Rent</p><p class="text-xl font-bold text-white">KES ' + h.price.toLocaleString() + '</p></div></div>' +
						actionBtn;
					
					add3DEffect(card);
					container.appendChild(card);
				});
			}

			function uploadHouse() {
				const formData = new FormData();
				formData.append("building_name", document.getElementById('building').value);
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
				.then(data => { 
					if(data.ResponseCode === "0") { showToast("Success! Check 'My Hub'"); fetchHouses(); } 
					else { showToast(data.CustomerMessage || "Connection Failed"); } 
				})
				.catch(err => { showToast("System Error"); });
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
		ID:           len(houses) + 1,
		BuildingName: r.FormValue("building_name"),
		Location:     r.FormValue("location"),
		Type:         r.FormValue("type"),
		Price:        p,
		Utilities:    u,
		Details:      r.FormValue("details"),
		ImageURLs:    imageURLs,
		Phone:        currentUser.Phone,
		Owner:        currentUser.Username,
		IsBooked:     false,
		MapURL:       r.FormValue("map_url"),
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
func payHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	rawPhone := r.URL.Query().Get("phone")
	phone := formatPhoneNumber(rawPhone)

	var selectedHouse *House
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

	err := initiateSTKPush(phone, "1")
	if err != nil {
		fmt.Fprintf(w, `{"ResponseCode": "1", "CustomerMessage": "M-Pesa Error: %s"}`, err.Error())
		return
	}

	selectedHouse.IsBooked = true
	selectedHouse.TenantPhone = phone
	saveData(houseFile, houses)
	fmt.Fprint(w, `{"ResponseCode": "0", "CustomerMessage": "Sent"}`)
}

// 7. FILE SERVER
func serveMedia(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "."+r.URL.Path) }

// --- CONNECTORS ---
func getHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: CookieName, Value: "", Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func uploadHouseHandler(w http.ResponseWriter, r *http.Request) { uploadHouse(w, r) }

// --- INTERNAL MPESA LOGIC ---
func initiateSTKPush(phoneNumber, amount string) error {
	token, err := getAccessToken()
	if err != nil {
		return err
	}
	timestamp := time.Now().Format("20060102150405")
	password := base64.StdEncoding.EncodeToString([]byte(shortCode + passkey + timestamp))
	headers := map[string]string{"Authorization": "Bearer " + token, "Content-Type": "application/json"}
	payload := map[string]string{
		"BusinessShortCode": shortCode, "Password": password, "Timestamp": timestamp,
		"TransactionType": "CustomerPayBillOnline", "Amount": amount, "PartyA": phoneNumber,
		"PartyB":           shortCode,
		"PhoneNumber":      phoneNumber,
		"CallBackURL":      callbackURL,
		"AccountReference": "NyumbaApp",
		"TransactionDesc":  "Viewing Fee",
	}
	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", mpesaPushURL, bytes.NewBuffer(jsonPayload))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed: %s", string(body))
	}
	return nil
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
