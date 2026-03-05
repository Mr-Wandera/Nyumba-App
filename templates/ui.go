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
		</main>` + GetNeighborhoodScroller() + `
	</body>
	</html>`
}
func GetNeighborhoodScroller() string {
	return `
	<div class="w-full overflow-hidden py-12 bg-black/40 backdrop-blur-sm border-y border-white/5 relative">
		<div class="absolute inset-y-0 left-0 w-32 bg-gradient-to-r from-[#0a0a0a] to-transparent z-10"></div>
		<div class="absolute inset-y-0 right-0 w-32 bg-gradient-to-l from-[#0a0a0a] to-transparent z-10"></div>
		
		<div class="flex whitespace-nowrap animate-scroll gap-8 items-center">
			<span class="text-3xl md:text-5xl font-black text-slate-800 hover:text-indigo-500 transition cursor-default tracking-tighter uppercase">Thika Town</span>
			<span class="text-3xl md:text-5xl font-black text-slate-800 hover:text-indigo-500 transition cursor-default tracking-tighter uppercase underline decoration-indigo-500/30">Section 9</span>
			<span class="text-3xl md:text-5xl font-black text-slate-800 hover:text-indigo-500 transition cursor-default tracking-tighter uppercase">Ngoingwa</span>
			<span class="text-3xl md:text-5xl font-black text-slate-800 hover:text-indigo-500 transition cursor-default tracking-tighter uppercase italic">Landless</span>
			<span class="text-3xl md:text-5xl font-black text-slate-800 hover:text-indigo-500 transition cursor-default tracking-tighter uppercase">Juja</span>
			<span class="text-3xl md:text-5xl font-black text-slate-800 hover:text-indigo-500 transition cursor-default tracking-tighter uppercase">Thika Town</span>
			<span class="text-3xl md:text-5xl font-black text-slate-800 hover:text-indigo-500 transition cursor-default tracking-tighter uppercase underline decoration-indigo-500/30">Section 9</span>
			<span class="text-3xl md:text-5xl font-black text-slate-800 hover:text-indigo-500 transition cursor-default tracking-tighter uppercase">Ngoingwa</span>
		</div>
	</div>
	<style>
		@keyframes scroll { 0% { transform: translateX(0); } 100% { transform: translateX(-50%); } }
		.animate-scroll { animation: scroll 25s linear infinite; width: max-content; }
	</style>`
}
func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	return fmt.Sprintf(`<!DOCTYPE html><html><head><title>Nyumba | Explore</title>
	<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@600;900&display=swap" rel="stylesheet">
	<script src="https://cdn.tailwindcss.com"></script>
	<style>
		body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: white; }
		.glass-card { background: rgba(30, 41, 59, 0.4); border: 1px solid rgba(255, 255, 255, 0.05); backdrop-filter: blur(16px); }
	</style></head>
	<body class="p-8">
		<header class="flex justify-between items-center mb-12">
			<h2 class="text-4xl font-black tracking-tighter">Explore <span class="text-indigo-500">Sanctuaries</span></h2>
			<div class="text-indigo-400 font-bold">%s</div>
		</header>
		<div id="results-area" class="grid grid-cols-1 md:grid-cols-3 gap-8"></div>
	</body></html>`, currentUsername)
}

func GetScripts(isLoggedIn bool, currentUsername string) string {
	return `<script>
		function fetchHouses() {
			fetch('/houses').then(res => res.json()).then(data => {
				const container = document.getElementById('results-area');
				container.innerHTML = "";
				data.forEach(h => {
					const div = document.createElement('div');
					div.className = "glass-card p-6 rounded-[2.5rem]";
					div.innerHTML = '<img src="'+h.image_urls[0]+'" class="rounded-[2rem] mb-6 h-56 w-full object-cover shadow-2xl"><h3 class="text-2xl font-bold mb-2">'+h.building_name+'</h3><p class="text-indigo-400 text-xs font-bold uppercase tracking-widest">📍 '+h.location+'</p><button class="w-full bg-indigo-600 py-4 mt-6 rounded-2xl font-bold shadow-xl shadow-indigo-600/20">Pay KES 1,000 to View</button>';
					container.appendChild(div);
				});
			});
		}
		window.onload = fetchHouses;
	</script>`
}
// Add this to your existing GetHTML function or a new component
func GetTrustSignals() string {
	return `
	<section class="mt-16 border-t border-white/5 pt-12">
		<h3 class="text-indigo-400 font-bold uppercase tracking-[0.2em] text-xs mb-6 text-center">The Science of the Search</h3>
		<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
			<div class="glass-card p-6 rounded-3xl border-emerald-500/20">
				<div class="flex items-center gap-3 mb-4">
					<div class="p-2 bg-emerald-500/10 rounded-lg text-emerald-500">🛡️</div>
					<h4 class="font-bold text-white">Full Verification</h4>
				</div>
				<p class="text-xs text-slate-400 leading-relaxed">Every sanctuary is cross-referenced with ID checks and ownership documents before appearing on your feed.</p>
			</div>

			<div class="glass-card p-6 rounded-3xl border-indigo-500/20">
				<div class="flex items-center gap-3 mb-4">
					<div class="p-2 bg-indigo-500/10 rounded-lg text-indigo-400">🤝</div>
					<h4 class="font-bold text-white">Landlord Accountability</h4>
				</div>
				<p class="text-xs text-slate-400 leading-relaxed">Review landlord responsiveness and transparency. We build communities, not just tenant-owner contracts.</p>
			</div>

			<div class="glass-card p-6 rounded-3xl border-purple-500/20">
				<div class="flex items-center gap-3 mb-4">
					<div class="p-2 bg-purple-500/10 rounded-lg text-purple-400">📈</div>
					<h4 class="font-bold text-white">Passive Income</h4>
				</div>
				<p class="text-xs text-slate-400 leading-relaxed">Landlords: Automate your occupancy. High-quality leads for high-quality properties.</p>
			</div>
		</div>
	</section>`
}