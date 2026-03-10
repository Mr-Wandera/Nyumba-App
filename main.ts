import express from "express";
import fs from "fs";
import path from "path";

const app = express();
const PORT = 3000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use("/uploads", express.static("uploads"));

// Mock data storage
let houses = JSON.parse(fs.readFileSync("houses.json", "utf-8"));

// Templates (mimicking Go templates)
const getLandingHTML = () => `
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
            background-size: 200% 200%;
            animation: meshGradient 20s ease infinite;
        }
        @keyframes meshGradient {
            0% { background-position: 0% 0%; }
            50% { background-position: 100% 100%; }
            100% { background-position: 0% 0%; }
        }
        @keyframes fadeInUp {
            from { opacity: 0; transform: translateY(20px); }
            to { opacity: 1; transform: translateY(0); }
        }
        .animate-fade-in-up {
            animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
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
                    <a href="/landlord" class="hover:text-white transition-colors">For Landlords</a>
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
        <div class="animate-fade-in-up inline-flex items-center gap-2 px-4 py-1.5 rounded-full border border-white/10 bg-white/5 text-[10px] font-bold tracking-widest uppercase mb-12">
            <span class="w-2 h-2 rounded-full bg-emerald-500 shadow-[0_0_8px_rgba(16,185,129,0.6)]"></span>
            Verified Listings Only
        </div>

        <h2 class="animate-fade-in-up [animation-delay:200ms] text-6xl md:text-9xl font-extrabold tracking-tighter mb-4 leading-[0.85]">
            Find Your Sanctuary.
        </h2>
        <h3 class="animate-fade-in-up [animation-delay:400ms] text-6xl md:text-9xl font-extrabold tracking-tighter mb-12 leading-[0.85] bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-cyan-400 to-emerald-400">
            Simplified.
        </h3>

        <p class="animate-fade-in-up [animation-delay:600ms] max-w-xl text-lg md:text-xl text-zinc-400 mb-12 leading-relaxed">
            An exclusive platform connecting serious renters with verified landlords. No agents. No endless scrolling. Just your next home.
        </p>

        <a href="/explore" class="animate-fade-in-up [animation-delay:800ms] group bg-white text-black px-10 py-5 rounded-full text-lg font-bold hover:scale-105 active:scale-95 transition-all flex items-center gap-3 shadow-[0_20px_50px_rgba(255,255,255,0.1)]">
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
</html>`;

