package templates

import "fmt"

func GetLandingHTML() string {
	return `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nyumba - Find Your Sanctuary</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Inter', sans-serif; }
        @keyframes marquee {
            0% { transform: translateX(0); }
            100% { transform: translateX(-50%); }
        }
        .marquee-container {
            overflow: hidden;
            white-space: nowrap;
            position: relative;
        }
        .marquee-content {
            display: inline-block;
            animation: marquee 40s linear infinite;
        }
        .marquee-content:hover {
            animation-play-state: paused;
        }
        .bg-mesh {
            background-color: #020617;
            background-image: 
                radial-gradient(at 0% 0%, rgba(30, 58, 138, 0.3) 0px, transparent 50%),
                radial-gradient(at 100% 0%, rgba(20, 184, 166, 0.2) 0px, transparent 50%),
                radial-gradient(at 50% 100%, rgba(88, 28, 135, 0.2) 0px, transparent 50%);
        }
    </style>
</head>
<body class="bg-mesh text-white min-h-screen">
    <!-- Navbar -->
    <div class="fixed top-8 left-1/2 -translate-x-1/2 z-50 w-full max-w-4xl px-4">
        <nav class="backdrop-blur-xl bg-black/40 border border-white/10 rounded-full px-8 py-3 flex justify-between items-center shadow-2xl">
            <div class="flex items-center gap-12">
                <h1 class="text-xl font-extrabold tracking-tighter">Nyumba.</h1>
                <div class="hidden md:flex items-center gap-8 text-sm font-medium text-zinc-400">
                    <a href="#" class="hover:text-white transition-colors">How it Works</a>
                    <a href="#neighborhoods" class="hover:text-white transition-colors">Neighborhoods</a>
                    <a href="#" class="hover:text-white transition-colors">For Landlords</a>
                </div>
            </div>
            <div class="flex items-center gap-6">
                <a href="#" class="text-sm font-medium text-zinc-400 hover:text-white transition-colors">Sign In</a>
                <a href="/explore" class="bg-gradient-to-r from-blue-600 to-indigo-600 text-white px-6 py-2 rounded-full text-sm font-bold hover:shadow-[0_0_20px_rgba(37,99,235,0.4)] transition-all">Explore</a>
            </div>
        </nav>
    </div>

    <!-- Hero Section -->
    <main class="relative pt-48 pb-32 flex flex-col items-center text-center px-6">
        <div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full border border-white/10 bg-white/5 text-[10px] font-bold tracking-widest uppercase mb-12">
            <span class="w-2 h-2 rounded-full bg-emerald-500 shadow-[0_0_8px_rgba(16,185,129,0.6)]"></span>
            Verified Listings Only
        </div>

        <h2 class="text-6xl md:text-8xl font-extrabold tracking-tight mb-4 leading-none">
            Find Your Sanctuary.
        </h2>
        <h3 class="text-6xl md:text-8xl font-extrabold tracking-tight mb-12 leading-none bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-cyan-400 to-emerald-400">
            Simplified.
        </h3>

        <p class="max-w-xl text-lg md:text-xl text-zinc-400 mb-12 leading-relaxed">
            An exclusive platform connecting serious renters with verified landlords. No agents. No endless scrolling. Just your next home.
        </p>

        <a href="/explore" class="group bg-white text-black px-10 py-5 rounded-full text-lg font-bold hover:scale-105 transition-all flex items-center gap-3">
            Start Your Search
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" class="group-hover:translate-x-1 transition-transform"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
        </a>

        <!-- Stats Section -->
        <div class="mt-24 flex gap-16 md:gap-32">
            <div class="flex flex-col items-center">
                <span class="text-4xl md:text-5xl font-extrabold mb-2">500+</span>
                <span class="text-[10px] uppercase tracking-[0.2em] text-zinc-500 font-bold">Verified Listings</span>
            </div>
            <div class="flex flex-col items-center">
                <span class="text-4xl md:text-5xl font-extrabold mb-2 text-emerald-500">0</span>
                <span class="text-[10px] uppercase tracking-[0.2em] text-zinc-500 font-bold">Scam Reports</span>
            </div>
        </div>
    </main>

    <!-- Neighborhoods Section -->
    <section id="neighborhoods" class="pb-32">
        <div class="max-w-7xl mx-auto px-6 mb-12 text-center">
            <h3 class="text-[10px] uppercase tracking-[0.3em] text-zinc-500 font-bold">Popular Neighborhoods</h3>
        </div>
        
        <div class="marquee-container">
            <div class="marquee-content flex gap-4 px-4">
                <div class="flex gap-4">
                    <a href="/explore?neighborhood=Westlands" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Westlands</a>
                    <a href="/explore?neighborhood=Kilimani" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Kilimani</a>
                    <a href="/explore?neighborhood=Karen" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Karen</a>
                    <a href="/explore?neighborhood=Thika" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Thika</a>
                    <a href="/explore?neighborhood=Langata" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Langata</a>
                    <a href="/explore?neighborhood=Ruiru" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Ruiru</a>
                    <a href="/explore?neighborhood=Kileleshwa" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Kileleshwa</a>
                    <a href="/explore?neighborhood=Lavington" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Lavington</a>
                </div>
                <div class="flex gap-4">
                    <a href="/explore?neighborhood=Westlands" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Westlands</a>
                    <a href="/explore?neighborhood=Kilimani" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Kilimani</a>
                    <a href="/explore?neighborhood=Karen" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Karen</a>
                    <a href="/explore?neighborhood=Thika" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Thika</a>
                    <a href="/explore?neighborhood=Langata" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Langata</a>
                    <a href="/explore?neighborhood=Ruiru" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Ruiru</a>
                    <a href="/explore?neighborhood=Kileleshwa" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Kileleshwa</a>
                    <a href="/explore?neighborhood=Lavington" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-bold text-lg backdrop-blur-sm">Lavington</a>
                </div>
            </div>
        </div>
    </section>
</body>
</html>`
}


