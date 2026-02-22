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
		body { font-family: 'Outfit', sans-serif; background: #0f172a; color: #f8fafc; overflow-x: hidden; }
		.glass-nav { background: rgba(15, 23, 42, 0.7); backdrop-filter: blur(16px); border-bottom: 1px solid rgba(255,255,255,0.05); }
		/* Floating animation for the mockup */
		@keyframes float { 0% { transform: translateY(0px) perspective(1000px) rotateX(4deg); } 50% { transform: translateY(-15px) perspective(1000px) rotateX(1deg); } 100% { transform: translateY(0px) perspective(1000px) rotateX(4deg); } }
		.animate-float { animation: float 6s ease-in-out infinite; }
	</style>
	</head>
	<body class="antialiased selection:bg-indigo-500/30">
		
		<div class="fixed top-[-20%] left-[-10%] w-[50vw] h-[50vw] bg-indigo-600/20 rounded-full blur-[120px] -z-10 pointer-events-none"></div>
		<div class="fixed top-[20%] right-[-10%] w-[40vw] h-[40vw] bg-cyan-500/10 rounded-full blur-[120px] -z-10 pointer-events-none"></div>

		<nav class="fixed w-full z-50 glass-nav transition-all">
			<div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
				<div class="text-2xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300">Nyumba.</div>
				<div class="flex items-center gap-4">
					<a href="/login" class="text-sm font-bold text-slate-300 hover:text-white transition">Login</a>
					<a href="/explore" class="bg-indigo-600 hover:bg-indigo-500 text-white text-sm font-bold py-2.5 px-6 rounded-full shadow-lg shadow-indigo-500/25 transition transform hover:-translate-y-0.5 active:scale-95">Explore Homes</a>
				</div>
			</div>
		</nav>

		<main class="pt-36 pb-16 px-6 max-w-7xl mx-auto flex flex-col items-center text-center">
			
			<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full border border-indigo-500/30 bg-indigo-500/10 text-indigo-300 text-xs font-bold tracking-widest uppercase mb-8 backdrop-blur-sm">
				<span class="relative flex h-2 w-2"><span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-indigo-400 opacity-75"></span><span class="relative inline-flex rounded-full h-2 w-2 bg-indigo-500"></span></span>
				Now Live in Kenya
			</div>
			
			<h1 class="text-5xl md:text-7xl lg:text-8xl font-extrabold tracking-tight mb-6 leading-[1.1]">
				Find Your Sanctuary. <br/>
				<span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 via-cyan-300 to-emerald-300">Without the Hassle.</span>
			</h1>
			
			<p class="text-slate-400 text-lg md:text-xl max-w-2xl mx-auto mb-10 leading-relaxed font-light">
				Nyumba is an exclusive platform connecting serious renters with verified landlords. Say goodbye to endless scrolling, scams, and agent fees.
			</p>
			
			<div class="flex flex-col sm:flex-row gap-4 w-full justify-center relative z-20">
				<a href="/explore" class="group relative inline-flex items-center justify-center bg-white text-slate-900 text-lg font-bold py-4 px-8 rounded-2xl transition-all duration-300 hover:-translate-y-1 hover:shadow-[0_0_40px_rgba(255,255,255,0.3)]">
					Browse Listings 
					<span class="ml-2 group-hover:translate-x-1 transition-transform duration-300">→</span>
				</a>
				<a href="/signup" class="bg-slate-800/80 backdrop-blur-md border border-slate-700 hover:bg-slate-700 hover:border-slate-600 text-white text-lg font-bold py-4 px-8 rounded-2xl transition-all duration-300 hover:-translate-y-1 hover:shadow-xl">
					I am a Landlord
				</a>
			</div>

			<div class="mt-20 w-full max-w-4xl mx-auto relative z-10 hidden md:block">
				<div class="absolute inset-0 bg-gradient-to-r from-indigo-500 to-cyan-400 blur-3xl opacity-20 rounded-full animate-pulse"></div>
				
				<div class="relative bg-slate-900/60 backdrop-blur-2xl border border-white/10 p-2 rounded-[2rem] shadow-2xl animate-float">
					<div class="bg-[#0f172a] rounded-[1.5rem] border border-white/5 overflow-hidden flex flex-col h-64 shadow-[inset_0_0_40px_rgba(0,0,0,0.5)]">
						<div class="bg-slate-800/50 px-4 py-3 flex items-center gap-2 border-b border-white/5">
							<div class="w-3 h-3 rounded-full bg-red-500/80"></div>
							<div class="w-3 h-3 rounded-full bg-amber-500/80"></div>
							<div class="w-3 h-3 rounded-full bg-emerald-500/80"></div>
							<div class="mx-auto bg-slate-900/80 text-slate-500 text-xs py-1 px-10 rounded-full border border-white/5 flex items-center gap-2">🔒 nyumba-app.onrender.com</div>
						</div>
						<div class="p-6 flex-1 flex flex-col gap-4">
							<div class="bg-slate-800/50 border border-white/5 rounded-2xl p-4 flex items-center gap-4 shadow-inner">
								<div class="text-2xl">📍</div>
								<div class="flex-1 h-6 bg-slate-700/50 rounded-lg w-1/2"></div>
								<div class="w-24 h-10 bg-indigo-600 rounded-xl shadow-lg shadow-indigo-500/20"></div>
							</div>
							<div class="grid grid-cols-3 gap-4 flex-1">
								<div class="bg-slate-800/40 rounded-2xl border border-white/5"></div>
								<div class="bg-slate-800/40 rounded-2xl border border-white/5"></div>
								<div class="bg-slate-800/40 rounded-2xl border border-white/5"></div>
							</div>
						</div>
					</div>
				</div>
			</div>

			<div class="mt-16 md:mt-24 grid grid-cols-2 md:grid-cols-4 gap-4 w-full max-w-5xl mx-auto z-20 relative">
				<div class="bg-slate-800/40 border border-white/5 backdrop-blur-md rounded-3xl p-6 md:p-8 transition-all hover:bg-slate-800/60 hover:-translate-y-1 hover:border-indigo-500/30">
					<h3 class="text-3xl md:text-4xl font-black text-white mb-2">500+</h3>
					<p class="text-[10px] md:text-xs text-slate-400 uppercase tracking-widest font-bold">Renters</p>
				</div>
				<div class="bg-slate-800/40 border border-white/5 backdrop-blur-md rounded-3xl p-6 md:p-8 transition-all hover:bg-slate-800/60 hover:-translate-y-1 hover:border-emerald-500/30">
					<h3 class="text-3xl md:text-4xl font-black text-white mb-2 text-transparent bg-clip-text bg-gradient-to-br from-emerald-400 to-cyan-400">100%</h3>
					<p class="text-[10px] md:text-xs text-slate-400 uppercase tracking-widest font-bold">Verified</p>
				</div>
				<div class="bg-slate-800/40 border border-white/5 backdrop-blur-md rounded-3xl p-6 md:p-8 transition-all hover:bg-slate-800/60 hover:-translate-y-1 hover:border-amber-500/30">
					<h3 class="text-3xl md:text-4xl font-black text-white mb-2">24/7</h3>
					<p class="text-[10px] md:text-xs text-slate-400 uppercase tracking-widest font-bold">Direct Access</p>
				</div>
				<div class="bg-slate-800/40 border border-white/5 backdrop-blur-md rounded-3xl p-6 md:p-8 transition-all hover:bg-slate-800/60 hover:-translate-y-1 hover:border-cyan-500/30">
					<h3 class="text-3xl md:text-4xl font-black text-white mb-2">Secure</h3>
					<p class="text-[10px] md:text-xs text-slate-400 uppercase tracking-widest font-bold">M-Pesa Pay</p>
				</div>
			</div>
		</main>

		<footer class="text-center py-10 mt-10 border-t border-white/5 text-slate-500 text-xs font-medium tracking-wide">
			<p>&copy; 2026 Nyumba. Curated Living in Kenya.</p>
		</footer>
	</body></html>`
}