const getExploreHTML = () => `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Explore Sanctuaries - Nyumba</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        @keyframes fadeInUp {
            from { opacity: 0; transform: translateY(20px); }
            to { opacity: 1; transform: translateY(0); }
        }
        .animate-card {
            opacity: 0;
            animation: fadeInUp 0.6s cubic-bezier(0.16, 1, 0.3, 1) forwards;
        }
        .bg-mesh {
            background-color: #09090b;
            background-image: 
                radial-gradient(at 0% 0%, rgba(30, 58, 138, 0.15) 0px, transparent 50%),
                radial-gradient(at 100% 100%, rgba(20, 184, 166, 0.1) 0px, transparent 50%);
        }
    </style>
</head>
<body class="bg-mesh text-white font-sans min-h-screen">
    <nav class="p-6 flex justify-between items-center max-w-7xl mx-auto border-b border-white/5 backdrop-blur-md sticky top-0 z-50">
        <h1 class="text-2xl font-black tracking-tighter">Nyumba.</h1>
        <div class="flex items-center gap-6">
            <span class="text-zinc-400 text-sm font-medium">Welcome, User</span>
            <a href="/" class="text-zinc-400 hover:text-white transition-colors text-sm font-medium">Home</a>
        </div>
    </nav>
    <main class="max-w-7xl mx-auto px-6 py-12">
        <div class="flex flex-col md:flex-row justify-between items-start md:items-end mb-12 gap-6">
            <div>
                <h2 class="text-5xl md:text-7xl font-black tracking-tighter leading-none" id="explore-title">Available Sanctuaries</h2>
                <p class="text-zinc-500 mt-4 font-medium" id="filter-status">Showing all verified listings</p>
            </div>
            <a href="/explore" id="clear-filter" class="hidden text-sm text-zinc-400 hover:text-white underline font-medium">Clear filters</a>
        </div>

        <!-- Filters -->
        <div class="flex flex-wrap gap-6 mb-16 p-8 bg-white/5 backdrop-blur-2xl rounded-[2.5rem] border border-white/10 shadow-2xl">
            <div class="flex flex-col gap-3">
                <label class="text-[10px] font-black text-zinc-500 uppercase tracking-[0.2em]">Min Price (KSh)</label>
                <input type="number" id="min-price" placeholder="0" class="bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all w-44">
            </div>
            <div class="flex flex-col gap-3">
                <label class="text-[10px] font-black text-zinc-500 uppercase tracking-[0.2em]">Max Price (KSh)</label>
                <input type="number" id="max-price" placeholder="Any" class="bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all w-44">
            </div>
            <div class="flex flex-col gap-3">
                <label class="text-[10px] font-black text-zinc-500 uppercase tracking-[0.2em]">Sort By</label>
                <select id="sort-by" class="bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all w-56 text-zinc-400 appearance-none">
                    <option value="default">Default</option>
                    <option value="price-asc">Price: Low to High</option>
                    <option value="price-desc">Price: High to Low</option>
                    <option value="bedrooms-desc">Bedrooms: Most to Least</option>
                    <option value="bathrooms-desc">Bathrooms: Most to Least</option>
                </select>
            </div>
            <div class="flex items-end">
                <button id="apply-filters" class="bg-blue-600 text-white font-black px-10 py-4 rounded-2xl hover:bg-blue-500 hover:scale-105 active:scale-95 transition-all shadow-xl shadow-blue-600/20">
                    Apply Filters
                </button>
            </div>
        </div>
        <div id="houses-grid" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            <!-- Loaded via JS -->
        </div>
    </main>

    <!-- Floating Mobile Filter Button -->
    <div class="fixed bottom-8 left-1/2 -translate-x-1/2 md:hidden z-50">
        <button onclick="window.scrollTo({top: 0, behavior: 'smooth'})" class="bg-white text-black px-8 py-4 rounded-full font-black text-sm shadow-2xl flex items-center gap-2">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><path d="M3 6h18"/><path d="M7 12h10"/><path d="M10 18h4"/></svg>
            Filters
        </button>
    </div>
    <script>
        const urlParams = new URLSearchParams(window.location.search);
        const neighborhoodFilter = urlParams.get('neighborhood');
        let allHouses = [];

        if (neighborhoodFilter) {
            document.getElementById('explore-title').innerText = 'Sanctuaries in ' + neighborhoodFilter;
            document.getElementById('filter-status').innerText = 'Showing listings for ' + neighborhoodFilter;
            document.getElementById('clear-filter').classList.remove('hidden');
        }

        function renderHouses(houses) {
            const grid = document.getElementById('houses-grid');
            grid.innerHTML = '';

            if (houses.length === 0) {
                grid.innerHTML = '<div class="col-span-full py-20 text-center text-zinc-500">No sanctuaries found matching your criteria.</div>';
                return;
            }

            houses.forEach((house, index) => {
                const card = document.createElement('div');
                // Bento grid logic: every 4th item spans 2 columns on large screens
                const isBento = index % 4 === 0;
                card.className = \`bg-white/5 backdrop-blur-xl rounded-[2.5rem] overflow-hidden border border-white/10 hover:border-white/20 transition-all group animate-card \${isBento ? 'lg:col-span-2' : ''}\`;
                card.style.animationDelay = \`\${index * 100}ms\`;
                
                card.innerHTML = \`
                    <div class="relative overflow-hidden \${isBento ? 'aspect-[21/9]' : 'aspect-video'}">
                        <img src="\${house.image_urls[0]}" alt="\${house.building_name}" class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-700">
                        <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-transparent to-transparent opacity-60"></div>
                        <div class="absolute bottom-6 left-6">
                             <span class="bg-blue-600 text-white px-4 py-1.5 rounded-full text-[10px] font-black tracking-widest uppercase flex items-center gap-1.5 shadow-xl shadow-blue-600/40">
                                <svg xmlns="http://www.w3.org/2000/svg" width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"><path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/><circle cx="12" cy="10" r="3"/></svg>
                                \${house.location}
                            </span>
                        </div>
                    </div>
                    <div class="p-8">
                        <div class="flex justify-between items-start mb-6">
                            <div>
                                <h3 class="text-3xl font-black tracking-tighter">\${house.building_name}</h3>
                                <p class="text-zinc-500 text-sm font-medium mt-1">Verified Sanctuary</p>
                            </div>
                            <div class="text-right">
                                <p class="text-[10px] font-black text-zinc-500 uppercase tracking-widest mb-1">Monthly</p>
                                <span class="text-2xl font-black text-white">KSh \${house.price.toLocaleString()}</span>
                            </div>
                        </div>
                        <div class="flex gap-6 pt-6 border-t border-white/5">
                            <span class="text-sm text-zinc-400 font-bold flex items-center gap-2 bg-white/5 px-4 py-2 rounded-full">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-blue-400"><path d="M2 4v16"/><path d="M2 8h18a2 2 0 0 1 2 2v10"/><path d="M2 17h20"/><path d="M6 8v9"/></svg>
                                \${house.bedrooms} Bed
                            </span>
                            <span class="text-sm text-zinc-400 font-bold flex items-center gap-2 bg-white/5 px-4 py-2 rounded-full">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-emerald-400"><path d="M9 6 6.5 3.5a1.5 1.5 0 0 0-2.12 0 1.5 1.5 0 0 0 0 2.12L7 8"/><rect width="16" height="12" x="2" y="8" rx="2"/><path d="M7 12h.01"/><path d="M17 12h.01"/><path d="M12 12h.01"/></svg>
                                \${house.bathrooms} Bath
                            </span>
                        </div>
                    </div>
                \`;
                grid.appendChild(card);
            });
        }

        function applyFilters() {
            const minPrice = parseFloat(document.getElementById('min-price').value) || 0;
            const maxPrice = parseFloat(document.getElementById('max-price').value) || Infinity;
            const sortBy = document.getElementById('sort-by').value;

            let filtered = [...allHouses];

            if (neighborhoodFilter) {
                filtered = filtered.filter(h => h.location.toLowerCase().includes(neighborhoodFilter.toLowerCase()));
            }

            filtered = filtered.filter(h => h.price >= minPrice && h.price <= maxPrice);

            // Sorting logic
            if (sortBy === 'price-asc') {
                filtered.sort((a, b) => a.price - b.price);
            } else if (sortBy === 'price-desc') {
                filtered.sort((a, b) => b.price - a.price);
            } else if (sortBy === 'bedrooms-desc') {
                filtered.sort((a, b) => b.bedrooms - a.bedrooms);
            } else if (sortBy === 'bathrooms-desc') {
                filtered.sort((a, b) => b.bathrooms - a.bathrooms);
            }

            renderHouses(filtered);
            document.getElementById('clear-filter').classList.remove('hidden');
        }

        document.getElementById('apply-filters').addEventListener('click', applyFilters);

        fetch('/houses')
            .then(res => res.json())
            .then(houses => {
                allHouses = houses;
                let filtered = houses;
                if (neighborhoodFilter) {
                    filtered = houses.filter(h => h.location.toLowerCase().includes(neighborhoodFilter.toLowerCase()));
                }
                renderHouses(filtered);
            });
    </script>
</body>
</html>`;

