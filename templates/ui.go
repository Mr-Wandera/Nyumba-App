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

package templates

// GetLandingHTML restores your high-end professional landing page design
func GetLandingHTML() string {
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Nyumba | Find Your Sanctuary</title>
		<script src="https://cdn.tailwindcss.com"></script>
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;900&display=swap" rel="stylesheet">
		<style>
			body { 
				font-family: 'Outfit', sans-serif; 
				background-color: #0a0a0a;
				background-image: radial-gradient(circle at 80% 20%, #1e1b4b 0%, #0a0a0a 50%);
				color: #f8fafc;
				overflow-x: hidden;
			}
			.glass-nav {
				background: rgba(15, 23, 42, 0.6);
				backdrop-filter: blur(12px);
				border: 1px solid rgba(255, 255, 255, 0.05);
			}
			.badge-glow {
				box-shadow: 0 0 15px rgba(16, 185, 129, 0.2);
			}
		</style>
	</head>
	<body class="min-h-screen flex flex-col">
		<nav class="fixed top-6 left-1/2 -translate-x-1/2 w-[90%] max-w-5xl glass-nav rounded-full px-8 py-4 flex justify-between items-center z-50">
			<div class="text-2xl font-black tracking-tighter">Nyumba<span class="text-indigo-500">.</span></div>
			<div class="hidden md:flex items-center gap-8 text-sm font-semibold text-slate-400">
				<a href="#" class="hover:text-white transition">How it Works</a>
				<a href="#" class="hover:text-white transition">Neighborhoods</a>
				<a href="#" class="hover:text-white transition">For Landlords</a>
			</div>
			<div class="flex items-center gap-4">
				<a href="/login" class="text-sm font-bold hover:text-indigo-400 transition">Sign In</a>
				<a href="/explore" class="bg-indigo-600 hover:bg-indigo-500 text-white px-6 py-2.5 rounded-full font-bold text-sm shadow-lg shadow-indigo-600/20 transition transform hover:scale-105">Explore</a>
			</div>
		</nav>

		<main class="flex-1 flex flex-col items-center justify-center pt-32 pb-20 px-6 relative">
			<div class="badge-glow inline-flex items-center gap-2 bg-emerald-500/10 border border-emerald-500/20 px-4 py-1.5 rounded-full mb-8">
				<div class="w-2 h-2 bg-emerald-500 rounded-full animate-pulse"></div>
				<span class="text-[10px] font-black text-emerald-500 uppercase tracking-[0.2em]">Verified Listings Only</span>
			</div>

			<h1 class="text-6xl md:text-8xl font-black text-center tracking-tighter leading-[0.9] mb-8">
				Find Your <span class="text-white">Sanctuary.</span><br>
				<span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 via-purple-400 to-cyan-400">Simplified.</span>
			</h1>

			<p class="text-slate-400 text-center text-lg md:text-xl max-w-2xl mb-12 leading-relaxed">
				An exclusive platform connecting serious renters with verified landlords. No agents. No endless scrolling. Just your next home.
			</p>

			<a href="/explore" class="bg-white text-black hover:bg-slate-200 px-10 py-5 rounded-full font-black text-lg transition-all transform hover:scale-105 shadow-xl shadow-white/5">
				Start Your Search →
			</a>
		</main>
	</body>
	</html>`
}
