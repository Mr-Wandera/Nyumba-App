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
	currentUsername := "" // 👈 We need to know WHO is logged in to match with the house owner

	welcomeMsg := "Welcome"
	navLinks := `<a href="/login" class="btn-secondary">Login</a>`
	addFormDisplay := "none"

	if currentUser != nil {
		isLoggedIn = "true"
		currentUsername = currentUser.Username
		welcomeMsg = "Hi, " + currentUser.Username
		navLinks = `<a href="/logout" class="btn-danger-outline">Logout</a>`
		if currentUser.Role == "landlord" {
			addFormDisplay = "block"
		}
	}

	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Nyumba</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700&display=swap" rel="stylesheet">
		<style>
			:root { --primary: #4f46e5; --bg: #f3f4f6; --text: #1f2937; --mpesa: #27ae60; --whatsapp: #25D366; }
			body { font-family: 'Inter', sans-serif; background: var(--bg); color: var(--text); margin: 0; padding-top: 80px; }
			.navbar { position: fixed; top: 0; left: 0; right: 0; background: white; height: 70px; display: flex; align-items: center; justify-content: space-between; padding: 0 5%; box-shadow: 0 1px 3px rgba(0,0,0,0.1); z-index: 100; }
			.container { max-width: 1200px; margin: 0 auto; padding: 20px; display: grid; grid-template-columns: 350px 1fr; gap: 30px; }
			@media (max-width: 768px) { .container { grid-template-columns: 1fr; } }
			.card { background: white; padding: 25px; border-radius: 16px; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.1); margin-bottom: 20px; position: relative; }
			
			.gallery { display: flex; overflow-x: auto; gap: 10px; padding-bottom: 10px; scroll-behavior: smooth; }
			.gallery img { width: 100%; height: 250px; object-fit: cover; border-radius: 12px; flex-shrink: 0; }

			.tag { display: inline-block; background: #e0e7ff; color: #4338ca; padding: 4px 10px; border-radius: 20px; font-size: 0.8rem; font-weight: 600; margin-bottom: 10px; }
			
			input, select { width: 100%; padding: 12px; margin-bottom: 12px; border: 1px solid #d1d5db; border-radius: 8px; box-sizing: border-box; }
			.btn-primary { background: var(--primary); color: white; padding: 10px 20px; border-radius: 8px; border: none; cursor: pointer; width: 100%; font-weight: bold; }
			.btn-mpesa { background: var(--mpesa); color: white; padding: 10px; border-radius: 8px; border: none; cursor: pointer; width: 100%; font-weight: bold; margin-top: 5px; }
			.btn-whatsapp { background: var(--whatsapp); color: white; padding: 10px; border-radius: 8px; border: none; cursor: pointer; width: 100%; font-weight: bold; margin-top: 5px; text-decoration: none; display: block; text-align: center; }
			.btn-delete { background: white; color: #ef4444; border: 1px solid #ef4444; padding: 5px 10px; border-radius: 6px; cursor: pointer; float: right; font-size: 0.8rem; margin-top: -5px; }
			
			.booked { opacity: 0.7; border: 2px solid #ccc; background: #f9f9f9; pointer-events: none; }
			.booked::after { content: "⛔ TAKEN"; position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%) rotate(-10deg); font-size: 3rem; font-weight: 900; color: #e53e3e; border: 5px solid #e53e3e; padding: 10px; border-radius: 10px; opacity: 0.8; z-index: 10; }
			
			/* Special style for the owner so they can still see the card even if booked */
			.booked.owner-view { opacity: 1; pointer-events: auto; border: 2px solid #4f46e5; }
			.booked.owner-view::after { display: none; } /* Hide the big TAKEN stamp for the owner */
			
			.tenant-info { background: #fee2e2; color: #991b1b; padding: 10px; border-radius: 8px; margin-top: 10px; font-weight: bold; text-align: center; }
		</style>
	</head>
	<body>
		<div class="navbar">
			<div style="font-size:1.5rem; font-weight:700; color:#4f46e5;">🏠 Nyumba</div>
			<div><span style="color:#666; margin-right:15px;">` + welcomeMsg + `</span>` + navLinks + `</div>
		</div>

		<div class="container">
			<div class="sidebar">
				<div class="card" style="display: ` + addFormDisplay + `;">
					<h3>➕ List Property</h3>
					<input id="loc" type="text" placeholder="Location">
					<select id="type">
						<option value="Bedsitter">Bedsitter</option>
						<option value="One Bedroom">One Bedroom</option>
						<option value="Two Bedroom">Two Bedroom</option>
						<option value="Studio">Studio</option>
					</select>
					<input id="price" type="number" placeholder="Rent (KES)">
					<input id="utils" type="number" placeholder="Bills (KES)">
					<label>📸 Photos (Select Multiple)</label> 
					<input id="photos" type="file" accept="image/*" multiple>
					<input id="details" type="text" placeholder="Description">
					<button class="btn-primary" onclick="uploadHouse()">Post Property</button>
				</div>
				<div class="card">
					<h3>🔍 Search</h3>
					<input id="searchTag" type="text" placeholder="Search...">
					<button class="btn-primary" style="background:#10b981" onclick="fetchHouses()">Filter</button>
				</div>
			</div>
			<div class="main-content" id="results-area"></div>
		</div>
		<div id="toast" style="visibility:hidden; min-width:250px; background:#333; color:#fff; text-align:center; border-radius:8px; padding:16px; position:fixed; z-index:2000; left:50%; bottom:30px; transform:translateX(-50%);">Notification</div>

		<script>
			const isLoggedIn = ` + isLoggedIn + `;
			const currentUsername = "` + currentUsername + `"; // 👈 Pass username to JS

			document.addEventListener("DOMContentLoaded", () => fetchHouses());

			function showToast(msg) {
				const x = document.getElementById("toast"); x.innerText = msg; x.style.visibility = "visible";
				setTimeout(() => { x.style.visibility = "hidden"; }, 3000);
			}

			function fetchHouses() {
				fetch('/houses').then(res => res.json()).then(data => {
					const container = document.getElementById('results-area');
					container.innerHTML = "";
					data.forEach(h => {
						const isOwner = (h.owner === currentUsername);
						
						// CSS Logic: If it's booked, dim it. BUT if I own it, keep it bright so I can manage it.
						let cardClass = "card";
						if (h.is_booked) {
							cardClass = isOwner ? "card booked owner-view" : "card booked";
						}

						// Delete Button (Only for Owner)
						let deleteBtn = "";
						if (isOwner) {
							deleteBtn = '<button class="btn-delete" onclick="deleteHouse(' + h.id + ')">🗑️ Delete</button>';
						}
						
						let whatsappLink = "https://wa.me/" + h.phone + "?text=Hi, is your " + h.type + " in " + h.location + " available?";
						
						let actionArea = "";
						
						// LOGIC: What shows at the bottom of the card?
						if (isOwner && h.is_booked) {
							// If I own it and it's booked -> Show me WHO booked it
							actionArea = '<div class="tenant-info">👤 Booked by: ' + h.tenant_phone + '</div>';
						} else if (isOwner) {
							// If I own it and it's empty -> Just show text
							actionArea = '<p style="color:#666; font-size:0.9rem;">(Your listing is active)</p>';
						} else if (isLoggedIn) {
							// If I'm a tenant -> Show Pay/Chat buttons
							actionArea = '<a href="' + whatsappLink + '" target="_blank" class="btn-whatsapp">💬 Chat on WhatsApp</a>' +
										 '<button class="btn-mpesa" onclick="payWithMpesa(' + h.id + ')">💳 Pay Booking (KES 1,000)</button>';
						} else {
							actionArea = '<a href="/login" style="display:block; text-align:center; margin-top:10px; color:#666;">Login to View Contacts</a>';
						}

						let imagesHtml = (h.image_urls && h.image_urls.length > 0) ? '<div class="gallery">' : '';
						if(h.image_urls) h.image_urls.forEach(url => { imagesHtml += '<img src="' + url + '">'; });
						if(h.image_urls && h.image_urls.length > 0) imagesHtml += '</div>';

						const html = '<div class="' + cardClass + '">' + 
							deleteBtn +
							imagesHtml + 
							'<div style="display:flex; justify-content:space-between; align-items:center;">' + 
								'<h3>' + h.location + '</h3>' + 
								'<span class="tag">' + h.type + '</span>' + 
							'</div>' +
							'<p>Rent: <b>' + h.price + '</b> | Bills: ' + h.utilities + '</p>' +
							'<p>' + h.details + '</p>' +
							actionArea + '</div>';
						container.innerHTML += html;
					});
				});
			}

			function deleteHouse(id) {
				if(!confirm("Are you sure you want to delete this house?")) return;
				fetch('/houses/delete?id=' + id, {method: 'POST'}).then(() => {
					showToast("🗑️ Deleted!");
					fetchHouses();
				});
			}

			function payWithMpesa(id) {
				let phone = prompt("📲 Enter M-Pesa Number (2547...):");
				if (!phone) return;
				showToast("⏳ Sending request...");
				fetch('/pay?id=' + id + '&phone=' + phone, {method: 'POST'})
				.then(res => res.json())
				.then(data => { 
					if(data.ResponseCode === "0") { showToast("✅ Check Phone!"); fetchHouses(); }
					else { showToast("⚠️ Check Console"); }
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

				fetch('/houses/upload', { method: 'POST', body: formData })
				.then(res => {
					if(res.status === 401) showToast("Login Required");
					else { fetchHouses(); showToast("Uploaded!"); }
				});
			}
		</script>
	</body>
	</html>`
	fmt.Fprint(w, html)
}
