package templates

import "fmt"

// GetLandingHTML restores the high-fidelity hero page
func GetLandingHTML() string {
	return `<!DOCTYPE html><html><head>
		<title>Nyumba | Find Your Sanctuary</title>
		<script src="https://cdn.tailwindcss.com"></script>
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;900&display=swap" rel="stylesheet">
		<style>
			body { font-family: 'Outfit', sans-serif; background-color: #0a0a0a; color: #f8fafc; overflow-x: hidden; }
			.glass-nav { background: rgba(15, 23, 42, 0.6); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.05); }
		</style>
	</head>
	<body class="min-h-screen flex flex-col">
		<nav class="fixed top-6 left-1/2 -translate-x-1/2 w-[90%] max-w-5xl glass-nav rounded-full px-8 py-4 flex justify-between items-center z-50">
			<div class="text-2xl font-black tracking-tighter">Nyumba<span class="text-indigo-500">.</span></div>
			<div class="flex gap-4">
				<a href="/login" class="text-sm font-bold hover:text-indigo-400 transition">Sign In</a>
				<a href="/explore" class="bg-indigo-600 text-white px-6 py-2.5 rounded-full font-bold text-sm">Explore</a>
			</div>
		</nav>
		<main class="flex-1 flex flex-col items-center justify-center pt-32 pb-20 px-6">
			<h1 class="text-6xl md:text-8xl font-black text-center tracking-tighter leading-[0.9] mb-8">
				Find Your <span class="text-white">Sanctuary.</span><br>
				<span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-400">Simplified.</span>
			</h1>
			<a href="/signup" class="bg-white text-black px-10 py-5 rounded-full font-black text-lg">Start Your Search →</a>
		</main>` + GetNeighborhoodScroller() + `</body></html>`
}

func GetNeighborhoodScroller() string {
	return `<div class="w-full overflow-hidden py-12 bg-black/40 border-y border-white/5 whitespace-nowrap">
		<div class="flex animate-scroll gap-12 items-center">
			<span class="text-4xl font-black text-slate-800 uppercase italic">Section 9</span>
			<span class="text-4xl font-black text-slate-800 uppercase">Thika Town</span>
			<span class="text-4xl font-black text-slate-800 uppercase underline decoration-indigo-500/30">Ngoingwa</span>
			<span class="text-4xl font-black text-slate-800 uppercase italic">Section 9</span>
		</div>
	</div>
	<style>
		@keyframes scroll { 0% { transform: translateX(0); } 100% { transform: translateX(-50%); } }
		.animate-scroll { animation: scroll 20s linear infinite; width: max-content; }
	</style>`
}

// GetHTML restores the premium Dashboard with Sidebar
func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	return fmt.Sprintf(`<!DOCTYPE html><html><head>
		<script src="https://cdn.tailwindcss.com"></script>
		<style>body { background: #0a0a0a; color: white; overflow: hidden; }
		.glass-sidebar { background: rgba(15, 23, 42, 0.8); border-right: 1px solid rgba(255, 255, 255, 0.05); }</style>
	</head>
	<body class="h-screen flex">
		<aside class="w-[350px] glass-sidebar p-6 flex flex-col">
			<h1 class="text-3xl font-black mb-10 tracking-tighter">Nyumba<span class="text-indigo-500">.</span></h1>
			<p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-10">Curated Living</p>
			<button class="w-full bg-white text-black py-4 rounded-xl font-black">Publish Listing</button>
			<div class="mt-auto p-4 bg-indigo-600/10 rounded-2xl flex justify-between items-center">
				<div class="flex items-center gap-3"><span class="font-bold">%%s</span></div>
				<a href="/" class="text-xs">Logout</a>
			</div>
		</aside>
		<main class="flex-1 p-10 overflow-y-auto">
			<h1 class="text-5xl font-black mb-8 text-indigo-400">Explore Sanctuaries</h1>
			<div id="results-area" class="grid grid-cols-1 lg:grid-cols-2 gap-8"></div>
		</main></body></html>`, currentUsername)
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