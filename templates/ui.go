package templates

import (
	"fmt"
	"html/template"
)

// GetLandingHTML returns the entry point for the application
func GetLandingHTML() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
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
	<main class="flex-1 flex flex-col items-center justify-center pt-32 pb-12 px-6 text-center">
		<h1 class="text-4xl md:text-6xl lg:text-7xl font-black tracking-tighter mb-8 leading-[0.9]">
			Find Your <span class="text-white">Sanctuary.</span><br>
			<span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-400">Directly.</span>
		</h1>
		<a href="/explore" class="bg-white text-black px-10 py-5 rounded-full font-black text-lg hover:scale-105 transition mb-12 inline-block">Start Your Search →</a>
	</main>

	<div class="flex justify-center gap-8 md:gap-24 py-16 border-y border-white/5 bg-white/[0.02]">
		<div class="text-center">
			<div class="text-3xl md:text-4xl font-black text-indigo-400 tracking-tighter">500+</div>
			<div class="text-white/40 text-[10px] font-bold uppercase tracking-[0.2em] mt-2">Verified Listings</div>
		</div>
		<div class="text-center">
			<div class="text-3xl md:text-4xl font-black text-indigo-400 tracking-tighter">0</div>
			<div class="text-white/40 text-[10px] font-bold uppercase tracking-[0.2em] mt-2">Scam Reports</div>
		</div>
		<div class="text-center">
			<div class="text-3xl md:text-4xl font-black text-indigo-400 tracking-tighter">KES 1K</div>
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

// GetHTML serves the exploration dashboard with the responsive sidebar and grid
func GetHTML(currentUsername string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Nyumba | Explore</title>
	<script src="https://cdn.tailwindcss.com"></script>
	<link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;600;900&display=swap" rel="stylesheet">
	<style>body{font-family:'Outfit',sans-serif; background:#0a0a0a; color:white;}</style>
</head>
<body class="min-h-screen flex flex-col md:flex-row bg-[#0a0a0a] text-white">
	<aside class="w-full md:w-[380px] border-b md:border-b-0 md:border-r border-white/5 p-8 flex flex-col bg-slate-900/20 backdrop-blur-xl">
		<div class="mb-12">
			<h1 class="text-4xl font-black tracking-tighter">Nyumba<span class="text-indigo-500">.</span></h1>
			<p class="text-[10px] text-slate-500 font-bold uppercase tracking-[0.3em] mt-2">Welcome, %s</p>
		</div>

		<form action="/add-house" method="POST" enctype="multipart/form-data" class="space-y-5">
			<input type="text" name="building_name" placeholder="Building Name" class="w-full p-4 rounded-2xl bg-white/5 border border-white/10 text-sm focus:border-indigo-500 outline-none transition" required>
			<input type="text" name="location" placeholder="Location (e.g. Section 9)" class="w-full p-4 rounded-2xl bg-white/5 border border-white/10 text-sm focus:border-indigo-500 outline-none transition" required>
			<input type="text" name="map_link" placeholder="📍 Google Maps Link" class="w-full p-4 rounded-2xl bg-white/5 border border-white/10 text-sm focus:border-indigo-500 outline-none transition">
			
			<div class="relative group">
				<input type="file" name="property_photo" accept="image/*" class="absolute inset-0 w-full h-full opacity-0 cursor-pointer z-10">
				<div class="w-full bg-white/5 text-slate-400 border border-white/10 py-4 rounded-2xl font-bold text-xs text-center group-hover:bg-white/10 transition">
					Upload Property Photo
				</div>
			</div>

			<button type="submit" class="w-full bg-white text-black py-5 rounded-2xl font-black hover:bg-slate-200 transition-all active:scale-95 shadow-xl shadow-white/5">Publish Sanctuary</button>
		</form>
	</aside>

	<main class="flex-1 p-8 md:p-12 overflow-y-auto">
		<header class="mb-12">
			<h2 class="text-4xl md:text-5xl font-black tracking-tighter">Available <span class="text-indigo-500">Sanctuaries</span></h2>
			<p class="text-slate-500 mt-2 font-medium">Verified listings direct from owners in Thika.</p>
		</header>

		<div id="results-area" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
			</div>
	</main>

	%s
</body>
</html>`, template.HTMLEscapeString(currentUsername), GetScripts())
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