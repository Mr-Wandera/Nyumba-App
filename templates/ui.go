package templates

func GetLandingHTML() string {
	return `<!DOCTYPE html>
<html>
<head>
    <title>Nyumba | Find Your Sanctuary</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet">
    <style>*{font-family:'Outfit',sans-serif;}</style>
</head>
<body class="bg-[#0a0a0a] text-white min-h-screen">
    <nav class="flex justify-between items-center px-6 py-6">
        <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-br from-indigo-500 to-cyan-400 rounded-xl flex items-center justify-center text-white font-bold">N</div>
            <span class="text-2xl font-bold">Nyumba.</span>
        </div>
        <a href="/explore" class="bg-white text-black px-6 py-3 rounded-full font-semibold">Start Searching</a>
    </nav>
    <section class="flex flex-col items-center justify-center min-h-[70vh] px-6 text-center">
        <h1 class="text-5xl md:text-7xl font-black mb-6 leading-tight">Find Your <br><span style="background:linear-gradient(135deg,#818cf8,#22d3ee);-webkit-background-clip:text;-webkit-text-fill-color:transparent;">Sanctuary.</span></h1>
        <p class="text-lg text-white/60 max-w-2xl mb-10">An exclusive platform connecting serious renters with verified landlords. No agents. No endless scrolling. Just your next home.</p>
        <a href="/explore" class="bg-gradient-to-r from-indigo-600 to-indigo-500 text-white px-8 py-4 rounded-full font-bold text-lg">Start Your Search</a>
        <div class="mt-16 grid grid-cols-3 gap-8 md:gap-16">
            <div class="text-center"><div class="text-3xl font-bold">500+</div><div class="text-xs text-white/50 uppercase tracking-wider">Verified Listings</div></div>
            <div class="text-center"><div class="text-3xl font-bold text-emerald-400">0</div><div class="text-xs text-white/50 uppercase tracking-wider">Scam Reports</div></div>
            <div class="text-center"><div class="text-3xl font-bold text-cyan-400">KES 1K</div><div class="text-xs text-white/50 uppercase tracking-wider">To Connect</div></div>
        </div>
    </section>
    <div class="w-full overflow-hidden py-6 border-y border-white/5 bg-black/20">
        <div style="display:flex;white-space:nowrap;animation:ticker 30s linear infinite;">
            <span class="mx-8 text-white/30">Thika Town</span>
            <span class="mx-8 text-indigo-400 font-semibold">Section 9</span>
            <span class="mx-8 text-white/30">Ngoingwa</span>
            <span class="mx-8 text-cyan-400 font-semibold">Kenyatta Road</span>
            <span class="mx-8 text-white/30">Makongeni</span>
            <span class="mx-8 text-indigo-400 font-semibold">Kamakis</span>
            <span class="mx-8 text-white/30">Ruiru</span>
            <span class="mx-8 text-cyan-400 font-semibold">Juja</span>
        </div>
    </div>
    <style>@keyframes ticker{0%{transform:translateX(0);}100%{transform:translateX(-50%);}}</style>
</body>
</html>`
}

func GetHTML(currentUsername string) string {
	return `<!DOCTYPE html>
<html>
<head>
    <title>Nyumba | Explore</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&display=swap" rel="stylesheet">
    <style>*{font-family:'Outfit',sans-serif;}</style>
</head>
<body class="bg-[#0a0a0a] text-white h-screen overflow-hidden">
    <div class="flex h-screen">
        <aside class="w-[350px] bg-[#111] border-r border-white/10 p-6 flex flex-col">
            <h1 class="text-2xl font-bold mb-6">Nyumba.</h1>
            <form action="/add-house" method="POST" enctype="multipart/form-data" class="space-y-4">
                <input type="text" name="building_name" placeholder="Property Name" class="w-full p-3 rounded-xl bg-black/50 border border-white/10 text-white" required>
                <input type="text" name="location" placeholder="Location" class="w-full p-3 rounded-xl bg-black/50 border border-white/10 text-white">
                <input type="url" name="map_link" placeholder="Google Maps Link" class="w-full p-3 rounded-xl bg-black/50 border border-white/10 text-white">
                <input type="number" name="price" placeholder="Monthly Rent (KES)" class="w-full p-3 rounded-xl bg-black/50 border border-white/10 text-white">
                <input type="file" name="property_photo" accept="image/*" class="w-full p-3 text-white/60 text-sm">
                <button type="submit" class="w-full bg-gradient-to-r from-indigo-600 to-cyan-500 text-white py-3 rounded-xl font-bold">Publish Listing</button>
            </form>
        </aside>
        <main class="flex-1 flex flex-col overflow-hidden">
            <header class="px-8 py-4 border-b border-white/10 flex justify-between items-center">
                <h1 class="text-2xl font-bold">Available Sanctuaries</h1>
                <a href="/" class="px-4 py-2 rounded-full border border-white/20 text-sm hover:bg-white/10">Back</a>
            </header>
            <div id="results-area" class="flex-1 overflow-y-auto p-8">
                <div class="flex items-center justify-center h-full">
                    <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-500"></div>
                </div>
            </div>
        </main>
    </div>
    <script>
        async function fetchHouses() {
            try {
                const res = await fetch('/houses');
                const data = await res.json();
                const container = document.getElementById('results-area');
                if (data.length === 0) {
                    container.innerHTML = '<div class="text-center text-white/60">No listings yet. Add one!</div>';
                    return;
                }
                let html = '<div class="grid grid-cols-1 lg:grid-cols-2 gap-6">';
                data.forEach(h => {
                    html += '<div class="bg-[#1a1a1a] rounded-2xl overflow-hidden border border-white/10 hover:border-indigo-500/30 transition">';
                    html += '<div class="relative h-48">';
                    html += '<img src="' + (h.image_urls[0] || '/uploads/default.jpg') + '" class="w-full h-full object-cover">';
                    html += '<div class="absolute top-4 left-4 bg-emerald-500/20 px-3 py-1 rounded-full text-xs text-emerald-300 border border-emerald-500/30">✓ Verified</div>';
                    html += '<div class="absolute bottom-4 left-4"><p class="text-xl font-bold">KES ' + (h.price || 15000) + '<span class="text-sm text-white/60">/mo</span></p><p class="text-sm text-white/60">📍 ' + (h.location || 'Thika') + '</p></div></div>';
                    html += '<div class="p-4"><h3 class="text-lg font-bold mb-2">' + h.building_name + '</h3>';
                    html += '<button onclick="alert(\'Payment integration coming soon!\')" class="w-full bg-gradient-to-r from-indigo-600 to-cyan-500 text-white py-3 rounded-xl font-bold flex items-center justify-center gap-2"><span>🔒</span>Unlock for KES 1,000</button>';
                    html += '</div></div>';
                });
                html += '</div>';
                container.innerHTML = html;
            } catch (err) {
                container.innerHTML = '<div class="text-red-400 text-center">Error loading listings</div>';
            }
        }
        window.addEventListener('load', fetchHouses);
    </script>
</body>
</html>`
}