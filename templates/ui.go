package templates

import "fmt"

func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	// We use %% to escape the percent sign for Go's Sprintf
	return fmt.Sprintf(`<!DOCTYPE html><html><head><title>Nyumba</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet">
	<script src="https://cdn.tailwindcss.com"></script>
	<style>
		body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: #f8fafc; overflow-x: hidden; }
		.glass-card { background: rgba(30, 41, 59, 0.4); border: 1px solid rgba(255, 255, 255, 0.1); backdrop-filter: blur(12px); }
		.detail-pane { background: rgba(15, 23, 42, 0.98); backdrop-filter: blur(25px); border-left: 1px solid rgba(255,255,255,0.1); }
		@keyframes slideIn { from { transform: translateX(100%%); } to { transform: translateX(0); } }
		.animate-slide { animation: slideIn 0.4s cubic-bezier(0.16, 1, 0.3, 1); }
	</style></head>
	<body class="h-screen flex flex-col md:flex-row overflow-hidden relative">
		<main class="flex-1 h-full overflow-y-auto relative z-10">
			<div class="p-8 max-w-[1600px] mx-auto">
				<header class="mb-10 flex justify-between items-center">
					<h2 class="text-4xl font-extrabold text-white tracking-tighter">Nyumba<span class="text-indigo-500">.</span></h2>
					<div class="flex items-center gap-4">
						<span class="text-slate-400 text-sm">Active Session:</span>
						<span class="font-bold text-indigo-400">%s</span>
					</div>
				</header>
				<div id="results-area" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 pb-24"></div>
			</div>
		</main>

		<div id="detail-overlay" class="fixed inset-0 z-[150] hidden">
			<div class="absolute inset-0 bg-black/60 backdrop-blur-sm" onclick="closeDetails()"></div>
			<div class="absolute right-0 top-0 h-full w-full md:w-[500px] detail-pane animate-slide p-8 overflow-y-auto shadow-2xl">
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
					div.className = "glass-card rounded-[2rem] p-5 cursor-pointer border border-white/5 hover:border-indigo-500/30 transition-all duration-300 hover:-translate-y-2";
					div.onclick = () => openDetails(h.id);
					div.innerHTML = '<div class="h-56 bg-slate-800 rounded-3xl overflow-hidden mb-5"><img src="'+h.image_urls[0]+'" class="w-full h-full object-cover"></div><h3 class="text-2xl font-bold text-white mb-1">'+h.building_name+'</h3><p class="text-indigo-400 font-bold text-xs uppercase tracking-widest">📍 '+h.location+'</p>';
					container.appendChild(div);
				});
			});
		};

		const openDetails = (id) => {
			const h = allHousesData.find(x => x.id == id);
			if(!h) return;
			document.getElementById('detail-content').innerHTML = '<div class="rounded-3xl overflow-hidden mb-8 shadow-2xl"><img src="'+h.image_urls[0]+'" class="w-full h-80 object-cover"></div><h1 class="text-3xl font-extrabold text-white mb-1">'+h.building_name+'</h1><p class="text-indigo-400 font-bold uppercase tracking-widest text-xs mb-6">📍 '+h.location+'</p><p class="text-slate-400 leading-relaxed mb-8">'+h.details+'</p><button class="w-full bg-indigo-600 hover:bg-indigo-500 text-white font-bold py-4 rounded-2xl transition shadow-lg shadow-indigo-500/20">Pay KES 1,000 to View</button>';
			document.getElementById('detail-overlay').classList.remove('hidden');
		};

		const closeDetails = () => document.getElementById('detail-overlay').classList.add('hidden');
		window.onload = fetchHouses;
	</script>`
}