const getLandlordHTML = () => `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Landlord Dashboard - Nyumba</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Inter', sans-serif; }
        .bg-mesh {
            background-color: #020617;
            background-image: 
                radial-gradient(at 0% 0%, rgba(30, 58, 138, 0.3) 0px, transparent 50%),
                radial-gradient(at 100% 0%, rgba(20, 184, 166, 0.2) 0px, transparent 50%),
                radial-gradient(at 50% 100%, rgba(88, 28, 135, 0.2) 0px, transparent 50%);
            background-size: 200% 200%;
            animation: meshGradient 20s ease infinite;
        }
        @keyframes meshGradient {
            0% { background-position: 0% 0%; }
            50% { background-position: 100% 100%; }
            100% { background-position: 0% 0%; }
        }
        ::-webkit-scrollbar {
            width: 6px;
        }
        ::-webkit-scrollbar-track {
            background: transparent;
        }
        ::-webkit-scrollbar-thumb {
            background: rgba(255, 255, 255, 0.1);
            border-radius: 10px;
        }
        ::-webkit-scrollbar-thumb:hover {
            background: rgba(255, 255, 255, 0.2);
        }
    </style>
</head>
<body class="bg-zinc-950 text-white min-h-screen flex">
    <!-- Sidebar -->
    <aside class="w-85 border-r border-white/5 flex flex-col h-screen sticky top-0 bg-black/40 backdrop-blur-3xl z-50">
        <div class="p-10">
            <h1 class="text-4xl font-black tracking-tighter mb-2">Nyumba.</h1>
            <p class="text-[10px] font-black tracking-[0.4em] text-zinc-500 uppercase">Landlord Portal</p>
        </div>

        <div class="flex-1 overflow-y-auto px-8 pb-10">
            <div class="bg-white/5 border border-white/10 rounded-[2rem] p-8 shadow-2xl">
                <h3 class="text-lg font-black tracking-tight mb-6">New Listing</h3>
                <form action="/add-house" method="POST" class="space-y-5">
                    <div class="space-y-2">
                        <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Building Name</label>
                        <input type="text" name="building_name" placeholder="e.g. Azure Heights" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all" required>
                    </div>
                    
                    <div class="space-y-2">
                        <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Location</label>
                        <input type="text" name="location" placeholder="e.g. Westlands" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all" required>
                    </div>
                    
                    <div class="space-y-2">
                        <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Monthly Rent (KSh)</label>
                        <input type="number" name="price" placeholder="75000" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all" required>
                    </div>

                    <div class="grid grid-cols-2 gap-4">
                        <div class="space-y-2">
                            <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Beds</label>
                            <input type="number" name="bedrooms" placeholder="2" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all">
                        </div>
                        <div class="space-y-2">
                            <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Baths</label>
                            <input type="number" name="bathrooms" placeholder="2" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all">
                        </div>
                    </div>

                    <button type="submit" class="w-full bg-white text-black font-black py-5 rounded-2xl hover:bg-zinc-200 hover:scale-[1.02] active:scale-95 transition-all shadow-2xl shadow-white/5 mt-4">
                        Publish Sanctuary
                    </button>
                </form>
            </div>
        </div>

        <div class="p-8 border-t border-white/5 flex items-center justify-between bg-black/20">
            <div class="flex items-center gap-4">
                <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-blue-500 to-indigo-600 flex items-center justify-center font-black text-white shadow-xl shadow-blue-500/20">
                    M
                </div>
                <div>
                    <p class="text-sm font-black tracking-tight">Musa Landlord</p>
                    <a href="/" class="text-[10px] font-bold text-zinc-500 hover:text-white transition-colors uppercase tracking-widest">Sign Out</a>
                </div>
            </div>
        </div>
    </aside>

    <!-- Main Content -->
    <main class="flex-1 bg-mesh p-16 overflow-y-auto">
        <div class="max-w-6xl">
            <div class="flex justify-between items-center mb-16">
                <div>
                    <h2 class="text-7xl font-black tracking-tighter leading-none mb-4">
                        Your <span class="text-blue-400">Portfolio</span>
                    </h2>
                    <p class="text-zinc-400 text-xl font-medium">Manage your verified sanctuaries and connect with serious renters.</p>
                </div>
                <div class="flex gap-4">
                    <div class="bg-white/5 backdrop-blur-xl border border-white/10 p-6 rounded-3xl text-center min-w-[140px]">
                        <p class="text-[10px] font-black text-zinc-500 uppercase tracking-widest mb-2">Active</p>
                        <p class="text-4xl font-black">12</p>
                    </div>
                    <div class="bg-white/5 backdrop-blur-xl border border-white/10 p-6 rounded-3xl text-center min-w-[140px]">
                        <p class="text-[10px] font-black text-zinc-500 uppercase tracking-widest mb-2">Views</p>
                        <p class="text-4xl font-black text-emerald-400">2.4k</p>
                    </div>
                </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-10">
                <!-- Preview Card -->
                <div class="bg-white/5 backdrop-blur-2xl border border-white/10 rounded-[3rem] p-8 group hover:border-white/20 transition-all shadow-2xl">
                    <div class="relative aspect-[16/10] rounded-[2rem] overflow-hidden mb-8">
                        <img src="https://picsum.photos/seed/apartment/800/600" alt="Preview" class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-700">
                        <div class="absolute top-6 right-6 bg-black/60 backdrop-blur-xl px-5 py-2.5 rounded-2xl text-sm font-black border border-white/10 shadow-2xl">
                            KES 75,000
                        </div>
                        <div class="absolute bottom-6 left-6 flex gap-2">
                            <span class="bg-blue-600 text-white px-5 py-2 rounded-full text-[10px] font-black tracking-widest uppercase flex items-center gap-2 shadow-xl shadow-blue-600/40">
                                <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"><path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/><circle cx="12" cy="10" r="3"/></svg>
                                WESTLANDS
                            </span>
                        </div>
                    </div>
                    
                    <div class="flex justify-between items-end mb-10">
                        <div>
                            <h3 class="text-4xl font-black tracking-tighter mb-2">Azure Heights</h3>
                            <p class="text-zinc-500 font-medium">Verified Listing • 24 Days Active</p>
                        </div>
                        <div class="flex gap-2">
                            <button class="p-4 bg-white/5 rounded-2xl hover:bg-white/10 transition-colors">
                                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5Z"/><path d="m15 5 4 4"/></svg>
                            </button>
                        </div>
                    </div>
                    
                    <div class="pt-10 border-t border-white/5 flex gap-4">
                        <button class="flex-1 bg-white text-black font-black py-5 rounded-2xl hover:bg-zinc-200 transition-all">
                            View Analytics
                        </button>
                        <button class="px-6 bg-rose-500/10 text-rose-500 font-black rounded-2xl hover:bg-rose-500/20 transition-all">
                            Unpublish
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </main>
</body>
</html>`;

// Routes
app.get("/", (req, res) => res.send(getLandingHTML()));
app.get("/explore", (req, res) => res.send(getExploreHTML()));
app.get("/landlord", (req, res) => res.send(getLandlordHTML()));
app.get("/houses", (req, res) => res.json(houses));

app.post("/add-house", (req, res) => {
    const newHouse = {
        id: houses.length + 1,
        building_name: req.body.building_name,
        location: req.body.location,
        price: parseFloat(req.body.price) || 0,
        image_urls: ["/uploads/default.jpg"],
        bedrooms: parseInt(req.body.bedrooms) || 1,
        bathrooms: parseInt(req.body.bathrooms) || 1,
    };
    houses.push(newHouse);
    fs.writeFileSync("houses.json", JSON.stringify(houses, null, 2));
    res.redirect("/explore");
});

app.listen(PORT, "0.0.0.0", () => {
    console.log("Server running on port " + PORT);
});
