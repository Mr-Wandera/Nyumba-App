package templates

import "fmt"

func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	return fmt.Sprintf(`<!DOCTYPE html><html><head><title>Nyumba | Explore Sanctuaries</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet">
	<script src="https://cdn.tailwindcss.com"></script>
	<style>
		body { font-family: 'Outfit', sans-serif; background: radial-gradient(circle at top right, #1e293b, #0a0a0a); color: #f8fafc; min-height: 100vh; }
		.glass-nav { background: rgba(15, 23, 42, 0.8); backdrop-filter: blur(20px); border-bottom: 1px solid rgba(255,255,255,0.05); }
		.glass-card { background: rgba(30, 41, 59, 0.4); border: 1px solid rgba(255, 255, 255, 0.05); backdrop-filter: blur(16px); }
		.detail-pane { background: rgba(15, 23, 42, 0.98); backdrop-filter: blur(30px); border-left: 1px solid rgba(255,255,255,0.1); }
		@keyframes slideIn { from { transform: translateX(100%%); } to { transform: translateX(0); } }
		.animate-slide { animation: slideIn 0.4s cubic-bezier(0.16, 1, 0.3, 1); }
	</style></head>
	<body class="overflow-x-hidden">
		<nav class="glass-nav sticky top-0 z-[100] px-6 py-4 flex justify-between items-center">
			<div class="text-2xl font-black tracking-tighter text-white">Nyumba<span class="text-indigo-500">.</span></div>
			<div class="flex items-center gap-6">
				<span class="text-slate-400 text-sm">Welcome, <span class="text-indigo-400 font-bold">%s</span></span>
				<a href="/logout" class="text-xs font-bold text-slate-500 hover:text-white transition">Logout</a>
			</div>
		</nav>

		<main class="p-6 md:p-12 max-w-7xl mx-auto">
			<header class="mb-12">
				<div class="inline-flex items-center gap-2 bg-emerald-500/10 border border-emerald-500/20 px-3 py-1 rounded-full mb-4">
					<div class="w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse"></div>
					<span class="text-[10px] font-bold text-emerald-500 uppercase tracking-widest">Verified Listings Only</span>
				</div>
				<h1 class="text-4xl md:text-6xl font-black text-white mb-4 tracking-tighter">Explore <span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-400">Sanctuaries</span></h1>
				<p class="text-slate-400 max-w-xl">Find your next home and connect directly with verified owners without the hassle of agents.</p>
			</header>

			<div id="results-area" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 pb-24"></div>
		</main>

		<div id="detail-overlay" class="fixed inset-0 z-[150] hidden">
			<div class="absolute inset-0 bg-black/60 backdrop-blur-sm" onclick="closeDetails()"></div>
			<div class="absolute right-0 top-0 h-full w-full md:w-[500px] lg:w-[600px] detail-pane animate-slide p-8 overflow-y-auto shadow-2xl">
				<button onclick="closeDetails()" class="mb-8 text-slate-400 hover:text-white flex items-center gap-2 font-bold text-sm">
					<span>←</span> Back to Search
				</button>
				<div id="detail-content"></div>
			</div>
		</div>
	</body></html>`, currentUsername)
}

func GetScripts(isLoggedIn bool, currentUsername string) string {
	return `
	<script>
		let allHousesData = [];
		const fetchHouses = () => {
			fetch('/houses').then(res => res.json()).then(data => {
				allHousesData = data;
				const container = document.getElementById('results-area');
				container.innerHTML = "";
				data.forEach(h => {
					const div = document.createElement('div');
					div.className = "glass-card rounded-[2.5rem] p-6 cursor-pointer border border-white/5 hover:border-indigo-500/30 transition-all duration-500 hover:-translate-y-2 group";
					div.onclick = () => openDetails(h.id);
					div.innerHTML = '<div class="h-64 bg-slate-800 rounded-[2rem] overflow-hidden mb-6 relative"><img src="'+h.image_urls[0]+'" class="w-full h-full object-cover group-hover:scale-110 transition duration-700"><div class="absolute top-4 right-4 bg-slate-900/90 px-4 py-2 rounded-2xl text-white font-bold text-sm">KES '+h.price.toLocaleString()+'</div></div><h3 class="text-2xl font-bold text-white mb-2">'+h.building_name+'</h3><div class="flex items-center gap-2 text-indigo-400 text-xs font-bold uppercase tracking-widest"><span>📍</span> '+h.location+'</div><button class="w-full mt-6 py-4 rounded-2xl bg-indigo-600/10 text-indigo-400 text-xs font-bold uppercase tracking-widest border border-indigo-500/20 group-hover:bg-indigo-600 group-hover:text-white transition-all duration-300">View Details</button>';
					container.appendChild(div);
				});
			});
		};

		const openDetails = (id) => {
			const h = allHousesData.find(x => x.id == id);
			if(!h) return;
			document.getElementById('detail-content').innerHTML = '<div class="rounded-[2rem] overflow-hidden mb-8 shadow-2xl"><img src="'+h.image_urls[0]+'" class="w-full h-80 object-cover"></div><div class="flex justify-between items-start mb-6"><div><h1 class="text-3xl font-black text-white mb-2">'+h.building_name+'</h1><p class="text-indigo-400 font-bold uppercase tracking-widest text-xs">📍 '+h.location+'</p></div><div class="text-right"><p class="text-2xl font-black text-white">KES '+h.price.toLocaleString()+'</p><p class="text-[10px] text-slate-500 uppercase font-bold tracking-tighter">Per Month</p></div></div><p class="text-slate-400 leading-relaxed mb-12 text-lg">'+h.details+'</p><div class="sticky bottom-0 pt-4 bg-slate-950/80 backdrop-blur-md border-t border-white/10"><button class="w-full bg-indigo-600 hover:bg-indigo-500 text-white font-bold py-5 rounded-2xl transition shadow-xl shadow-indigo-500/20 text-lg">Pay KES 1,000 to View</button></div>';
			document.getElementById('detail-overlay').classList.remove('hidden');
		};

		const closeDetails = () => document.getElementById('detail-overlay').classList.add('hidden');
		window.onload = fetchHouses;
	</script>`
}

func GetLandingHTML() string {
	return `<!DOCTYPE html><html><head><title>Nyumba | Find Your Sanctuary</title><script src="https://cdn.tailwindcss.com"></script><link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;600;900&display=swap" rel="stylesheet"></head>
	<body style="font-family:'Outfit',sans-serif;" class="bg-[#0a0a0a] h-screen flex items-center justify-center text-white overflow-hidden relative">
		<div class="absolute top-[-20%%] left-[-10%%] w-[60vw] h-[60vw] bg-indigo-600/10 rounded-full blur-[120px]"></div>
		<div class="text-center relative z-10 px-6">
			<h1 class="text-6xl md:text-8xl font-black mb-8 tracking-tighter leading-none">Find Your <span class="text-transparent bg-clip-text bg-gradient-to-r from-white to-slate-500">Sanctuary.</span><br><span class="text-indigo-500">Simplified.</span></h1>
			<p class="text-slate-400 mb-12 text-lg md:text-xl max-w-2xl mx-auto leading-relaxed">An exclusive platform connecting serious renters with verified landlords. No agents. No endless scrolling. Just your next home.</p>
			<a href="/explore" class="bg-indigo-600 hover:bg-indigo-500 px-12 py-5 rounded-2xl font-bold transition-all transform hover:scale-105 inline-block shadow-2xl shadow-indigo-600/20 text-lg">Start Your Search →</a>
		</div>
	</body></html>`
}