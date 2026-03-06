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
		.animate-scroll { display: flex; white-space: nowrap; animation: scroll 40s linear infinite; }
		.animate-scroll:hover { animation-play-state: paused; }
	</style>
</head>
<body class="min-h-screen flex flex-col">
	<main class="flex-col items-center justify-center pt-32 pb-12 px-6 text-center">
		<h1 class="text-6xl md:text-8xl font-black tracking-tighter mb-8 leading-[0.9]">
			Find Your <span class="text-white">Sanctuary.</span><br>
			<span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-400">Directly.</span>
		</h1>
		<a href="/explore" class="bg-white text-black px-10 py-5 rounded-full font-black text-lg hover:scale-105 transition mb-12 inline-block">Start Your Search →</a>
	</main>

	<div class="flex justify-center gap-12 md:gap-24 py-16 border-y border-white/5 bg-white/[0.02]">
		<div class="text-center">
			<div class="text-4xl font-black text-indigo-400 tracking-tighter">500+</div>
			<div class="text-white/40 text-[10px] font-bold uppercase tracking-[0.2em] mt-2">Verified Listings</div>
		</div>
		<div class="text-center">
			<div class="text-4xl font-black text-indigo-400 tracking-tighter">0</div>
			<div class="text-white/40 text-[10px] font-bold uppercase tracking-[0.2em] mt-2">Scam Reports</div>
		</div>
		<div class="text-center">
			<div class="text-4xl font-black text-indigo-400 tracking-tighter">KES 1K</div>
			<div class="text-white/40 text-[10px] font-bold uppercase tracking-[0.2em] mt-2">To Connect</div>
		</div>
	</div>

	<section class="w-full overflow-hidden py-12 relative">
		<div class="animate-scroll gap-24">
			<span class="text-2xl font-black uppercase tracking-widest text-slate-700">Section 9 •</span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-700">Ngoingwa •</span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-700">Kenyatta Road •</span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-700">Section 9 •</span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-700">Ngoingwa •</span>
			<span class="text-2xl font-black uppercase tracking-widest text-slate-700">Kenyatta Road •</span>
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
					<div class="group relative bg-gradient-to-b from-white/5 to-white/[0.02] backdrop-blur-xl rounded-[2.5rem] border border-white/10 overflow-hidden hover:border-indigo-500/30 transition-all duration-500 hover:shadow-[0_0_40px_-10px_rgba(99,102,241,0.3)]">
						
						<div class="relative h-72 overflow-hidden">
							<img src="${h.image_urls && h.image_urls.length > 0 ? h.image_urls[0] : '/uploads/default.jpg'}" 
								 class="w-full h-full object-cover group-hover:scale-110 transition duration-700">
							<div class="absolute inset-0 bg-gradient-to-t from-[#0a0a0a] via-transparent to-transparent"></div>
							
							<div class="absolute top-6 left-6 bg-indigo-500/20 backdrop-blur-md border border-indigo-500/30 text-indigo-300 px-4 py-1.5 rounded-full text-[10px] font-black uppercase tracking-widest flex items-center gap-2">
								<span class="w-2 h-2 bg-indigo-400 rounded-full animate-pulse"></span> Verified Sanctuary
							</div>
							
							<div class="absolute bottom-6 left-6">
								<p class="text-3xl font-black text-white tracking-tighter">KES ${h.price.toLocaleString()}<span class="text-sm font-normal text-white/50">/mo</span></p>
								<p class="text-white/60 text-xs font-bold uppercase tracking-widest mt-1">📍 ${h.location}</p>
							</div>
						</div>
						
						<div class="p-8">
							<h3 class="text-2xl font-black text-white tracking-tight mb-3">${h.building_name}</h3>
							<div class="flex gap-6 text-xs font-bold text-white/40 uppercase tracking-widest mb-8">
								<span class="flex items-center gap-2">🛏 ${h.type || '2 Beds'}</span>
								<span class="flex items-center gap-2">📐 ${h.details || 'Ready'}</span>
							</div>
							
							${h.is_booked ? ` + "`" + `
								<div class="space-y-3 animate-in fade-in slide-in-from-bottom-4 duration-700">
									<a href="tel:${h.phone}" class="w-full bg-green-500/10 hover:bg-green-500/20 text-green-400 py-5 rounded-2xl font-black transition flex items-center justify-center gap-3 border border-green-500/20">
										📞 Call Owner Directly
									</a>
									<a href="https://wa.me/${h.phone.replace('+', '')}" target="_blank" class="w-full bg-white/5 hover:bg-white/10 text-white py-4 rounded-2xl font-bold transition flex items-center justify-center gap-3">
										💬 WhatsApp Chat
									</a>
								</div>` + "`" + ` : ` + "`" + `
								<button onclick="handlePayment(${h.id})" class="w-full bg-white text-black py-5 rounded-2xl font-black hover:bg-indigo-50 transition-all flex items-center justify-center gap-3 shadow-xl active:scale-95">
									<span>🔓</span> Unlock for KES 1,000
								</button>` + "`" + `}
						</div>
					</div>` + "`" + `).join("");
			} catch (err) {
				console.error("Sanctuary Load Error:", err);
			}
		}

		async function handlePayment(houseId) {
			// Trigger M-Pesa STK Push
			alert("Connecting to M-Pesa for Sanctuary ID: " + houseId);
			const res = await fetch('/pay', { method: 'POST', body: JSON.stringify({ id: houseId }) });
		}

		window.onload = fetchHouses;
	</script>`
}