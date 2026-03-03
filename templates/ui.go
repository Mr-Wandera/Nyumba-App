package templates

import "fmt"

func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	return fmt.Sprintf(`<!DOCTYPE html><html><head><title>Nyumba</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet">
	<script src="https://cdn.tailwindcss.com"></script>
	<style>
		body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: #f8fafc; overflow-x: hidden; }
		.glass-card { background: rgba(30, 41, 59, 0.4); border: 1px solid rgba(255, 255, 255, 0.05); backdrop-filter: blur(16px); }
		.glass-sidebar { background: rgba(15, 23, 42, 0.8); backdrop-filter: blur(20px); border-right: 1px solid rgba(255, 255, 255, 0.05); }
		.detail-pane { background: rgba(15, 23, 42, 0.98); backdrop-filter: blur(30px); border-left: 1px solid rgba(255,255,255,0.1); }
		@keyframes slideIn { from { transform: translateX(100%%); } to { transform: translateX(0); } }
		.animate-slide { animation: slideIn 0.4s cubic-bezier(0.16, 1, 0.3, 1); }
		#toast.hidden { display: none; }
		@keyframes blob { 0%% { transform: translate(0px, 0px) scale(1); } 33%% { transform: translate(150px, -150px) scale(1.2); } 66%% { transform: translate(-100px, 100px) scale(0.8); } 100%% { transform: translate(0px, 0px) scale(1); } }
		.animate-blob { animation: blob 10s infinite alternate ease-in-out; }
	</style></head>
	<body class="h-screen flex flex-col md:flex-row overflow-hidden relative">
		<div class="fixed top-[-10%%] left-[-10%%] w-[50vw] h-[50vw] bg-indigo-600/20 rounded-full blur-[120px] -z-10 animate-blob"></div>
		
		<aside id="sidebar" class="fixed inset-y-0 left-0 z-50 w-80 md:static md:flex flex-col h-full transform -translate-x-full md:translate-x-0 transition-transform duration-300 glass-sidebar">
			<div class="p-8 pb-4 flex justify-between items-center">
				<div><a href="/" class="text-4xl font-extrabold tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-white to-slate-400">Nyumba.</a></div>
			</div>
			<div class="px-6 py-4 space-y-6 flex-1 overflow-y-auto">
				%s 
				<div style="display: %s;" class="glass-card rounded-[2rem] p-6 mb-8 border border-indigo-500/20">
					<h3 class="text-xs font-extrabold text-indigo-400 uppercase tracking-widest mb-4">Landlord Hub</h3>
					<div class="space-y-3">
						<input id="building" type="text" placeholder="Apartment Name" class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-2.5 text-sm text-white outline-none">
						<input id="loc" type="text" placeholder="Location" class="w-full bg-slate-900/50 border border-slate-700 rounded-xl px-4 py-2.5 text-sm text-white outline-none">
						<button onclick="uploadHouse()" class="w-full bg-white text-slate-900 font-bold py-3 rounded-xl text-sm transition">Publish Listing</button>
					</div>
				</div>
				<div class="space-y-4">
					<input id="searchLoc" onkeyup="fetchHouses()" type="text" placeholder="Search Neighborhood..." class="w-full bg-slate-900 border border-slate-700 rounded-2xl px-5 py-3.5 text-white outline-none">
				</div>
			</div>
			<div class="p-6 border-t border-white/5 flex items-center justify-between">
				<div class="flex items-center gap-3"><div class="font-bold text-white">%s</div></div>
				<a href="/logout" class="text-xs font-bold text-slate-400">Logout</a>
			</div>
		</aside>

		<main class="flex-1 h-full overflow-y-auto relative z-10">
			<div class="p-4 md:p-8 lg:px-12 max-w-[1600px] mx-auto">
				<header class="mb-10">
					<h2 class="text-3xl md:text-5xl font-extrabold text-white tracking-tight">Explore <span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-300">Sanctuaries</span></h2>
				</header>
				<div id="results-area" class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-8 pb-24"></div>
			</div>
		</main>

		<div id="detail-overlay" class="fixed inset-0 z-[150] hidden">
			<div class="absolute inset-0 bg-black/60 backdrop-blur-sm" onclick="closeDetails()"></div>
			<div class="absolute right-0 top-0 h-full w-full md:w-[500px] lg:w-[600px] detail-pane animate-slide p-8 overflow-y-auto shadow-2xl">
				<button onclick="closeDetails()" class="mb-8 text-slate-400 hover:text-white flex items-center gap-2 font-bold text-sm">
					<span>&larr;</span> Back to Search
				</button>
				<div id="detail-content"></div>
			</div>
		</div>

		<div id="toast" class="hidden fixed top-6 left-1/2 -translate-x-1/2 bg-indigo-600 px-6 py-3 rounded-full text-sm font-bold text-white z-[200]">
			<span id="toast-msg"></span>
		</div>`, myHubButton, landlordPanelDisplay, currentUsername)
}
func GetScripts(isLoggedIn bool, currentUsername string) string {
	return fmt.Sprintf(`
	<script>
		const isLoggedIn = %v;
		const currentUsername = "%s";
		let allHousesData = [];

		document.addEventListener("DOMContentLoaded", () => {
			const params = new URLSearchParams(window.location.search);
			if (params.get('loc')) document.getElementById('searchLoc').value = params.get('loc');
			fetchHouses();
		});

		function fetchHouses() {
			fetch('/houses').then(res => res.json()).then(data => {
				allHousesData = data;
				renderList(data);
			});
		}

		function renderList(data) {
			const container = document.getElementById('results-area');
			const sLoc = document.getElementById('searchLoc').value.toLowerCase();
			container.innerHTML = "";

			data.filter(h => !sLoc || h.location.toLowerCase().includes(sLoc)).forEach(h => {
				const card = document.createElement('div');
				card.className = "glass-card rounded-[2rem] p-5 flex flex-col relative group transition-all duration-300 hover:-translate-y-2 cursor-pointer border border-white/5 hover:border-indigo-500/30";
				card.onclick = () => openDetails(h.id);
				card.innerHTML = '<div class="w-full h-56 bg-slate-800 rounded-3xl overflow-hidden relative mb-5"><img src="'+h.image_urls[0]+'" class="w-full h-full object-cover group-hover:scale-110 transition duration-700"><div class="absolute top-4 right-4 bg-slate-900/90 px-4 py-1.5 rounded-full text-white font-extrabold text-sm">KES '+h.price.toLocaleString()+'</div></div><h3 class="text-2xl font-bold text-white mb-2">'+h.building_name+'</h3><p class="text-slate-400 text-sm mb-4">&nbsp;'+h.location+'</p><button class="w-full py-3 rounded-xl bg-slate-800 text-slate-300 text-xs font-bold uppercase tracking-widest">View Details</button>';
				container.appendChild(card);
			});
		}

		function openDetails(id) {
			const house = allHousesData.find(h => h.id == id);
			if (!house) return;
			const content = document.getElementById('detail-content');
			content.innerHTML = '<div class="rounded-3xl overflow-hidden mb-8"><img src="'+house.image_urls[0]+'" class="w-full h-80 object-cover"></div><div class="flex justify-between items-start mb-6"><div><h1 class="text-3xl font-extrabold text-white">'+house.building_name+'</h1><p class="text-indigo-400 font-bold uppercase tracking-widest text-xs">📍 '+house.location+'</p></div><div class="text-right"><p class="text-2xl font-black text-white">KES '+house.price.toLocaleString()+'</p></div></div><div class="grid grid-cols-2 gap-4 mb-8"><div class="bg-white/5 p-4 rounded-2xl border border-white/5"><p class="text-slate-500 text-[10px] font-bold uppercase mb-1">Type</p><p class="text-white font-bold">'+house.type+'</p></div><div class="bg-white/5 p-4 rounded-2xl border border-white/5"><p class="text-slate-500 text-[10px] font-bold uppercase mb-1">Deposit</p><p class="text-white font-bold">1 Month</p></div></div><p class="text-slate-400 leading-relaxed mb-8">'+house.details+'</p><div class="sticky bottom-0 bg-slate-900/80 backdrop-blur-md pt-4 border-t border-white/10">' + (house.is_booked ? '<button class="w-full bg-emerald-500 text-white font-bold py-4 rounded-2xl">Unlocked</button>' : '<button onclick="payWithMpesa('+house.id+')" class="w-full bg-indigo-600 text-white font-bold py-4 rounded-2xl shadow-lg shadow-indigo-500/20">Pay KES 1,000 to View</button>') + '</div>';
			document.getElementById('detail-overlay').classList.remove('hidden');
		}

		function closeDetails() { document.getElementById('detail-overlay').classList.add('hidden'); }
		function showToast(msg) { const t = document.getElementById("toast"); document.getElementById("toast-msg").innerText = msg; t.classList.remove("hidden"); setTimeout(() => t.classList.add("hidden"), 3000); }
		function payWithMpesa(id) {
			let phone = prompt("Enter M-Pesa Number:");
			if (!phone) return;
			showToast("Requesting M-Pesa...");
			fetch('/pay?id=' + id + '&phone=' + phone, {method: 'POST'})
			.then(res => res.json())
			.then(data => { 
				if(data.ResponseCode === "0") { showToast("Check phone for STK Push!"); } 
				else { showToast(data.CustomerMessage || "Failed"); } 
			});
		}
	</script>`, isLoggedIn, currentUsername)
}