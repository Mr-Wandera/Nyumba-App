package templates

import "fmt"

func GetLandingHTML() string {
	return `<!DOCTYPE html><html><head>
		<title>Nyumba | Find Your Sanctuary</title>
		<script src="https://cdn.tailwindcss.com"></script>
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;900&display=swap" rel="stylesheet">
		<style>
			body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: white; overflow-x: hidden; }
			.scroll-container { mask-image: linear-gradient(to right, transparent, black 15%, black 85%, transparent); -webkit-mask-image: linear-gradient(to right, transparent, black 15%, black 85%, transparent); }
			@keyframes scroll { 0% { transform: translateX(0); } 100% { transform: translateX(-50%); } }
			.scrolling-text { display: flex; white-space: nowrap; animation: scroll 40s linear infinite; }
			.scrolling-text:hover { animation-play-state: paused; }
		</style>
	</head>
	<body class="min-h-screen flex flex-col">
		<main class="flex-1 flex flex-col items-center justify-center pt-32 pb-12">
			<h1 class="text-6xl md:text-8xl font-black text-center mb-8 tracking-tighter leading-[0.9]">
				Find Your <span class="text-white">Sanctuary.</span><br>
				<span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-400">Simplified.</span>
			</h1>
			<a href="/explore" class="bg-white text-black px-10 py-5 rounded-full font-black text-lg">Start Your Search →</a>
		</main>
		<section class="scroll-container w-full py-16 bg-black/40 border-y border-white/5 overflow-hidden">
			<div class="scrolling-text gap-24 items-center">
				` + getTickerItems() + getTickerItems() + `
			</div>
		</section>
	</body></html>`
}

func getTickerItems() string {
	locs := []string{"Thika Town", "Section 9", "Ngoingwa", "Landless"}
	html := ""
	for _, l := range locs {
		html += fmt.Sprintf(`<a href="/explore?location=%s" class="text-2xl font-black uppercase tracking-widest text-slate-500 hover:text-indigo-400 transition">%s <span class="mx-8">•</span></a>`, l, l)
	}
	return html
}

func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	return fmt.Sprintf(`<!DOCTYPE html><html><head><script src="https://cdn.tailwindcss.com"></script></head>
	<body class="h-screen flex bg-[#0a0a0a] text-white overflow-hidden">
		<aside class="w-[350px] border-r border-white/5 p-6 flex flex-col overflow-y-auto">
			<h1 class="text-3xl font-black mb-10">Nyumba<span class="text-indigo-500">.</span></h1>
			
			<form action="/add-house" method="POST" class="space-y-4 mb-10">
				<input type="text" name="building_name" placeholder="Apartment Name" class="w-full p-3 rounded-xl bg-slate-900 border border-white/5 text-sm">
				<input type="text" name="location" placeholder="Thika (e.g. Section 9)" class="w-full p-3 rounded-xl bg-slate-900 border border-white/5 text-sm">
				<button type="submit" class="w-full bg-white text-black py-4 rounded-xl font-black">Publish Listing</button>
			</form>

			<div class="mt-auto flex items-center justify-between p-4 bg-indigo-600/10 rounded-2xl">
				<span class="font-bold text-sm">%%s</span> <a href="/" class="text-xs">Logout</a>
			</div>
		</aside>
		<main class="flex-1 p-10 overflow-y-auto">
			<h1 class="text-5xl font-black mb-8">Explore <span class="text-indigo-400">Sanctuaries</span></h1>
			<div id="results-area" class="grid grid-cols-1 lg:grid-cols-2 gap-8"></div>
		</main></body></html>`, currentUsername)
}

func GetSignupHTML() string {
	return `<!DOCTYPE html><html><head><script src="https://cdn.tailwindcss.com"></script></head>
	<body class="bg-[#0a0a0a] flex items-center justify-center min-h-screen">
		<form action="/signup" method="POST" class="bg-slate-900/40 p-10 rounded-[2.5rem] border border-white/5 w-full max-w-md">
			<h1 class="text-4xl font-black text-center text-white mb-8">Create Account</h1>
			<button type="submit" class="w-full bg-indigo-600 py-5 rounded-2xl font-bold text-white">Start Journey</button>
		</form>
	</body></html>`
}

func GetScripts(isLoggedIn bool, currentUsername string) string {
	return `<script>
		async function fetchHouses() {
			const params = new URLSearchParams(window.location.search);
			const loc = params.get('location') || '';
			const res = await fetch('/houses?location=' + loc);
			const data = await res.json();
			const container = document.getElementById('results-area');
			container.innerHTML = data.map(h => '<div class="bg-slate-900/40 p-8 rounded-[2.5rem] border border-white/5"><img src="'+h.image_urls[0]+'" class="rounded-[2rem] h-64 w-full object-cover mb-6"><h2 class="text-3xl font-bold">'+h.building_name+'</h2><button class="w-full bg-indigo-500 py-5 mt-6 rounded-2xl font-bold">Pay KES 1,000 to View</button></div>').join("");
		}
		window.onload = fetchHouses;
	</script>`
}