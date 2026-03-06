package templates

import "html/template"

func GetLandingHTML() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
	<title>Nyumba | Sanctuary</title>
	<script src="https://cdn.tailwindcss.com"></script>
	<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;900&display=swap" rel="stylesheet">
	<style>
		body { font-family: 'Outfit', sans-serif; background: #0a0a0a; color: white; overflow-x: hidden; }
		@keyframes scroll {
			0% { transform: translateX(0); }
			100% { transform: translateX(-50%); }
		}
		.animate-scroll {
			display: flex;
			white-space: nowrap;
			animation: scroll 40s linear infinite;
		}
		.animate-scroll:hover { animation-play-state: paused; }
		.scroll-container {
			mask-image: linear-gradient(to right, transparent, black 15%, black 85%, transparent);
			-webkit-mask-image: linear-gradient(to right, transparent, black 15%, black 85%, transparent);
		}
	</style>
</head>
<body class="min-h-screen flex flex-col">
	<main class="flex-1 flex flex-col items-center justify-center pt-32 pb-12 px-6 text-center">
		<h1 class="text-6xl md:text-8xl font-black tracking-tighter mb-8 leading-[0.9]">
			Find Your <span class="text-white">Sanctuary.</span><br>
			<span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-400">Simplified.</span>
		</h1>
		<a href="/explore" class="bg-white text-black px-10 py-5 rounded-full font-black text-lg hover:scale-105 transition">Start Your Search →</a>
	</main>

	<section class="w-full overflow-hidden bg-black/40 border-y border-white/5 py-12 relative scroll-container">
		<div class="animate-scroll gap-24">
			<span class="text-2xl font-black uppercase tracking-widest text-slate-500 hover:text-indigo-400 transition">Thika Town <span class="text-indigo-500 mx-8">•</span></span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-500 hover:text-indigo-400 transition">Section 9 <span class="text-indigo-500 mx-8">•</span></span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-500 hover:text-indigo-400 transition">Ngoingwa <span class="text-indigo-500 mx-8">•</span></span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-500 hover:text-indigo-400 transition">Landless <span class="text-indigo-500 mx-8">•</span></span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-500 hover:text-indigo-400 transition">Kenyatta Road <span class="text-indigo-500 mx-8">•</span></span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-500 hover:text-indigo-400 transition">Thika Town <span class="text-indigo-500 mx-8">•</span></span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-500 hover:text-indigo-400 transition">Section 9 <span class="text-indigo-500 mx-8">•</span></span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-500 hover:text-indigo-400 transition">Ngoingwa <span class="text-indigo-500 mx-8">•</span></span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-500 hover:text-indigo-400 transition">Landless <span class="text-indigo-500 mx-8">•</span></span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-500 hover:text-indigo-400 transition">Kenyatta Road <span class="text-indigo-500 mx-8">•</span></span>
		</div>
	</section>
</body>
</html>`
}

func GetHTML(currentUsername string) string {
	// Use strings.Builder or fmt.Sprintf with script injection
	html := `<!DOCTYPE html>
<html>
<head>
	<script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="h-screen flex bg-[#0a0a0a] text-white overflow-hidden">
	<aside class="w-[350px] border-r border-white/5 p-6 flex flex-col">
		<h1 class="text-3xl font-black mb-10">Nyumba.</h1>
		<p class="text-sm text-gray-400 mb-4">Welcome, ` + template.HTMLEscapeString(currentUsername) + `</p>
		<form action="/add-house" method="POST" enctype="multipart/form-data" class="space-y-4">
			<input type="text" name="building_name" placeholder="Apartment Name" class="w-full p-3 rounded-xl bg-slate-900 border border-white/5 text-white" required>
			<input type="text" name="location" placeholder="Location" class="w-full p-3 rounded-xl bg-slate-900 border border-white/5 text-white">
			<input type="url" name="map_link" placeholder="📍 Google Maps Link" class="w-full p-3 rounded-xl bg-slate-900 border border-white/5 text-white">
			<input type="file" name="property_photo" accept="image/*" class="text-xs text-gray-400">
			<button type="submit" class="w-full bg-white text-black py-4 rounded-xl font-black hover:bg-gray-200 transition">Publish</button>
		</form>
	</aside>
	<main id="results-area" class="flex-1 p-10 grid grid-cols-2 gap-8 overflow-y-auto"></main>
	` + GetScripts() + `
</body>
</html>`
	return html
}

func GetScripts() string {
	return `<script>
		async function fetchHouses() {
			const res = await fetch('/houses');
			const data = await res.json();
			const container = document.getElementById('results-area');
			
			container.innerHTML = data.map(h => ` + "`" + `
				<div class="bg-slate-900/40 p-6 rounded-[2rem] border border-white/5 group">
					<img src="${h.image_urls && h.image_urls.length > 0 ? h.image_urls[0] : '/uploads/default.jpg'}" class="h-48 w-full object-cover rounded-xl mb-4 group-hover:scale-105 transition duration-300">
					<div class="flex justify-between items-center mb-4">
						<h2 class="text-2xl font-bold">${h.building_name}</h2>
						${h.map_link ? ` + "`" + `<a href="${h.map_link}" target="_blank" class="text-indigo-400 font-bold text-xs">📍 Map</a>` + "`" + ` : ''}
					</div>
					<button onclick="handlePayment(${h.id})" class="w-full bg-indigo-600 py-4 rounded-xl font-bold">Pay KES 1,000</button>
				</div>` + "`" + `).join("");
		}
		window.onload = fetchHouses;
	</script>`
}