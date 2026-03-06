package templates

import "html/template"

func GetLandingHTML() string {
	return `<!DOCTYPE html>
<html>
<head>
	<title>Nyumba</title>
	<script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-[#0a0a0a] text-white flex flex-col items-center justify-center h-screen">
	<h1 class="text-7xl font-black mb-8 tracking-tighter">Find Your Sanctuary.</h1>
	<a href="/explore" class="bg-white text-black px-10 py-5 rounded-full font-black hover:scale-105 transition">Start Search</a>
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
			if (!res.ok) throw new Error('Failed to fetch');
			const data = await res.json();
			const container = document.getElementById('results-area');
			
			if (data.length === 0) {
				container.innerHTML = '<p class="text-gray-500 col-span-2 text-center">No listings yet.</p>';
				return;
			}
			
			container.innerHTML = data.map(h => 
				'<div class="bg-slate-900/40 p-6 rounded-[2rem] border border-white/5 group hover:border-white/10 transition">' +
					'<img src="' + (h.image_urls?.[0] || '/uploads/default.jpg') + '" class="h-48 w-full object-cover rounded-xl mb-4 group-hover:scale-105 transition duration-300">' +
					'<div class="flex justify-between items-center mb-2">' +
						'<h2 class="text-2xl font-bold">' + (h.building_name || 'Unnamed Property') + '</h2>' +
						(h.map_link ? '<a href="' + h.map_link + '" target="_blank" rel="noopener" class="text-indigo-400 text-xs font-bold hover:text-indigo-300">📍 Map</a>' : '') +
					'</div>' +
					(h.location ? '<p class="text-gray-400 text-sm mb-4">📍 ' + h.location + '</p>' : '') +
					'<button class="w-full bg-indigo-600 hover:bg-indigo-500 py-4 rounded-xl font-bold transition">Pay KES 1,000</button>' +
				'</div>'
			).join('');
		} catch (err) {
			console.error('Error:', err);
			document.getElementById('results-area').innerHTML = '<p class="text-red-500 col-span-2 text-center">Failed to load listings.</p>';
		}
	}
	window.addEventListener('load', fetchHouses);
	</script>`
}