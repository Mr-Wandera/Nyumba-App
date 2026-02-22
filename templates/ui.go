package templates

import "fmt"

func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	
	return fmt.Sprintf(`<!DOCTYPE html><html><head><title>Nyumba</title><meta name="viewport" content="width=device-width, initial-scale=1"><link rel="manifest" href="/manifest.json"><meta name="theme-color" content="#0f172a"><link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet"><script src="https://cdn.tailwindcss.com"></script>
	<style>
		body { font-family: 'Outfit', sans-serif; background: #0f172a; color: #f8fafc; }
		.glass-card { background: rgba(30, 41, 59, 0.7); border: 1px solid rgba(255, 255, 255, 0.1); backdrop-filter: blur(10px); transform-style: preserve-3d; transition: transform 0.1s ease-out; }
		.glass-card > .absolute, .glass-card > div.mt-4 { transform: translateZ(40px); box-shadow: 0 10px 20px rgba(0,0,0,0.3); }
		.glass-sidebar { background: #1e293b; border-right: 1px solid rgba(255, 255, 255, 0.05); }
		#toast.hidden { display: none; } 
	</style></head>
	<body class="h-screen flex flex-col md:flex-row overflow-hidden">
		<div class="md:hidden flex items-center justify-between p-4 bg-slate-900 border-b border-white/5 z-40"><h1 class="text-xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300">Nyumba.</h1><button onclick="toggleMenu()" class="text-white text-2xl px-2">☰</button></div>
		<div id="backdrop" onclick="toggleMenu()" class="fixed inset-0 bg-black/80 z-40 hidden md:hidden transition-opacity"></div>

		<aside id="sidebar" class="fixed inset-y-0 left-0 z-50 w-80 bg-[#1e293b] md:bg-transparent md:static md:flex flex-col h-full transform -translate-x-full md:translate-x-0 transition-transform duration-300 glass-sidebar">
			<div class="p-8 pb-4 flex justify-between items-center"><div><h1 class="text-4xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300">Nyumba.</h1><p class="text-xs text-slate-500 font-medium tracking-widest uppercase mt-2">Curated Living</p></div><button onclick="toggleMenu()" class="md:hidden text-white text-3xl">&times;</button></div>
			<div class="px-6 py-4 space-y-6 flex-1 overflow-y-auto">
				%s 
				<div style="display: %s;" class="glass-card rounded-2xl p-5 mb-8">
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
				<div class="flex items-center gap-3"><div class="w-8 h-8 rounded-full bg-indigo-500 flex items-center justify-center text-xs font-bold">👤</div><div class="text-sm"><div class="font-bold text-white">%s</div></div></div>
				<a href="/logout" class="text-sm font-bold text-red-400 border border-red-500/30 px-3 py-1 rounded-full hover:bg-red-500/10 transition">Logout</a>
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

		<div id="toast" class="hidden fixed top-6 left-1/2 -translate-x-1/2 bg-indigo-600 px-6 py-3 rounded-full text-sm font-bold text-white shadow-2xl z-[60] flex items-center gap-2 transition-all duration-300"><span class="text-lg">✨</span> <span id="toast-msg">Notification</span></div>
		<script>
			const isLoggedIn = %s;
			const currentUsername = "%s";
			let houseImages = {}; let currentGalleryID = 0; let galleryIndex = 0; let autoScrollInterval;
			let allHousesData = []; 

			const skeletonHTML = '<div class="glass-card rounded-3xl p-4 flex flex-col relative animate-pulse border border-white/5"><div class="w-full h-48 bg-slate-800 rounded-2xl mb-4"></div><div class="h-4 bg-slate-800 rounded w-1/3 mb-2"></div><div class="h-6 bg-slate-800 rounded w-3/4 mb-4"></div><div class="h-12 bg-slate-800 rounded-xl mt-4"></div></div>';

			document.addEventListener("DOMContentLoaded", () => { fetchHouses(); startAutoScroll(); });

			function openDashboard() {
				const container = document.getElementById('dashboard-list');
				container.innerHTML = "";
				const myHouses = allHousesData.filter(h => h.is_booked === true); 
				if(myHouses.length === 0) {
					container.innerHTML = "<div class='text-center text-slate-500 py-10'>You haven't paid any viewing fees yet.</div>";
				} else {
					myHouses.forEach(h => {
						let item = '<div class="bg-slate-800 p-4 rounded-xl mb-3 border border-white/10">' +
							'<div class="flex justify-between mb-2"><span class="font-bold text-white">' + h.building_name + '</span><span class="text-xs text-emerald-400 font-bold">UNLOCKED</span></div>' +
							'<div class="grid grid-cols-2 gap-2">' +
								'<a href="tel:' + h.phone + '" class="bg-slate-700 hover:bg-slate-600 text-white py-2 rounded-lg text-xs font-bold text-center">📞 Call Owner</a>' +
								'<a href="https://wa.me/' + h.phone + '" class="bg-emerald-600 hover:bg-emerald-500 text-white py-2 rounded-lg text-xs font-bold text-center">💬 WhatsApp</a>' +
							'</div>' +
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
					const rotateX = ((y - rect.height/2) / (rect.height/2)) * -5; 
					const rotateY = ((x - rect.width/2) / (rect.width/2)) * 5;
					card.style.transform = "perspective(1000px) rotateX(" + rotateX + "deg) rotateY(" + rotateY + "deg) scale(1.02)";
				});
				card.addEventListener('mouseleave', () => { card.style.transform = "perspective(1000px) rotateX(0) rotateY(0) scale(1)"; });
			}

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
			function startAutoScroll() { if (autoScrollInterval) clearInterval(autoScrollInterval); autoScrollInterval = setInterval(() => { document.querySelectorAll('[id^="img-"]').forEach(img => { let id = img.id.split('-')[1]; if (document.getElementById('gallery-modal').classList.contains('hidden')) { changeSlide(id, 1); } }); }, 3500); }
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

				// UPDATED: Professional Empty State UI
				if (filtered.length === 0) { 
					container.innerHTML = "<div class='col-span-full flex flex-col items-center justify-center text-center py-20 px-4 bg-slate-800/30 rounded-3xl border border-white/5'>" +
						"<div class='text-6xl mb-4'>🏙️</div>" +
						"<h3 class='text-xl font-bold text-white mb-2'>No properties here yet</h3>" +
						"<p class='text-sm text-slate-400 max-w-md mb-6'>We are currently verifying landlords in this area. Check back soon or adjust your search filters.</p>" +
						"<button onclick=\"document.getElementById('searchLoc').value=''; document.getElementById('searchPrice').value=''; fetchHouses();\" class='bg-indigo-600/20 text-indigo-400 border border-indigo-500/30 px-6 py-2 rounded-full text-sm font-bold hover:bg-indigo-600/40 transition'>Clear Filters</button>" +
					"</div>";
					return; 
				}
				
				filtered.forEach((h) => {
					houseImages[h.id] = h.image_urls;
					let imageSrc = (h.image_urls && h.image_urls.length > 0) ? h.image_urls[0] : 'https://via.placeholder.com/600x400?text=No+Image';
					
					let actionBtn;
					if (h.is_booked) {
						actionBtn = '<button onclick="openDashboard()" class="mt-4 w-full py-3 rounded-xl bg-emerald-500/10 text-emerald-400 border border-emerald-500/50 text-xs font-bold tracking-widest uppercase">🔓 Contact Unlocked</button>';
					} else if (isLoggedIn) {
						// UPDATED: Payment Clarity
						actionBtn = '<div class="mt-4">' +
							'<p class="text-[10px] text-center text-slate-400 mb-2 uppercase font-bold tracking-wider">Unlocks Direct Phone & WhatsApp</p>' +
							'<button onclick="payWithMpesa(' + h.id + ')" class="w-full bg-gradient-to-r from-indigo-600 to-indigo-700 hover:from-indigo-500 hover:to-indigo-600 text-white font-bold py-3 rounded-xl shadow-lg shadow-indigo-500/30 transition transform active:scale-95 flex items-center justify-center gap-2">Pay via M-Pesa (KES 1,000)</button>' +
						'</div>';
					} else {
						actionBtn = '<a href="/login" class="block mt-4 w-full py-3 rounded-xl bg-slate-800 hover:bg-slate-700 text-white text-center text-xs font-bold transition">Login to Unlock Details</a>';
					}

					const card = document.createElement('div');
					card.className = "glass-card rounded-3xl p-4 flex flex-col relative group transition hover:-translate-y-1 hover:shadow-2xl";
					card.innerHTML = 
						'<div class="w-full h-48 bg-slate-800 rounded-2xl overflow-hidden relative mb-4 cursor-pointer" onclick="openGallery(' + h.id + ')"><img id="img-' + h.id + '" src="' + imageSrc + '" class="w-full h-full object-cover transition duration-700 ease-out"></div>' +
						'<div class="flex-1">' + 
							'<h3 class="text-xl font-bold text-white">' + h.building_name + '</h3>' + 
							'<p class="text-xs text-slate-400 mb-2">📍 ' + h.location + '</p>' + 
							'<p class="text-slate-400 text-sm line-clamp-2">' + h.details + '</p>' + 
						'</div>' +
						'<div class="mt-4 pt-4 border-t border-white/5 flex items-end justify-between"><div><p class="text-[10px] text-slate-500 uppercase font-bold">Monthly Rent</p><p class="text-xl font-bold text-white">KES ' + h.price.toLocaleString() + '</p></div></div>' +
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
		body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: #f8fafc; overflow-x: hidden; }
		.glass-pill { background: rgba(15, 23, 42, 0.6); backdrop-filter: blur(20px); border: 1px solid rgba(255,255,255,0.08); box-shadow: 0 20px 40px rgba(0,0,0,0.4); }
		@keyframes float { 0% { transform: translateY(0px) rotateX(4deg); } 50% { transform: translateY(-10px) rotateX(1deg); } 100% { transform: translateY(0px) rotateX(4deg); } }
		.animate-float { animation: float 6s ease-in-out infinite; perspective: 1000px; }
		
		/* Infinite Marquee Animation */
		@keyframes scroll { 0% { transform: translateX(0); } 100% { transform: translateX(-50%); } }
		.animate-marquee { display: flex; width: max-content; animation: scroll 30s linear infinite; }
		.animate-marquee:hover { animation-play-state: paused; } /* Pauses when user hovers over a chip */
	</style>
	</head>
	<body class="antialiased selection:bg-indigo-500/30">
		
		<div class="fixed top-[-10%] left-[10%] w-[40vw] h-[40vw] bg-indigo-600/10 rounded-full blur-[100px] -z-10 pointer-events-none"></div>
		<div class="fixed bottom-[10%] right-[10%] w-[30vw] h-[30vw] bg-cyan-500/10 rounded-full blur-[100px] -z-10 pointer-events-none"></div>

		<div class="fixed top-6 left-0 w-full flex justify-center z-50 px-4">
			<nav class="glass-pill rounded-full px-6 py-3 w-full max-w-4xl flex items-center justify-between transition-all">
				<div class="text-xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-white to-slate-400">Nyumba.</div>
				<div class="hidden md:flex items-center gap-8 text-sm font-semibold text-slate-300">
					<a href="#" class="hover:text-white transition">How it Works</a>
					<a href="#" class="hover:text-white transition">Neighborhoods</a>
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

			<div class="w-full max-w-4xl border-t border-white/10 pt-8 mt-4">
				<p class="text-xs text-slate-500 uppercase tracking-widest font-bold mb-4">Popular Neighborhoods</p>
				
				<div class="relative overflow-hidden w-full flex">
					<div class="absolute inset-y-0 left-0 w-16 bg-gradient-to-r from-[#0a0a0a] to-transparent z-10 pointer-events-none"></div>
					<div class="absolute inset-y-0 right-0 w-16 bg-gradient-to-l from-[#0a0a0a] to-transparent z-10 pointer-events-none"></div>
					
					<div class="animate-marquee gap-3">
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Kileleshwa</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Kilimani</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Thika</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Juja</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Roysambu</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Westlands</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Ruiru</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Karen</a>

						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Kileleshwa</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Kilimani</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Thika</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Juja</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Roysambu</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Westlands</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Ruiru</a>
						<a href="/explore" class="shrink-0 px-5 py-2.5 rounded-full bg-slate-800/40 border border-white/5 text-slate-300 text-sm font-medium hover:bg-indigo-600/20 hover:border-indigo-500/30 transition backdrop-blur-sm cursor-pointer">📍 Karen</a>
					</div>
				</div>
			</div>
		</main>

		<section class="pb-20 px-6">
			<div class="max-w-5xl mx-auto grid grid-cols-2 md:grid-cols-4 gap-4">
				<div class="bg-slate-900/40 border border-white/5 rounded-3xl p-6 text-center">
					<h3 class="text-3xl font-bold text-white mb-1">500+</h3>
					<p class="text-[10px] text-slate-500 uppercase tracking-widest font-bold">Properties</p>
				</div>
				<div class="bg-slate-900/40 border border-white/5 rounded-3xl p-6 text-center">
					<h3 class="text-3xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-cyan-400 mb-1">100%</h3>
					<p class="text-[10px] text-slate-500 uppercase tracking-widest font-bold">Verified</p>
				</div>
				<div class="bg-slate-900/40 border border-white/5 rounded-3xl p-6 text-center">
					<h3 class="text-3xl font-bold text-white mb-1">24/7</h3>
					<p class="text-[10px] text-slate-500 uppercase tracking-widest font-bold">Direct Access</p>
				</div>
				<div class="bg-slate-900/40 border border-white/5 rounded-3xl p-6 text-center">
					<h3 class="text-3xl font-bold text-white mb-1">Secure</h3>
					<p class="text-[10px] text-slate-500 uppercase tracking-widest font-bold">M-Pesa</p>
				</div>
			</div>
		</section>

	</body></html>`
}