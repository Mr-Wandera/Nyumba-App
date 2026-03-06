package templates

import "fmt"

// GetLandingHTML provides the refined hero and the infinite scroller
func GetLandingHTML() string {
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Nyumba | Find Your Sanctuary</title>
		<script src="https://cdn.tailwindcss.com"></script>
		<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;900&display=swap" rel="stylesheet">
		<style>
			body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: white; overflow-x: hidden; }
			.scroll-container {
				mask-image: linear-gradient(to right, transparent, black 15%, black 85%, transparent);
				-webkit-mask-image: linear-gradient(to right, transparent, black 15%, black 85%, transparent);
			}
			@keyframes scroll { 0% { transform: translateX(0); } 100% { transform: translateX(-50%); } }
			.scrolling-text { display: flex; white-space: nowrap; animation: scroll 40s linear infinite; }
			.scroll-container:hover .scrolling-text { animation-play-state: paused; }
		</style>
	</head>
	<body class="min-h-screen flex flex-col">
		<main class="flex-1 flex flex-col items-center justify-center pt-32 pb-12 px-6">
			<h1 class="text-6xl md:text-8xl font-black text-center tracking-tighter mb-8 leading-[0.9]">
				Find Your <span class="text-white">Sanctuary.</span><br>
				<span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-400">Simplified.</span>
			</h1>
			<a href="/explore" class="bg-white text-black px-10 py-5 rounded-full font-black text-lg">Start Your Search →</a>
		</main>
		<section class="scroll-container w-full py-16 bg-black/40 border-y border-white/5 relative overflow-hidden">
			<div class="scrolling-text gap-24 items-center">
				` + getTickerContent() + getTickerContent() + `
			</div>
		</section>
	</body></html>`
}

func getTickerContent() string {
	items := []string{"Verified Listings", "Zero Agent Fees", "50+ New Sanctuaries", "Direct Contact"}
	html := ""
	for _, item := range items {
		html += fmt.Sprintf(`<span class="text-2xl font-black uppercase tracking-widest text-slate-500">%s <span class="text-indigo-500 mx-8">•</span></span>`, item)
	}
	return html
}

func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	return fmt.Sprintf(`<!DOCTYPE html><html><head><script src="https://cdn.tailwindcss.com"></script></head>
	<body class="h-screen flex bg-[#0a0a0a] text-white overflow-hidden">
		<aside class="w-[350px] border-r border-white/5 p-6 flex flex-col">
			<h1 class="text-3xl font-black mb-10">Nyumba<span class="text-indigo-500">.</span></h1>
			<div class="mt-auto flex items-center justify-between p-4 bg-indigo-600/10 rounded-2xl">
				<span class="font-bold text-sm">%%s</span> <a href="/" class="text-xs">Logout</a>
			</div>
		</aside>
		<main class="flex-1 p-10 overflow-y-auto">
			<h1 class="text-5xl font-black mb-8">Explore <span class="text-indigo-400">Sanctuaries</span></h1>
			<div id="results-area" class="grid grid-cols-1 lg:grid-cols-2 gap-8"></div>
		</main>
	</body></html>`, currentUsername)
}

func GetSignupHTML() string {
	return `<!DOCTYPE html><html><head><script src="https://cdn.tailwindcss.com"></script></head>
	<body class="bg-[#0a0a0a] flex items-center justify-center min-h-screen">
		<form action="/signup" method="POST" class="bg-slate-900/40 p-10 rounded-[2.5rem] border border-white/5 w-full max-w-md text-center">
			<h1 class="text-4xl font-black text-white mb-8">Create Account</h1>
			<input type="text" name="username" placeholder="Username" class="w-full p-4 rounded-2xl bg-black border border-white/10 text-white mb-4">
			<button type="submit" class="w-full bg-indigo-600 py-5 rounded-2xl font-bold text-white">Start Journey</button>
		</form>
	</body></html>`
}

func GetScripts(isLoggedIn bool, currentUsername string) string {
	return `<script>
		async function fetchHouses() {
			const res = await fetch('/houses');
			const data = await res.json();
			const container = document.getElementById('results-area');
			container.innerHTML = data.map(h => '<div class="bg-slate-900/40 p-8 rounded-[2.5rem] border border-white/5"><div class="relative h-64 bg-slate-800 rounded-[2rem] overflow-hidden mb-6"><img src="'+h.image_urls[0]+'" class="w-full h-full object-cover"></div><h2 class="text-3xl font-bold">'+h.building_name+'</h2><button class="w-full bg-indigo-500 py-5 mt-6 rounded-2xl font-bold">Pay KES 1,000 to View</button></div>').join("");
		}
		window.onload = fetchHouses;
	</script>`
}