func GetHTML(user string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Explore Sanctuaries - Nyumba</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-zinc-950 text-white font-sans">
    <nav class="p-6 flex justify-between items-center max-w-7xl mx-auto border-b border-zinc-800">
        <h1 class="text-2xl font-bold tracking-tighter">Nyumba.</h1>
        <div class="flex items-center gap-6">
            <span class="text-zinc-400">Welcome, %s</span>
            <a href="/" class="text-zinc-400 hover:text-white transition-colors">Home</a>
        </div>
    </nav>
    <main class="max-w-7xl mx-auto px-6 py-12">
        <div class="flex justify-between items-end mb-12">
            <div>
                <h2 class="text-4xl font-bold" id="explore-title">Available Sanctuaries</h2>
                <p class="text-zinc-500 mt-2" id="filter-status">Showing all verified listings</p>
            </div>
            <a href="/explore" id="clear-filter" class="hidden text-sm text-zinc-400 hover:text-white underline">Clear filters</a>
        </div>
        <div id="houses-grid" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            <!-- Loaded via JS -->
        </div>
    </main>
    <script>
        const urlParams = new URLSearchParams(window.location.search);
        const neighborhoodFilter = urlParams.get('neighborhood');

        if (neighborhoodFilter) {
            document.getElementById('explore-title').innerText = 'Sanctuaries in ' + neighborhoodFilter;
            document.getElementById('filter-status').innerText = 'Showing listings for ' + neighborhoodFilter;
            document.getElementById('clear-filter').classList.remove('hidden');
        }

        fetch('/houses')
            .then(res => res.json())
            .then(houses => {
                const grid = document.getElementById('houses-grid');
                const filteredHouses = neighborhoodFilter 
                    ? houses.filter(h => h.location.toLowerCase().includes(neighborhoodFilter.toLowerCase()))
                    : houses;

                if (filteredHouses.length === 0) {
                    grid.innerHTML = '<div class="col-span-full py-20 text-center text-zinc-500">No sanctuaries found in this neighborhood yet.</div>';
                    return;
                }

                filteredHouses.forEach(house => {
                    const card = document.createElement('div');
                    card.className = 'bg-zinc-900 rounded-3xl overflow-hidden border border-zinc-800 hover:border-zinc-700 transition-all group';
                    card.innerHTML = %s
                        <div class="aspect-video overflow-hidden">
                            <img src="${house.image_urls[0]}" alt="${house.building_name}" class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500">
                        </div>
                        <div class="p-6">
                            <div class="flex justify-between items-start mb-4">
                                <div>
                                    <h3 class="text-xl font-bold">${house.building_name}</h3>
                                    <p class="text-zinc-500 text-sm">${house.location}</p>
                                </div>
                                <span class="bg-white text-black px-3 py-1 rounded-full text-sm font-bold">KSh ${house.price.toLocaleString()}</span>
                            </div>
                            <div class="flex gap-4 pt-4 border-t border-zinc-800">
                                <span class="text-sm text-zinc-400 flex items-center gap-1">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2 4v16"/><path d="M2 8h18a2 2 0 0 1 2 2v10"/><path d="M2 17h20"/><path d="M6 8v9"/></svg>
                                    ${house.bedrooms} Bed
                                </span>
                                <span class="text-sm text-zinc-400 flex items-center gap-1">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 6 6.5 3.5a1.5 1.5 0 0 0-2.12 0 1.5 1.5 0 0 0 0 2.12L7 8"/><rect width="16" height="12" x="2" y="8" rx="2"/><path d="M7 12h.01"/><path d="M17 12h.01"/><path d="M12 12h.01"/></svg>
                                    ${house.bathrooms} Bath
                                </span>
                            </div>
                        </div>
                    %s;
                    grid.appendChild(card);
                });
            });
    </script>
</body>
</html>`, user, "`", "`")
}
