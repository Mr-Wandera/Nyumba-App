package templates

import "fmt"

// Define the Head section separately for consistency
func GetHead() string {
	return `
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Nyumba | Sanctuary Simplified</title>
		<script src="https://cdn.tailwindcss.com"></script>
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet">
		<style>
			:root { --accent: #6366f1; }
			body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: #f8fafc; }
			
			/* Typography Scale */
			h1 { font-size: clamp(3rem, 8vw, 6rem); font-weight: 800; line-height: 0.9; letter-spacing: -0.05em; }
			h2 { font-size: 2.5rem; font-weight: 700; letter-spacing: -0.02em; }
			h3 { font-size: 1.5rem; font-weight: 600; }
			.body-text { font-size: 1.125rem; line-height: 1.6; color: #94a3b8; }
			.caption { font-size: 0.75rem; font-weight: 800; text-transform: uppercase; letter-spacing: 0.2em; color: #64748b; }

			/* Motion & Micro-interactions */
			.hover-lift { transition: transform 0.3s cubic-bezier(0.34, 1.56, 0.64, 1); }
			.hover-lift:hover { transform: translateY(-8px); }
			@keyframes fadeIn { from { opacity: 0; transform: translateY(20px); } to { opacity: 1; transform: translateY(0); } }
			.reveal { animation: fadeIn 0.8s ease-out forwards; }
		</style>
	</head>`
}

func GetLandingHTML() string {
	return `<!DOCTYPE html>` + GetHead() + `
	<body class="min-h-screen flex flex-col overflow-x-hidden">
		<main class="h-screen w-full flex flex-col items-center justify-center px-6 relative">
			<div class="absolute inset-0 bg-[radial-gradient(circle_at_50%_50%,#1e1b4b_0%,#0a0a0a_70%)] opacity-50"></div>
			
			<div class="reveal text-center z-10">
				<p class="caption mb-6 text-indigo-400">Verified Listings Only</p>
				<h1 class="mb-8">Find Your <span class="text-white">Sanctuary.</span><br>
				<span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-400">Simplified.</span></h1>
				
				<p class="body-text max-w-2xl mx-auto mb-12">
					An exclusive platform connecting serious renters with verified landlords. 
					No agents. No endless scrolling. Just your next home.
				</p>

				<a href="/explore" class="bg-white text-black px-12 py-5 rounded-full font-black text-xl hover:bg-indigo-500 hover:text-white transition-all shadow-2xl shadow-white/5">
					Start Your Search →
				</a>
			</div>
		</main>` + GetNeighborhoodScroller() + `
	</body></html>`
}
			

// GetNeighborhoodScroller provides the infinite horizontal loop
func GetNeighborhoodScroller() string {
	return `
	<div class="w-full overflow-hidden py-12 bg-black/40 backdrop-blur-sm border-y border-white/5 relative">
		<div class="absolute inset-y-0 left-0 w-32 bg-gradient-to-r from-[#0a0a0a] to-transparent z-10"></div>
		<div class="absolute inset-y-0 right-0 w-32 bg-gradient-to-l from-[#0a0a0a] to-transparent z-10"></div>
		
		<div class="flex whitespace-nowrap animate-scroll gap-12 items-center">
			<span class="text-4xl font-black text-slate-800 uppercase tracking-tighter">Thika Town</span>
			<span class="text-4xl font-black text-slate-800 uppercase tracking-tighter underline decoration-indigo-500/30">Section 9</span>
			<span class="text-4xl font-black text-slate-800 uppercase tracking-tighter italic">Ngoingwa</span>
			<span class="text-4xl font-black text-slate-800 uppercase tracking-tighter">Landless</span>
			<span class="text-4xl font-black text-slate-800 uppercase tracking-tighter">Thika Town</span>
			<span class="text-4xl font-black text-slate-800 uppercase tracking-tighter underline decoration-indigo-500/30">Section 9</span>
		</div>
	</div>
	<style>
		@keyframes scroll { 0% { transform: translateX(0); } 100% { transform: translateX(-50%); } }
		.animate-scroll { animation: scroll 25s linear infinite; width: max-content; }
	</style>`
}

