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
			try {
				const res = await fetch('/houses');
				const data = await res.json();
				const container = document.getElementById('results-area');
				
				container.innerHTML = data.map(h => ` + "`" + `
					<div class="bg-slate-900/40 p-8 rounded-[2.5rem] border border-white/5 group transition-all">
						<img src="${h.image_urls && h.image_urls.length > 0 ? h.image_urls[0] : '/uploads/default.jpg'}" 
							 class="rounded-[2rem] h-64 w-full object-cover mb-6 group-hover:scale-105 transition duration-500">
						
						<div class="flex justify-between items-start mb-6">
							<div>
								<h2 class="text-3xl font-bold tracking-tight">${h.building_name}</h2>
								<p class="text-slate-500 text-sm font-medium">📍 ${h.location}</p>
							</div>
							${h.map_link ? ` + "`" + `<a href="${h.map_link}" target="_blank" class="text-indigo-400 font-bold text-xs hover:text-white transition">📍 Map</a>` + "`" + ` : ''}
						</div>

						${h.is_booked ? ` + "`" + `
							<div class="bg-green-500/10 border border-green-500/30 rounded-3xl p-6 mt-4 animate-in fade-in zoom-in duration-500">
								<div class="flex items-center gap-4 mb-4">
									<div class="w-12 h-12 rounded-full bg-green-500/20 flex items-center justify-center text-green-400 text-xl">✓</div>
									<div>
										<p class="text-green-400 font-black tracking-tight">Access Granted</p>
										<p class="text-white/40 text-xs font-bold uppercase tracking-widest">Sanctuary Unlocked</p>
									</div>
								</div>
								<div class="space-y-3 mt-3">
									<a href="tel:${h.phone}" class="flex items-center justify-center gap-3 bg-white/5 hover:bg-white/10 text-white py-4 rounded-2xl font-bold transition">
										📞 Call Owner
									</a>
									<a href="https://wa.me/${h.phone.replace('+', '')}" target="_blank" class="flex items-center justify-center gap-3 bg-green-500/10 hover:bg-green-500/20 text-green-400 py-4 rounded-2xl font-bold transition">
										💬 WhatsApp Chat
									</a>
								</div>
							</div>` + "`" + ` : ` + "`" + `
							<button onclick="handlePayment(${h.id})" class="w-full bg-indigo-600 hover:bg-indigo-500 text-white py-5 rounded-2xl font-black shadow-lg shadow-indigo-600/20 transition-all active:scale-95">
								Pay KES 1,000 to View
							</button>` + "`" + `}
					</div>` + "`" + `).join("");
			} catch (err) {
				console.error("Failed to load sanctuaries:", err);
			}
		}

		function handlePayment(houseId) {
			// Trigger your STK Push handler here
			fetch('/pay', { 
				method: 'POST', 
				body: JSON.stringify({ house_id: houseId }) 
			}).then(() => alert("Check your phone for the M-Pesa prompt!"));
		}

		window.onload = fetchHouses;
	</script>`
}