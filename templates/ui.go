package templates

import (
	"encoding/json"
	"fmt"
	"nyumba/models"
)

// House struct for template-specific use
type House struct {
	ID           int      `json:"id"`
	BuildingName string   `json:"building_name"`
	Location     string   `json:"location"`
	Price        float64  `json:"price"`
	Deposit      float64  `json:"deposit"`
	Type         string   `json:"type"`
	ImageURLs    []string `json:"image_urls"`
	IsPaid       bool     `json:"is_paid"`
	Bedrooms     int      `json:"bedrooms"`
	Bathrooms    int      `json:"bathrooms"`
	Phone        string   `json:"landlord_phone"`
	Description  string   `json:"description"`
}

func getHeader() string {
	return `
    <header class="fixed top-0 left-0 w-full z-50 backdrop-blur-xl bg-black/40 border-b border-white/5">
        <nav class="max-w-7xl mx-auto px-6 py-4 flex justify-between items-center">
            <div class="flex items-center gap-12">
                <a href="/" class="text-2xl font-black tracking-tighter">Nyumba.</a>
                <div class="hidden md:flex items-center gap-8 text-sm font-bold text-zinc-400">
                    <a href="/" class="hover:text-white transition-colors">Home</a>
                    <a href="/explore" class="hover:text-white transition-colors">Listings</a>
                    <a href="/about" class="hover:text-white transition-colors">About</a>
                    <a href="/contact" class="hover:text-white transition-colors">Contact</a>
                </div>
            </div>
            <div class="flex items-center gap-4">
                <a href="/login" class="hidden md:block text-sm font-bold text-zinc-400 hover:text-white transition-colors">Login</a>
                <a href="/signup" class="bg-white text-black px-6 py-2.5 rounded-full text-sm font-bold hover:scale-105 transition-all">Sign Up</a>
                <button id="mobile-menu-btn" class="md:hidden p-2 text-zinc-400">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><line x1="4" x2="20" y1="12" y2="12"/><line x1="4" x2="20" y1="6" y2="6"/><line x1="4" x2="20" y1="18" y2="18"/></svg>
                </button>
            </div>
        </nav>
        <div id="mobile-menu" class="hidden md:hidden bg-zinc-950 border-b border-white/5 p-6 flex flex-col gap-4">
            <a href="/" class="text-lg font-bold">Home</a>
            <a href="/explore" class="text-lg font-bold">Listings</a>
            <a href="/about" class="text-lg font-bold">About</a>
            <a href="/contact" class="text-lg font-bold">Contact</a>
            <hr class="border-white/5">
            <a href="/login" class="text-lg font-bold">Login</a>
        </div>
    </header>
    <script>
        document.getElementById('mobile-menu-btn').addEventListener('click', () => {
            document.getElementById('mobile-menu').classList.toggle('hidden');
        });
    </script>`
}
func GetLandingHTML(featured []models.House) string {
	featuredJSON, _ := json.Marshal(featured)
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nyumba - Find Your Sanctuary</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800;900&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Inter', sans-serif; background-color: #09090b; color: white; }
        @keyframes marquee { 0%% { transform: translateX(0); } 100%% { transform: translateX(-50%%); } }
        .marquee-container { overflow: hidden; white-space: nowrap; position: relative; }
        .marquee-content { display: inline-block; animation: marquee 40s linear infinite; }
        .bg-mesh {
            background-image: 
                radial-gradient(at 0%% 0%%, rgba(30, 58, 138, 0.3) 0px, transparent 50%%),
                radial-gradient(at 100%% 0%%, rgba(20, 184, 166, 0.2) 0px, transparent 50%%);
        }
    </style>
</head>
<body class="bg-mesh">
    %s

    <section class="pt-48 pb-32 flex flex-col items-center text-center px-6">
        <div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full border border-white/10 bg-white/5 text-[10px] font-black tracking-widest uppercase mb-12">
            <span class="w-2 h-2 rounded-full bg-emerald-500 shadow-[0_0_8px_rgba(16,185,129,0.6)]"></span>
            Verified Listings Only
        </div>
        <h1 class="text-6xl md:text-9xl font-black tracking-tighter mb-8 leading-[0.85]">
            Find Your <br> <span class="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-cyan-400 to-emerald-400">Sanctuary.</span>
        </h1>
        <p class="max-w-xl text-lg md:text-xl text-zinc-400 mb-12 leading-relaxed font-medium">
            An exclusive platform connecting serious renters with verified landlords. No agents. No endless scrolling. Just your next home.
        </p>
        <div class="flex flex-col md:flex-row gap-4">
            <a href="/explore" class="group bg-white text-black px-10 py-5 rounded-full text-lg font-black hover:scale-105 transition-all flex items-center gap-3 shadow-[0_20px_50px_rgba(255,255,255,0.1)]">
                Start Your Search
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" class="group-hover:translate-x-1 transition-transform"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
            </a>
            <a href="/landlord" class="px-10 py-5 rounded-full text-lg font-black border border-white/10 hover:bg-white/5 transition-all">List Your Property</a>
        </div>
    </section>

    <section class="py-32 bg-white/5 border-y border-white/5">
        <div class="max-w-7xl mx-auto px-6">
            <div class="flex justify-between items-end mb-16">
                <div>
                    <h3 class="text-[10px] uppercase tracking-[0.3em] text-zinc-500 font-black mb-4">Featured Sanctuaries</h3>
                    <h2 class="text-5xl font-black tracking-tighter">Hand-picked for you</h2>
                </div>
                <a href="/explore" class="text-zinc-400 hover:text-white font-bold underline">View all listings</a>
            </div>
            <div id="featured-grid" class="grid grid-cols-1 md:grid-cols-3 gap-8"></div>
        </div>
    </section>

    <section class="pb-32">
        <div class="max-w-7xl mx-auto px-6 mb-12 text-center">
            <h3 class="text-[10px] uppercase tracking-[0.3em] text-zinc-500 font-black">Popular Neighborhoods</h3>
        </div>
        <div class="marquee-container">
            <div class="marquee-content flex gap-4 px-4">
                <div class="flex gap-4">
                    <a href="/explore?neighborhood=Westlands" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm">Westlands</a>
                    <a href="/explore?neighborhood=Kilimani" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm">Kilimani</a>
                    <a href="/explore?neighborhood=Karen" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm">Karen</a>
                    <a href="/explore?neighborhood=Thika" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm">Thika</a>
                </div>
            </div>
        </div>
    </section>

    %s

    <script>
        const featured = %s;
        const grid = document.getElementById('featured-grid');
        featured.forEach(house => {
            const card = document.createElement('div');
            card.className = 'bg-white/5 backdrop-blur-xl rounded-[2.5rem] overflow-hidden border border-white/10 hover:border-white/20 transition-all group';
            card.innerHTML = %s
                <div class="aspect-video overflow-hidden relative">
                    <img src="${house.image_urls[0]}" alt="${house.building_name}" class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-700">
                </div>
                <div class="p-8">
                    <h3 class="text-2xl font-black tracking-tighter mb-2">${house.building_name}</h3>
                    <div class="flex justify-between items-center pt-6 border-t border-white/5">
                        <span class="text-xl font-black">KSh ${house.price.toLocaleString()}</span>
                    </div>
                </div>
            %s;
            grid.appendChild(card);
        });
    </script>
</body>
</html>`, getHeader(), getFooter(), featuredJSON, "`", "`", "`", "`")
}
func getFooter() string {
	return `
    <footer class="bg-zinc-950 border-t border-white/5 py-20">
        <div class="max-w-7xl mx-auto px-6 grid grid-cols-1 md:grid-cols-4 gap-12">
            <div class="col-span-1 md:col-span-2">
                <h2 class="text-3xl font-black tracking-tighter mb-6">Nyumba.</h2>
                <p class="text-zinc-500 max-w-sm leading-relaxed mb-8">Kenya's premier sanctuary discovery platform.</p>
            </div>
            <div>
                <h4 class="text-sm font-black uppercase tracking-widest text-zinc-500 mb-6">Platform</h4>
                <ul class="space-y-4 text-zinc-400 font-bold">
                    <li><a href="/explore" class="hover:text-white transition-colors">Listings</a></li>
                    <li><a href="/landlord" class="hover:text-white transition-colors">For Landlords</a></li>
                </ul>
            </div>
        </div>
        <div class="max-w-7xl mx-auto px-6 mt-20 pt-8 border-t border-white/5 text-center text-zinc-600 text-sm font-bold">
            &copy; 2026 Nyumba Technologies. All rights reserved.
        </div>
    </footer>`
}

func GetExploreHTML() string {
	return fmt.Sprintf(`<!DOCTYPE html><html>%s<body class="bg-black text-white pt-32 px-6"><h1 class="text-5xl font-black">Explore Sanctuaries</h1>%s</body></html>`, getHeader(), getFooter())
}

func GetLandlordHTML() string {
	return fmt.Sprintf(`<!DOCTYPE html><html>%s<body class="bg-black text-white pt-32 px-6"><h1 class="text-5xl font-black">Landlord Portal</h1>%s</body></html>`, getHeader(), getFooter())
}

func GetAuthHTML(mode string) string {
	return fmt.Sprintf(`<!DOCTYPE html><html><body class="bg-black text-white flex items-center justify-center min-h-screen"><h1 class="text-4xl font-black">%s</h1></body></html>`, mode)
}

func GetStaticHTML(title, content string) string {
	return fmt.Sprintf(`<!DOCTYPE html><html>%s<body class="bg-black text-white pt-32 px-6"><h1 class="text-5xl font-black">%s</h1><p class="mt-8 text-zinc-400">%s</p>%s</body></html>`, getHeader(), title, content, getFooter())
}

func GetHTML(currentUsername string) string {
	return GetExploreHTML()
}