// GetHTML provides the dashboard layout with the property submission sidebar
func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Nyumba | Explore</title>
		<script src="https://cdn.tailwindcss.com"></script>
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;600;900&display=swap" rel="stylesheet">
		<style>
			body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: white; overflow: hidden; }
			.glass-sidebar { background: rgba(15, 23, 42, 0.8); border-right: 1px solid rgba(255, 255, 255, 0.05); }
			.glass-card { background: rgba(30, 41, 59, 0.4); border: 1px solid rgba(255, 255, 255, 0.05); backdrop-filter: blur(16px); }
			.input-field { background: #0f172a; border: 1px solid #1e293b; color: white; }
		</style>
	</head>
	<body class="h-screen flex">
		<aside class="w-[350px] glass-sidebar p-6 flex flex-col overflow-y-auto">
			<div class="mb-10">
				<h1 class="text-3xl font-black tracking-tighter">Nyumba<span class="text-indigo-500">.</span></h1>
				<p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest">Curated Living</p>
			</div>
			<div class="space-y-4 mb-10">
				<input type="text" placeholder="Apartment Name" class="input-field w-full p-3 rounded-xl text-sm">
				<input type="text" placeholder="Thika" class="input-field w-full p-3 rounded-xl text-sm">
				<select class="input-field w-full p-3 rounded-xl text-sm"><option>Bedsitter</option><option>One Bedroom</option></select>
				<button class="w-full bg-white text-black py-4 rounded-xl font-black">Publish Listing</button>
			</div>
			<div class="mt-auto flex items-center justify-between p-4 bg-indigo-600/10 rounded-2xl">
				<div class="flex items-center gap-3">
					<div class="w-8 h-8 bg-indigo-500 rounded-full flex items-center justify-center font-bold text-xs">👤</div>
					<span class="font-bold text-sm">%%s</span>
				</div>
				<a href="/logout" class="text-xs font-bold text-slate-500">Logout</a>
			</div>
		</aside>
		<main class="flex-1 p-10 overflow-y-auto">
			<h1 class="text-5xl font-black tracking-tighter mb-8 text-indigo-400">Explore Sanctuaries</h1>
			<div id="results-area" class="grid grid-cols-1 lg:grid-cols-2 gap-8"></div>
			` + GetTrustSignals() + `
		</main>
	</body></html>`, currentUsername)
}

// GetSignupHTML restores the premium signup page
func GetSignupHTML() string {
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Join Nyumba | Curated Living</title>
		<script src="https://cdn.tailwindcss.com"></script>
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@400;800&display=swap" rel="stylesheet">
		<style>
			body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: white; }
			.glass-card { background: rgba(30, 41, 59, 0.4); border: 1px solid rgba(255, 255, 255, 0.05); backdrop-filter: blur(16px); }
			.input-field { background: #0f172a; border: 1px solid #1e293b; color: white; }
		</style>
	</head>
	<body class="min-h-screen flex items-center justify-center p-6">
		<div class="glass-card w-full max-w-md rounded-[2.5rem] p-10 shadow-2xl">
			<div class="text-center mb-8">
				<h1 class="text-4xl font-black tracking-tighter mb-2 text-white">Create Account</h1>
				<p class="text-slate-400 text-sm font-semibold uppercase tracking-widest">Join 500+ curated renters today</p>
			</div>
			<form action="/signup" method="POST" class="space-y-5">
				<input type="text" name="username" placeholder="Username" class="input-field w-full py-4 px-6 rounded-2xl font-semibold outline-none focus:border-indigo-500">
				<input type="text" name="phone" placeholder="M-Pesa Number" class="input-field w-full py-4 px-6 rounded-2xl font-semibold outline-none focus:border-indigo-500">
				<input type="password" name="password" placeholder="Password" class="input-field w-full py-4 px-6 rounded-2xl font-semibold outline-none focus:border-indigo-500">
				<button type="submit" class="w-full bg-indigo-600 hover:bg-indigo-500 text-white font-bold py-5 rounded-2xl transition-all shadow-xl shadow-indigo-600/20">Start Journey</button>
			</form>
		</div>
	</body></html>`
}

