package templates

import "fmt"

func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	
	// UPGRADED: Synced theme color to #0a0a0a and added floating badges, premium typography, and image zoom
	return fmt.Sprintf(`<!DOCTYPE html><html><head><title>Nyumba</title><meta name="viewport" content="width=device-width, initial-scale=1"><link rel="manifest" href="/manifest.json"><meta name="theme-color" content="#0a0a0a"><link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet"><script src="https://cdn.tailwindcss.com"></script>
	<style>
		body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: #f8fafc; overflow-x: hidden; }
		.glass-card { background: rgba(30, 41, 59, 0.4); border: 1px solid rgba(255, 255, 255, 0.05); backdrop-filter: blur(16px); }
		.glass-sidebar { background: rgba(15, 23, 42, 0.8); backdrop-filter: blur(20px); border-right: 1px solid rgba(255, 255, 255, 0.05); }
		#toast.hidden { display: none; } 
		
		/* Live Background Blob Animations */
		@keyframes blob { 0%% { transform: translate(0px, 0px) scale(1); } 33%% { transform: translate(150px, -150px) scale(1.2); } 66%% { transform: translate(-100px, 100px) scale(0.8); } 100%% { transform: translate(0px, 0px) scale(1); } }
		.animate-blob { animation: blob 10s infinite alternate ease-in-out; }
		.animation-delay-2000 { animation-delay: 2s; }
		.animation-delay-4000 { animation-delay: 4s; }
	</style></head>
	<body class="h-screen flex flex-col md:flex-row overflow-hidden relative">
		
		<div class="fixed top-[-10%%] left-[-10%%] w-[50vw] h-[50vw] bg-indigo-600/20 rounded-full blur-[120px] -z-10 pointer-events-none animate-blob mix-blend-lighten"></div>
		<div class="fixed bottom-[-10%%] right-[-10%%] w-[40vw] h-[40vw] bg-cyan-500/20 rounded-full blur-[120px] -z-10 pointer-events-none animate-blob animation-delay-2000 mix-blend-lighten"></div>

		<div class="md:hidden flex items-center justify-between p-4 bg-slate-900/80 backdrop-blur-md border-b border-white/5 z-40"><h1 class="text-xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300">Nyumba.</h1><button onclick="toggleMenu()" class="text-white text-2xl px-2">☰</button></div>
		<div id="backdrop" onclick="toggleMenu()" class="fixed inset-0 bg-black/80 z-40 hidden md:hidden transition-opacity"></div>

		<aside id="sidebar" class="fixed inset-y-0 left-0 z-50 w-80 md:w-80 md:static md:flex flex-col h-full transform -translate-x-full md:translate-x-0 transition-transform duration-300 glass-sidebar shadow-2xl">
			<div class="p-8 pb-4 flex justify-between items-center"><div><a href="/" class="text-4xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-white to-slate-400 hover:opacity-80 transition">Nyumba.</a><p class="text-xs text-slate-500 font-medium tracking-widest uppercase mt-2">Curated Living</p></div><button onclick="toggleMenu()" class="md:hidden text-white text-3xl">&times;</button></div>
			<div class="px-6 py-4 space-y-6 flex-1 overflow-y-auto">
				%s 
				<div style="display: %s;" class="glass-card rounded-[2rem] p-6 mb-8 border border-indigo-500/20 shadow-[0_0_30px_rgba(99,102,241,0.1)]">
					<h3 class="text-xs font-extrabold text-indigo-400 uppercase tracking-widest mb-4 flex items-center gap-2"><span class="w-2 h-2 rounded-full bg-indigo-500 animate-pulse"></span> Landlord Hub</h3>
					<div class="space-y-3">
						<input id="building" type="text" placeholder="Apartment Name" class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-2.5 text-sm text-white outline-none focus:border-indigo-500 transition">
						<input id="loc" type="text" placeholder="Location (e.g. Juja)" class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-2.5 text-sm text-white outline-none focus:border-indigo-500 transition">
						<input id="map_url" type="text" placeholder="📍 Google Maps Link" class="w-full bg-slate-900/50 border border-indigo-500/30 rounded-xl px-4 py-2.5 text-sm text-indigo-300 outline-none focus:border-indigo-500 transition">
						<select id="type" class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-2.5 text-sm outline-none text-slate-300"><option>Bedsitter</option><option>One Bedroom</option><option>Two Bedroom</option><option>Studio</option></select>
						<div class="grid grid-cols-2 gap-3"><input id="price" type="number" placeholder="Rent (KES)" class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-2.5 text-sm outline-none focus:border-indigo-500 transition"><input id="utils" type="number" placeholder="Bills" class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-2.5 text-sm outline-none focus:border-indigo-500 transition"></div>
						<input id="photos" type="file" multiple class="text-xs text-slate-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-xs file:font-bold file:bg-indigo-500/10 file:text-indigo-400 hover:file:bg-indigo-500/20 transition cursor-pointer w-full">
						<textarea id="details" placeholder="Beautiful apartment with..." class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-3 text-sm h-20 resize-none outline-none focus:border-indigo-500 transition"></textarea>
						<button onclick="uploadHouse()" class="w-full bg-white hover:bg-slate-200 text-slate-900 font-bold py-3 rounded-xl text-sm transition transform active:scale-95 mt-2">Publish Listing</button>
					</div>
				</div>
				<div class="space-y-4">
					<div><label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1.5 block ml-1">Search Neighborhood</label><input id="searchLoc" onkeyup="fetchHouses()" type="text" placeholder="Try 'Kileleshwa'..." class="w-full bg-slate-900/80 border border-slate-700 rounded-2xl px-5 py-3.5 text-white focus:border-indigo-500 outline-none transition shadow-inner"></div>
					<div><label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1.5 block ml-1">Max Budget</label><input id="searchPrice" onkeyup="fetchHouses()" type="number" placeholder="Any Price" class="w-full bg-slate-900/80 border border-slate-700 rounded-2xl px-5 py-3.5 text-white focus:border-emerald-500 outline-none transition shadow-inner"></div>
				</div>
			</div>
			<div class="p-6 border-t border-white/5 flex items-center justify-between">
				<div class="flex items-center gap-3"><div class="w-10 h-10 rounded-full bg-gradient-to-br from-indigo-500 to-cyan-400 flex items-center justify-center text-sm font-bold shadow-lg">👤</div><div class="text-sm"><div class="font-bold text-white tracking-wide">%s</div></div></div>
				<a href="/logout" class="text-xs font-bold text-slate-400 hover:text-red-400 px-3 py-2 rounded-lg hover:bg-red-500/10 transition">Logout</a>
			</div>
		</aside>

		<main class="flex-1 h-full overflow-y-auto relative z-10">
			<div class="p-4 md:p-8 lg:px-12 max-w-[1600px] mx-auto">
				<header class="flex flex-col md:flex-row md:justify-between md:items-end mb-10 mt-6 md:mt-2">
					<div><h2 class="text-3xl md:text-5xl font-extrabold text-white tracking-tight mb-2">Explore <span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300">Sanctuaries</span></h2><p class="text-slate-400 text-sm md:text-base">Find your next home and connect directly with verified owners.</p></div>
					<div id="offline-badge" class="hidden mt-4 md:mt-0 bg-amber-500/20 text-amber-500 border border-amber-500/50 px-4 py-2 rounded-full text-xs font-bold animate-pulse inline-block">⚠️ OFFLINE MODE</div>
				</header>
				<div id="results-area" class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-8 pb-24"></div>
			</div>
		</main>

		<div id="dashboard-modal" class="fixed inset-0 z-[100] bg-black/80 backdrop-blur-sm hidden flex flex-col items-center justify-center p-4 transition-opacity">
			<div class="bg-slate-900/90 w-full max-w-md rounded-[2.5rem] p-8 border border-white/10 relative shadow-2xl">
				<button onclick="closeDashboard()" class="absolute top-6 right-6 text-slate-400 hover:text-white text-2xl transition">&times;</button>
				<h2 class="text-2xl font-extrabold text-white mb-2">Unlocked Contacts</h2>
				<p class="text-sm text-slate-400 mb-8 leading-relaxed">You have full direct access to these landlords. Skip the agents and negotiate directly.</p>
				<div id="dashboard-list" class="space-y-4 max-h-[60vh] overflow-y-auto pr-2"></div>
			</div>
		</div>

		<div id="gallery-modal" class="fixed inset-0 z-[100] bg-black/95 hidden flex flex-col items-center justify-center p-4">
			<button onclick="closeGallery()" class="absolute top-6 right-6 text-white text-4xl hover:scale-110 transition">&times;</button>
			<img id="gallery-img" src="" class="max-h-[85vh] max-w-full rounded-2xl shadow-[0_0_50px_rgba(0,0,0,0.8)] object-contain mb-6">
			<div class="flex items-center gap-8 bg-slate-900/50 px-6 py-3 rounded-full border border-white/10 backdrop-blur-md">
                <button onclick="navGallery(-1)" class="text-white text-2xl hover:text-indigo-400 transition">❮</button>
                <p id="gallery-counter" class="text-slate-300 font-bold tracking-widest text-sm">1 / 1</p>
                <button onclick="navGallery(1)" class="text-white text-2xl hover:text-indigo-400 transition">❯</button>
            </div>
		</div>

		<div id="toast" class="hidden fixed top-6 left-1/2 -translate-x-1/2 bg-indigo-600/95 backdrop-blur-md px-6 py-3.5 rounded-full text-sm font-bold text-white shadow-2xl shadow-indigo-500/30 z-[60] flex items-center gap-3 transition-all duration-300 border border-indigo-500/50"><span class="text-lg">✨</span> <span id="toast-msg">Notification</span></div>

		<script>
			const isLoggedIn = %s;
			const currentUsername = "%s";
			let houseImages = {}; let currentGalleryID = 0; let galleryIndex = 0; let autoScrollInterval;
			let allHousesData = []; 

			const skeletonHTML = '<div class="glass-card rounded-[2rem] p-5 flex flex-col relative animate-pulse"><div class="w-full h-56 bg-slate-800/50 rounded-3xl mb-5"></div><div class="h-6 bg-slate-800/50 rounded-md w-2/3 mb-3"></div><div class="h-4 bg-slate-800/50 rounded-md w-full mb-2"></div><div class="h-4 bg-slate-800/50 rounded-md w-4/5 mb-6"></div><div class="h-14 bg-slate-800/50 rounded-xl mt-auto"></div></div>';

			document.addEventListener("DOMContentLoaded", () => { 
				const params = new URLSearchParams(window.location.search);
				const loc = params.get('loc');
				if (loc) { document.getElementById('searchLoc').value = loc; }

				fetchHouses(); 
				startAutoScroll(); 
			});

			function openDashboard() {
				const container = document.getElementById('dashboard-list');
				container.innerHTML = "";
				const myHouses = allHousesData.filter(h => h.is_booked === true); 
				if(myHouses.length === 0) {
					container.innerHTML = "<div class='text-center text-slate-500 py-10 font-medium'>No viewing fees paid yet.</div>";
				} else {
					myHouses.forEach(h => {
						let item = '<div class="bg-slate-800/50 p-5 rounded-2xl mb-4 border border-white/5 transition hover:border-emerald-500/30">' +
							'<div class="flex justify-between items-center mb-4"><span class="font-bold text-white text-lg">' + h.building_name + '</span><span class="text-[10px] text-emerald-400 font-extrabold tracking-widest bg-emerald-500/10 px-2 py-1 rounded-md">UNLOCKED</span></div>' +
							'<div class="grid grid-cols-2 gap-3">' +
								'<a href="tel:' + h.phone + '" class="bg-slate-700 hover:bg-slate-600 text-white py-2.5 rounded-xl text-xs font-bold text-center transition flex items-center justify-center gap-2">📞 Call</a>' +
								'<a href="https://wa.me/' + h.phone + '" class="bg-emerald-600 hover:bg-emerald-500 text-white py-2.5 rounded-xl text-xs font-bold text-center transition flex items-center justify-center gap-2 shadow-lg shadow-emerald-500/20">💬 WhatsApp</a>' +
							'</div>' +
						'</div>';
						container.innerHTML += item;
					});
				}
				document.getElementById('dashboard-modal').classList.remove('hidden');
			}
			function closeDashboard() { document.getElementById('dashboard-modal').classList.add('hidden'); }

			function toggleMenu() { document.getElementById('sidebar').classList.toggle('-translate-x-full'); document.getElementById('backdrop').classList.toggle('hidden'); }
			
			function showToast(msg) { 
				const t = document.getElementById("toast"); 
				document.getElementById("toast-msg").innerText = msg; 
				t.classList.remove("hidden"); 
				setTimeout(() => t.classList.add("hidden"), 3000); 
			}

			function openGallery(id) { const images = houseImages[id]; if(!images || images.length === 0) return; currentGalleryID = id; galleryIndex = 0; updateGalleryView(); document.getElementById('gallery-modal').classList.remove('hidden'); }
			function closeGallery() { document.getElementById('gallery-modal').classList.add('hidden'); }
			function navGallery(step) { const images = houseImages[currentGalleryID]; galleryIndex += step; if(galleryIndex >= images.length) galleryIndex = 0; if(galleryIndex < 0) galleryIndex = images.length - 1; updateGalleryView(); }
			function updateGalleryView() { const images = houseImages[currentGalleryID]; document.getElementById('gallery-img').src = images[galleryIndex]; document.getElementById('gallery-counter').innerText = (galleryIndex + 1) + " / " + images.length; }
			function startAutoScroll() { if (autoScrollInterval) clearInterval(autoScrollInterval); autoScrollInterval = setInterval(() => { document.querySelectorAll('[id^="img-"]').forEach(img => { let id = img.id.split('-')[1]; if (document.getElementById('gallery-modal').classList.contains('hidden')) { changeSlide(id, 1); } }); }, 4000); }
			function changeSlide(id, step) { const images = houseImages[id]; if (!images || images.length <= 1) return; let imgEl = document.getElementById('img-' + id); let current = parseInt(imgEl.dataset.index || 0); let next = current + step; if (next >= images.length) next = 0; if (next < 0) next = images.length - 1; imgEl.dataset.index = next; imgEl.src = images[next]; }

			function fetchHouses() {
				const container = document.getElementById('results-area');
				container.innerHTML = skeletonHTML + skeletonHTML + skeletonHTML;
				fetch('/houses').then(res => res.json()).then(data => {
					allHousesData = data; 
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

				if (filtered.length === 0) { 
					container.innerHTML = "<div class='col-span-full flex flex-col items-center justify-center text-center py-24 px-4 bg-slate-800/20 rounded-[3rem] border border-white/5 backdrop-blur-sm mt-8'>" +
						"<div class='text-6xl mb-6 opacity-80'>🏙️</div>" +
						"<h3 class='text-2xl font-extrabold text-white mb-3'>No Sanctuaries Found</h3>" +
						"<p class='text-base text-slate-400 max-w-md mb-8 leading-relaxed'>We are actively vetting properties in this area. Check back soon or try adjusting your filters.</p>" +
						"<button onclick=\"document.getElementById('searchLoc').value=''; document.getElementById('searchPrice').value=''; fetchHouses();\" class='bg-white text-slate-900 px-8 py-3.5 rounded-full text-sm font-bold shadow-lg transition transform hover:-translate-y-0.5 active:scale-95'>Clear Filters</button>" +
					"</div>";
					return; 
				}
				
				filtered.forEach((h) => {
					houseImages[h.id] = h.image_urls;
					let imageSrc = (h.image_urls && h.image_urls.length > 0) ? h.image_urls[0] : 'https://via.placeholder.com/600x400?text=No+Image';
					
					let actionBtn;
					if (h.is_booked) {
						actionBtn = '<button onclick="openDashboard()" class="mt-2 w-full py-3.5 rounded-xl bg-emerald-500/10 text-emerald-400 border border-emerald-500/30 text-xs font-extrabold tracking-widest uppercase hover:bg-emerald-500/20 transition shadow-inner">🔓 Contact Unlocked</button>';
					} else if (isLoggedIn) {
						actionBtn = '<div class="mt-2">' +
							'<p class="text-[10px] text-center text-slate-500 mb-2 uppercase font-extrabold tracking-widest">Unlocks Direct Phone & WhatsApp</p>' +
							'<button onclick="payWithMpesa(' + h.id + ')" class="w-full bg-gradient-to-r from-indigo-500 to-indigo-600 hover:from-indigo-400 hover:to-indigo-500 text-white font-bold py-3.5 rounded-xl shadow-[0_0_20px_rgba(99,102,241,0.2)] transition transform hover:-translate-y-0.5 active:scale-95 flex items-center justify-center gap-2 border border-indigo-400/30">Pay KES 1,000 to View</button>' +
						'</div>';
					} else {
						actionBtn = '<a href="/login" class="block mt-2 w-full py-3.5 rounded-xl bg-slate-800/50 border border-slate-700 hover:bg-slate-700 text-slate-300 hover:text-white text-center text-sm font-bold transition backdrop-blur-sm">Sign in to Unlock Details</a>';
					}

					// UPGRADED JOEDOWNS.AI STYLE CARDS
					const card = document.createElement('div');
					card.className = "glass-card rounded-[2rem] p-5 flex flex-col relative group transition-all duration-300 hover:-translate-y-2 hover:shadow-[0_20px_40px_rgba(0,0,0,0.5)] border border-white/5 hover:border-indigo-500/30";
					card.innerHTML = 
						'<div class="w-full h-56 bg-slate-800 rounded-3xl overflow-hidden relative mb-5 cursor-pointer shadow-inner" onclick="openGallery(' + h.id + ')">' +
							'<img id="img-' + h.id + '" src="' + imageSrc + '" class="w-full h-full object-cover transition-transform duration-700 ease-in-out group-hover:scale-110">' +
							'<div class="absolute inset-0 bg-gradient-to-t from-slate-900/60 to-transparent pointer-events-none"></div>' +
							'<div class="absolute top-4 right-4 bg-slate-900/90 backdrop-blur-xl px-4 py-1.5 rounded-full border border-white/10 text-white font-extrabold text-sm shadow-xl tracking-wide">KES ' + h.price.toLocaleString() + '</div>' +
							'<div class="absolute bottom-4 left-4 bg-indigo-600/90 backdrop-blur-xl px-3.5 py-1.5 rounded-full border border-white/10 text-white text-[11px] font-bold tracking-widest uppercase shadow-xl flex items-center gap-1">📍 ' + h.location + '</div>' +
						'</div>' +
						'<div class="flex-1">' + 
							'<h3 class="text-2xl font-bold text-white mb-2 leading-tight tracking-tight">' + h.building_name + '</h3>' + 
							'<p class="text-slate-400 text-sm line-clamp-2 leading-relaxed font-light">' + h.details + '</p>' + 
						'</div>' +
						'<div class="mt-4 pt-4 border-t border-white/10">' +
						actionBtn +
						'</div>';
					
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
				let phone = prompt("Enter M-Pesa Number to Pay Viewing Fee (e.g., 0712345678):");
				if (!phone) return;
				
				showToast("Requesting M-Pesa...");
				fetch('/pay?id=' + id + '&phone=' + phone, {method: 'POST'})
				.then(res => res.json())
				.then(data => { 
					if(data.ResponseCode === "0") { 
						showToast("📲 check your phone to enter PIN!"); 
					} else { 
						showToast(data.CustomerMessage || "Connection Failed"); 
					} 
				})
				.catch(err => { 
					console.error(err);
					showToast("System Error"); 
				});
			}
		</script>
	</body>
	</html>`, myHubButton, landlordPanelDisplay, currentUsername, isLoggedIn, currentUsername)
}
func GetLandingHTML() string {
	return `<!DOCTYPE html><html><head><title>Nyumba • Curated Living</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet">
	<script src="https://cdn.tailwindcss.com"></script>
	<style>
		html { scroll-behavior: smooth; }
		body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: #f8fafc; overflow-x: hidden; }
		.glass-pill { background: rgba(15, 23, 42, 0.6); backdrop-filter: blur(20px); border: 1px solid rgba(255,255,255,0.08); box-shadow: 0 20px 40px rgba(0,0,0,0.4); }
		@keyframes float { 0% { transform: translateY(0px) rotateX(4deg); } 50% { transform: translateY(-10px) rotateX(1deg); } 100% { transform: translateY(0px) rotateX(4deg); } }
		.animate-float { animation: float 6s ease-in-out infinite; perspective: 1000px; }
		
		/* Infinite Marquee Animation */
		@keyframes scroll { 0% { transform: translateX(0); } 100% { transform: translateX(-50%); } }
		.animate-marquee { display: flex; width: max-content; animation: scroll 30s linear infinite; }
		.animate-marquee:hover { animation-play-state: paused; }
		
		/* Live Background Blob Animations */
		@keyframes blob {
			0% { transform: translate(0px, 0px) scale(1); }
			33% { transform: translate(150px, -150px) scale(1.2); }
			66% { transform: translate(-100px, 100px) scale(0.8); }
			100% { transform: translate(0px, 0px) scale(1); }
		}
		.animate-blob { animation: blob 10s infinite alternate ease-in-out; }
		.animation-delay-2000 { animation-delay: 2s; }
		.animation-delay-4000 { animation-delay: 4s; }
	</style>
	</head>
	<body class="antialiased selection:bg-indigo-500/30 relative">
		
		<div class="fixed top-[-10%] left-[-10%] w-[50vw] h-[50vw] bg-indigo-600/40 rounded-full blur-[120px] -z-10 pointer-events-none animate-blob mix-blend-lighten"></div>
		<div class="fixed top-[10%] right-[-10%] w-[40vw] h-[40vw] bg-cyan-500/40 rounded-full blur-[120px] -z-10 pointer-events-none animate-blob animation-delay-2000 mix-blend-lighten"></div>
		<div class="fixed bottom-[-10%] left-[20%] w-[50vw] h-[50vw] bg-emerald-500/30 rounded-full blur-[120px] -z-10 pointer-events-none animate-blob animation-delay-4000 mix-blend-lighten"></div>

		<div class="fixed top-6 left-0 w-full flex justify-center z-50 px-4">
			<nav class="glass-pill rounded-full px-6 py-3 w-full max-w-4xl flex items-center justify-between transition-all">
				<div class="text-xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-white to-slate-400">Nyumba.</div>
				<div class="hidden md:flex items-center gap-8 text-sm font-semibold text-slate-300">
					<a href="#how" class="hover:text-white transition">How it Works</a>
					<a href="#neighborhoods" class="hover:text-white transition">Neighborhoods</a>
					<a href="/signup" class="hover:text-white transition">For Landlords</a>
				</div>
				<div class="flex items-center gap-4">
					<a href="/login" class="text-sm font-bold text-slate-300 hover:text-white transition">Sign In</a>
					<a href="/explore" class="bg-indigo-600 hover:bg-indigo-500 text-white text-sm font-bold py-2 px-6 rounded-full shadow-lg shadow-indigo-500/25 transition transform active:scale-95">Explore</a>
				</div>
			</nav>
		</div>

		<main class="pt-40 pb-16 px-6 max-w-5xl mx-auto flex flex-col items-center text-center">
			
			<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full border border-slate-700 bg-slate-800/50 text-slate-300 text-xs font-bold tracking-widest uppercase mb-8 backdrop-blur-sm">
				<span class="relative flex h-2 w-2"><span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span><span class="relative inline-flex rounded-full h-2 w-2 bg-emerald-500"></span></span>
				Verified Listings Only
			</div>
			
			<h1 class="text-6xl md:text-8xl font-extrabold tracking-tight mb-6 leading-[1.05]">
				Find Your Sanctuary. <br/>
				<span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300">Simplified.</span>
			</h1>
			
			<p class="text-slate-400 text-lg md:text-xl max-w-2xl mx-auto mb-10 leading-relaxed font-light">
				An exclusive platform connecting serious renters with verified landlords. No agents. No endless scrolling. Just your next home.
			</p>
			
			<div class="flex flex-col sm:flex-row gap-4 w-full justify-center relative z-20 mb-12">
				<a href="/explore" class="group relative inline-flex items-center justify-center bg-white text-slate-900 text-lg font-bold py-4 px-10 rounded-full transition-all duration-300 hover:-translate-y-1 hover:shadow-[0_0_30px_rgba(255,255,255,0.2)]">
					Start Your Search
					<span class="ml-2 group-hover:translate-x-1 transition-transform duration-300">→</span>
				</a>
			</div>

			<div id="neighborhoods" class="w-full max-w-4xl border-t border-white/10 pt-8 mt-4 scroll-mt-24">
				<p class="text-xs text-slate-500 uppercase tracking-widest font-bold mb-4">Popular Neighborhoods</p>
				
				<div class="relative overflow-hidden w-full flex">
					<div class="absolute inset-y-0 left-0 w-16 bg-gradient-to-r from-[#0a0a0a] to-transparent z-10 pointer-events-none"></div>
					<div class="absolute inset-y-0 right-0 w-16 bg-gradient-to-l from-[#0a0a0a] to-transparent z-10 pointer-events-none"></div>
					
					<div class="animate-marquee gap-3">
						<a href="/explore?loc=kileleshwa" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Kileleshwa</a>
						<a href="/explore?loc=kilimani" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Kilimani</a>
						<a href="/explore?loc=thika" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Thika</a>
						<a href="/explore?loc=juja" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Juja</a>
						<a href="/explore?loc=roysambu" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Roysambu</a>
						<a href="/explore?loc=westlands" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Westlands</a>
						<a href="/explore?loc=ruiru" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Ruiru</a>
						<a href="/explore?loc=karen" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Karen</a>

						<a href="/explore?loc=kileleshwa" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Kileleshwa</a>
						<a href="/explore?loc=kilimani" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Kilimani</a>
						<a href="/explore?loc=thika" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Thika</a>
						<a href="/explore?loc=juja" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Juja</a>
						<a href="/explore?loc=roysambu" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Roysambu</a>
						<a href="/explore?loc=westlands" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Westlands</a>
						<a href="/explore?loc=ruiru" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Ruiru</a>
						<a href="/explore?loc=karen" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Karen</a>
					</div>
				</div>
			</div>
		</main>

		<section id="how" class="py-24 px-6 relative z-10 border-t border-white/5 bg-slate-900/20 scroll-mt-10">
			<div class="max-w-5xl mx-auto text-center">
				<h2 class="text-3xl md:text-5xl font-bold mb-16 text-white">How it Works</h2>
				<div class="grid grid-cols-1 md:grid-cols-3 gap-8">
					<div class="bg-slate-800/40 p-8 rounded-3xl border border-white/5 backdrop-blur-sm transition hover:-translate-y-2">
						<div class="text-4xl mb-4">🔍</div>
						<h3 class="text-xl font-bold text-white mb-2">1. Discover</h3>
						<p class="text-slate-400 text-sm">Browse our curated list of high-quality apartments in top neighborhoods.</p>
					</div>
					<div class="bg-slate-800/40 p-8 rounded-3xl border border-white/5 backdrop-blur-sm transition hover:-translate-y-2">
						<div class="text-4xl mb-4">💳</div>
						<h3 class="text-xl font-bold text-white mb-2">2. Unlock</h3>
						<p class="text-slate-400 text-sm">Pay a secure KES 1,000 viewing fee via M-Pesa to show you're serious.</p>
					</div>
					<div class="bg-slate-800/40 p-8 rounded-3xl border border-white/5 backdrop-blur-sm transition hover:-translate-y-2">
						<div class="text-4xl mb-4">🔑</div>
						<h3 class="text-xl font-bold text-white mb-2">3. Connect</h3>
						<p class="text-slate-400 text-sm">Get the landlord's direct phone number instantly. No middlemen.</p>
					</div>
				</div>
			</div>
		</section>

		<section class="pb-20 pt-10 px-6 border-t border-white/5">
			<div class="max-w-5xl mx-auto grid grid-cols-2 md:grid-cols-4 gap-4">
				<div class="bg-slate-900/40 border border-white/5 rounded-3xl p-6 text-center z-10 relative">
					<h3 class="text-3xl font-bold text-white mb-1">500+</h3>
					<p class="text-[10px] text-slate-500 uppercase tracking-widest font-bold">Properties</p>
				</div>
				<div class="bg-slate-900/40 border border-white/5 rounded-3xl p-6 text-center z-10 relative">
					<h3 class="text-3xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-cyan-400 mb-1">100%</h3>
					<p class="text-[10px] text-slate-500 uppercase tracking-widest font-bold">Verified</p>
				</div>
				<div class="bg-slate-900/40 border border-white/5 rounded-3xl p-6 text-center z-10 relative">
					<h3 class="text-3xl font-bold text-white mb-1">24/7</h3>
					<p class="text-[10px] text-slate-500 uppercase tracking-widest font-bold">Direct Access</p>
				</div>
				<div class="bg-slate-900/40 border border-white/5 rounded-3xl p-6 text-center z-10 relative">
					<h3 class="text-3xl font-bold text-white mb-1">Secure</h3>
					<p class="text-[10px] text-slate-500 uppercase tracking-widest font-bold">M-Pesa</p>
				</div>
			</div>
		</section>

	</body></html>`
}