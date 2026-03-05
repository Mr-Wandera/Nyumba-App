package templates

import "fmt"

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

func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
    // This provides the structural UI for your Explore page
	return fmt.Sprintf(`<!DOCTYPE html><html><head><title>Explore | Nyumba</title>
	<script src="https://cdn.tailwindcss.com"></script></head>
	<body class="bg-[#0a0a0a] text-white p-8">
		<header class="flex justify-between items-center mb-10">
			<h2 class="text-4xl font-extrabold">Explore <span class="text-indigo-400">Sanctuaries</span></h2>
			<div class="text-indigo-400 font-bold">` + currentUsername + `</div>
		</header>
		<div id="results-area" class="grid grid-cols-1 md:grid-cols-3 gap-8"></div>
	</body></html>`)
}

func GetScripts(isLoggedIn bool, currentUsername string) string {
	return `<script>
		function fetchHouses() {
			fetch('/houses').then(res => res.json()).then(data => {
				const container = document.getElementById('results-area');
				container.innerHTML = "";
				data.forEach(h => {
					const div = document.createElement('div');
					div.className = "bg-slate-900 border border-white/5 p-6 rounded-[2rem] shadow-xl";
					div.innerHTML = '<img src="'+h.image_urls[0]+'" class="rounded-2xl mb-4 h-48 w-full object-cover"><h3 class="text-xl font-bold">'+h.building_name+'</h3><p class="text-indigo-400 font-bold uppercase text-xs">📍 '+h.location+'</p><button class="w-full bg-indigo-600 py-4 mt-6 rounded-xl font-bold">Pay KES 1,000 to View</button>';
					container.appendChild(div);
				});
			});
		}
		window.onload = fetchHouses;
	</script>`
}