func GetTrustSignals() string {
	return `<section class="mt-16 border-t border-white/5 pt-12">
		<h3 class="text-indigo-400 font-bold uppercase tracking-[0.2em] text-xs mb-6 text-center">The Science of the Search</h3>
		<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
			<div class="glass-card p-6 rounded-3xl border border-emerald-500/20"><h4 class="font-bold text-white mb-2">🛡️ Verification</h4><p class="text-xs text-slate-400">Cross-referenced ID checks.</p></div>
			<div class="glass-card p-6 rounded-3xl border border-indigo-500/20"><h4 class="font-bold text-white mb-2">🤝 Accountability</h4><p class="text-xs text-slate-400">Landlord transparency ratings.</p></div>
			<div class="glass-card p-6 rounded-3xl border border-purple-500/20"><h4 class="font-bold text-white mb-2">📈 Passive Income</h4><p class="text-xs text-slate-400">Verified leads for owners.</p></div>
		</div>
	</section>`
}

func GetScripts(isLoggedIn bool, currentUsername string) string {
	return `<script>
		async function fetchHouses() {
			try {
				const response = await fetch('/houses');
				const data = await response.json();
				const container = document.getElementById('results-area');
				container.innerHTML = "";
				data.forEach(h => {
					const div = document.createElement('div');
					div.className = "glass-card p-8 rounded-[2.5rem] relative group";
					div.innerHTML = '<div class="relative h-64 bg-slate-800 rounded-[2rem] overflow-hidden mb-6"><img src="'+h.image_urls[0]+'" class="w-full h-full object-cover group-hover:scale-110 transition duration-700"><div class="absolute top-4 right-4 bg-slate-900/90 px-4 py-2 rounded-2xl text-white font-bold text-sm">KES '+h.price.toLocaleString()+'</div></div><h2 class="text-3xl font-bold mb-4">'+h.building_name+'</h2><button class="w-full bg-indigo-500 hover:bg-indigo-400 text-white py-5 rounded-2xl font-bold">Pay KES 1,000 to View</button>';
					container.appendChild(div);
				});
			} catch (e) { console.error("Data fetch failed", e); }
		}
		window.onload = fetchHouses;
	</script>`
}
func GetScripts(isLoggedIn bool, currentUsername string) string {
	return `<script>
		async function fetchHouses() {
			const res = await fetch('/houses');
			const data = await res.json();
			const container = document.getElementById('results-area');
			container.innerHTML = data.map(h => ` + "`" + `
				<div class="glass-card hover-lift p-6 rounded-[2.5rem] bg-slate-900/40 border border-white/5 group">
					<div class="relative h-64 rounded-[2rem] overflow-hidden mb-6">
						<img src="${h.image_urls[0]}" class="w-full h-full object-cover group-hover:scale-110 transition duration-700">
						<div class="absolute top-4 right-4 bg-black/80 backdrop-blur-md px-4 py-2 rounded-2xl font-bold text-sm">
							KES ${h.price.toLocaleString()}
						</div>
					</div>
					<h3 class="mb-2">${h.building_name}</h3>
					<p class="caption text-indigo-400 mb-6">📍 ${h.location}</p>
					<button class="w-full bg-indigo-600 hover:bg-indigo-500 py-5 rounded-2xl font-black transition shadow-lg shadow-indigo-600/20">
						Pay KES 1,000 to View
					</button>
				</div>
			` + "`" + `).join("");
		}
		window.onload = fetchHouses;
	</script>`
}