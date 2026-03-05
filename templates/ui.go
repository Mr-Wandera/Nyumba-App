package templates

import "fmt"


func GetLandingHTML() string {
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Nyumba | Find Your Sanctuary</title>
		<script src="https://cdn.tailwindcss.com"></script>
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;600;900&display=swap" rel="stylesheet">
		<style>
			body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: white; overflow-x: hidden; }
			
			/* 1. The "Fade" Effect (Essential for UI) */
			.scroll-container {
				mask-image: linear-gradient(to right, transparent, black 15%, black 85%, transparent);
				-webkit-mask-image: linear-gradient(to right, transparent, black 15%, black 85%, transparent);
			}

			/* 3. Smoothness & Performance */
			@keyframes scroll {
				0% { transform: translateX(0); }
				100% { transform: translateX(-50%); }
			}

			.scrolling-text {
				display: flex;
				white-space: nowrap;
				/* Slowed down for a premium feel */
				animation: scroll 40s linear infinite; 
			}

			/* 4. Pause on Hover Interaction */
			.scroll-container:hover .scrolling-text {
				animation-play-state: paused;
			}

			/* 2. Typography & Spacing */
			.ticker-item {
				letter-spacing: 0.15em;
				color: #a1a1aa; /* Muted color to keep focus on CTA */
			}
		</style>
	</head>
	<body class="min-h-screen flex flex-col">
		<section class="scroll-container w-full py-16 bg-black/40 border-y border-white/5 relative overflow-hidden">
			<div class="scrolling-text gap-24 items-center">
				` + getTickerContent() + getTickerContent() + `
			</div>
		</section>
	</body>
	</html>`
}

func getTickerContent() string {
	// 5. Content Variation
	items := []string{
		"Verified Listings Only",
		"Zero Agent Fees Guaranteed",
		"50+ New Sanctuaries in Thika",
		"Direct Contact with Vetted Owners",
	}
	
	html := ""
	for _, item := range items {
		html += fmt.Sprintf(`
			<span class="ticker-item text-2xl md:text-3xl font-black uppercase whitespace-nowrap">
				%s <span class="text-indigo-500 mx-8">•</span>
			</span>`, item)
	}
	return html
}


// GetHTML restores the premium Dashboard with the Sidebar
func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Nyumba | Explore Sanctuaries</title>
		<script src="https://cdn.tailwindcss.com"></script>
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;900&display=swap" rel="stylesheet">
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
				<input type="text" placeholder="📍 Google Maps Link" class="input-field w-full p-3 rounded-xl text-sm">
				<select class="input-field w-full p-3 rounded-xl text-sm">
					<option>Bedsitter</option>
					<option>One Bedroom</option>
				</select>
				<div class="flex gap-2">
					<input type="number" placeholder="7500" class="input-field w-1/2 p-3 rounded-xl text-sm">
					<input type="number" placeholder="2000" class="input-field w-1/2 p-3 rounded-xl text-sm">
				</div>
				<button class="w-full bg-indigo-600/20 text-indigo-400 border border-indigo-500/20 py-3 rounded-xl font-bold text-xs uppercase tracking-widest">Choose files</button>
				<textarea placeholder="Beautiful apartment with..." class="input-field w-full p-3 rounded-xl text-sm h-24"></textarea>
				<button class="w-full bg-white text-black py-4 rounded-xl font-black shadow-xl">Publish Listing</button>
			</div>

			<div class="mt-auto flex items-center justify-between p-4 bg-indigo-600/10 rounded-2xl">
				<div class="flex items-center gap-3">
					<div class="w-8 h-8 bg-indigo-500 rounded-full flex items-center justify-center font-bold text-xs">👤</div>
					<span class="font-bold text-sm">%s</span>
				</div>
				<a href="/logout" class="text-xs font-bold text-slate-500">Logout</a>
			</div>
		</aside>

		<main class="flex-1 p-10 overflow-y-auto relative">
			<div class="absolute top-0 right-0 w-[50vw] h-[50vw] bg-indigo-600/5 rounded-full blur-[120px] pointer-events-none"></div>
			
			<header class="mb-12">
				<h1 class="text-5xl font-black tracking-tighter mb-2">Explore <span class="text-indigo-400">Sanctuaries</span></h1>
				<p class="text-slate-400">Find your next home and connect directly with verified owners.</p>
			</header>

			<div id="results-area" class="grid grid-cols-1 lg:grid-cols-2 gap-8">
				</div>
		</main>
	</body>
	</html>`, currentUsername)
}

func GetSignupHTML() string {
	return `<!DOCTYPE html><html><head><script src="https://cdn.tailwindcss.com"></script></head>
	<body class="bg-[#0a0a0a] min-h-screen flex items-center justify-center p-6">
		<div class="bg-slate-900/40 p-10 rounded-[2.5rem] w-full max-w-md border border-white/5">
			<h1 class="text-4xl font-black text-center mb-8 text-white">Create Account</h1>
			<form action="/signup" method="POST" class="space-y-5">
				<input type="text" name="username" placeholder="Username" class="w-full p-4 rounded-2xl bg-slate-950 text-white outline-none border border-slate-800">
				<button type="submit" class="w-full bg-indigo-600 text-white py-5 rounded-2xl font-bold">Start Journey</button>
			</form>
		</div></body></html>`
}

func GetScripts(isLoggedIn bool, currentUsername string) string {
	return `<script>
		async function fetchHouses() {
			const res = await fetch('/houses');
			const data = await res.json();
			const container = document.getElementById('results-area');
			container.innerHTML = data.map(h => '<div class="bg-slate-900/40 p-8 rounded-[2.5rem] border border-white/5"><div class="relative h-64 bg-slate-800 rounded-[2rem] overflow-hidden mb-6"><img src="'+h.image_urls[0]+'" class="w-full h-full object-cover"><div class="absolute top-4 right-4 bg-black/80 px-4 py-2 rounded-xl text-sm font-bold">KES '+h.price.toLocaleString()+'</div></div><h2 class="text-3xl font-bold">'+h.building_name+'</h2><button class="w-full bg-indigo-500 py-5 mt-6 rounded-2xl font-bold">Pay KES 1,000 to View</button></div>').join("");
		}
		window.onload = fetchHouses;
	</script>`
}