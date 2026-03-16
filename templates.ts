
export const getHeader = () => `
    <header class="fixed top-0 left-0 w-full z-50 backdrop-blur-xl bg-black/60 border-b border-white/5">
        <nav class="max-w-7xl mx-auto px-6 py-4 flex justify-between items-center">
            <div class="flex items-center gap-12">
                <a href="/" class="text-2xl font-black tracking-tighter text-white flex items-center gap-2">
                    <div class="w-8 h-8 bg-blue-600 rounded-lg flex items-center justify-center">
                        <span class="text-white text-sm">N</span>
                    </div>
                    Nyumba.
                </a>
                <div class="hidden md:flex items-center gap-8 text-sm font-bold text-zinc-400">
                    <a href="/" class="hover:text-white transition-colors">Home</a>
                    <a href="/explore" class="hover:text-white transition-colors">Explore</a>
                    <a href="/about" class="hover:text-white transition-colors">About</a>
                </div>
            </div>
            <div class="flex items-center gap-4">
                <button onclick="toggleAuthModal()" class="hidden md:block text-sm font-bold text-zinc-400 hover:text-white transition-colors">Sign In</button>
                <a href="/landlord" class="bg-blue-600 text-white px-6 py-2.5 rounded-full text-sm font-bold hover:bg-blue-500 transition-all">List Property</a>
                <button id="mobile-menu-btn" class="md:hidden p-2 text-zinc-400">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><line x1="4" x2="20" y1="12" y2="12"/><line x1="4" x2="20" y1="6" y2="6"/><line x1="4" x2="20" y1="18" y2="18"/></svg>
                </button>
            </div>
        </nav>
        <!-- Mobile Menu -->
        <div id="mobile-menu" class="hidden md:hidden bg-zinc-950 border-b border-white/5 p-6 flex flex-col gap-4">
            <a href="/" class="text-lg font-bold text-white">Home</a>
            <a href="/explore" class="text-lg font-bold text-white">Explore</a>
            <a href="/about" class="text-lg font-bold text-white">About</a>
            <hr class="border-white/5">
            <button onclick="toggleAuthModal()" class="text-left text-lg font-bold text-white">Sign In</button>
        </div>
    </header>
    
    <!-- Auth Modal -->
    <div id="auth-modal" class="fixed inset-0 z-[100] hidden flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/80 backdrop-blur-sm" onclick="toggleAuthModal()"></div>
        <div class="relative w-full max-w-md bg-zinc-900 border border-white/10 rounded-[2.5rem] p-10 shadow-2xl overflow-hidden">
            <div class="absolute top-0 left-0 w-full h-1 bg-gradient-to-r from-blue-500 via-cyan-500 to-emerald-500"></div>
            <button onclick="toggleAuthModal()" class="absolute top-6 right-6 text-zinc-500 hover:text-white transition-colors">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
            </button>
            
            <div id="auth-content">
                <div class="text-center mb-10">
                    <h2 class="text-3xl font-black tracking-tighter text-white mb-2">Welcome Back</h2>
                    <p class="text-zinc-500 font-bold text-sm">Sign in to manage your sanctuaries</p>
                </div>
                
                <div class="space-y-4">
                    <button class="w-full flex items-center justify-center gap-3 bg-white text-black font-bold py-4 rounded-2xl hover:bg-zinc-200 transition-all">
                        <svg class="w-5 h-5" viewBox="0 0 24 24"><path fill="currentColor" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/><path fill="currentColor" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/><path fill="currentColor" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l3.66-2.84z"/><path fill="currentColor" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/></svg>
                        Continue with Google
                    </button>
                    <div class="relative flex items-center py-4">
                        <div class="flex-grow border-t border-white/5"></div>
                        <span class="flex-shrink mx-4 text-zinc-600 text-xs font-black uppercase tracking-widest">or</span>
                        <div class="flex-grow border-t border-white/5"></div>
                    </div>
                    <div class="space-y-4">
                        <input type="email" placeholder="Email Address" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all">
                        <input type="password" placeholder="Password" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all">
                        <button class="w-full bg-blue-600 text-white font-black py-4 rounded-2xl hover:bg-blue-500 transition-all shadow-xl shadow-blue-600/20">
                            Sign In
                        </button>
                    </div>
                </div>
                
                <p class="mt-8 text-center text-zinc-500 font-bold text-sm">
                    Don't have an account? <button onclick="toggleAuthMode()" class="text-white underline">Sign Up</button>
                </p>
            </div>
        </div>
    </div>

    <script>
        document.getElementById('mobile-menu-btn')?.addEventListener('click', () => {
            document.getElementById('mobile-menu').classList.toggle('hidden');
        });
        
        function toggleAuthModal() {
            const modal = document.getElementById('auth-modal');
            modal.classList.toggle('hidden');
            if (!modal.classList.contains('hidden')) {
                document.body.style.overflow = 'hidden';
            } else {
                document.body.style.overflow = 'auto';
            }
        }
        
        function toggleAuthMode() {
            // Simple toggle for demo
            const title = document.querySelector('#auth-content h2');
            const p = document.querySelector('#auth-content p');
            const btn = document.querySelector('#auth-content button:last-of-type');
            const footer = document.querySelector('#auth-content p:last-of-type');
            
            if (title.innerText === 'Welcome Back') {
                title.innerText = 'Create Account';
                p.innerText = 'Join the sanctuary network';
                btn.innerText = 'Sign Up';
                footer.innerHTML = 'Already have an account? <button onclick="toggleAuthMode()" class="text-white underline">Sign In</button>';
            } else {
                title.innerText = 'Welcome Back';
                p.innerText = 'Sign in to manage your sanctuaries';
                btn.innerText = 'Sign In';
                footer.innerHTML = 'Don\\'t have an account? <button onclick="toggleAuthMode()" class="text-white underline">Sign Up</button>';
            }
        }
    </script>
`;

export const getFooter = () => `
    <footer class="bg-zinc-950 border-t border-white/5 py-20">
        <div class="max-w-7xl mx-auto px-6 grid grid-cols-1 md:grid-cols-4 gap-12">
            <div class="col-span-1 md:col-span-2">
                <h2 class="text-3xl font-black tracking-tighter mb-6 text-white">Nyumba.</h2>
                <p class="text-zinc-500 max-w-sm leading-relaxed mb-8">Kenya's premier sanctuary discovery platform. Connecting serious renters with verified landlords directly.</p>
                <div class="flex gap-4">
                    <a href="#" class="p-3 bg-white/5 rounded-2xl hover:bg-white/10 transition-colors text-white"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"/></svg></a>
                    <a href="#" class="p-3 bg-white/5 rounded-2xl hover:bg-white/10 transition-colors text-white"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="20" height="20" x="2" y="2" rx="5" ry="5"/><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"/><line x1="17.5" x2="17.51" y1="6.5" y2="6.5"/></svg></a>
                </div>
            </div>
            <div>
                <h4 class="text-sm font-black uppercase tracking-widest text-zinc-500 mb-6">Platform</h4>
                <ul class="space-y-4 text-zinc-400 font-bold">
                    <li><a href="/explore" class="hover:text-white transition-colors">Listings</a></li>
                    <li><a href="/landlord" class="hover:text-white transition-colors">For Landlords</a></li>
                    <li><a href="/about" class="hover:text-white transition-colors">How it Works</a></li>
                </ul>
            </div>
            <div>
                <h4 class="text-sm font-black uppercase tracking-widest text-zinc-500 mb-6">Support</h4>
                <ul class="space-y-4 text-zinc-400 font-bold">
                    <li><a href="/contact" class="hover:text-white transition-colors">Contact Us</a></li>
                    <li><a href="#" class="hover:text-white transition-colors">Privacy Policy</a></li>
                    <li><a href="#" class="hover:text-white transition-colors">Terms of Service</a></li>
                </ul>
            </div>
        </div>
        <div class="max-w-7xl mx-auto px-6 mt-20 pt-8 border-t border-white/5 text-center text-zinc-600 text-sm font-bold">
            &copy; 2026 Nyumba Technologies. All rights reserved.
        </div>
    </footer>
`;

export const getLandingHTML = (featuredHouses: any[] = []) => `
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
        @keyframes marquee { 0% { transform: translateX(0); } 100% { transform: translateX(-50%); } }
        .marquee-container { overflow: hidden; white-space: nowrap; position: relative; }
        .marquee-content { display: inline-block; animation: marquee 40s linear infinite; }
        .bg-mesh {
            background-image: 
                radial-gradient(at 0% 0%, rgba(30, 58, 138, 0.3) 0px, transparent 50%),
                radial-gradient(at 100% 0%, rgba(20, 184, 166, 0.2) 0px, transparent 50%);
        }
    </style>
</head>
<body class="bg-mesh">
    ${getHeader()}

    <!-- Hero -->
    <section class="pt-48 pb-32 flex flex-col items-center text-center px-6">
        <div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full border border-white/10 bg-white/5 text-[10px] font-black tracking-widest uppercase mb-12">
            <span class="w-2 h-2 rounded-full bg-emerald-500 shadow-[0_0_8px_rgba(16,185,129,0.6)]"></span>
            Verified Listings Only
        </div>
        <h1 class="text-6xl md:text-9xl font-black tracking-tighter mb-8 leading-[0.85]">
            Find Your <br> <span class="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-cyan-400 to-emerald-400">Sanctuary.</span>
        </h1>
        <p class="max-w-xl text-lg md:text-xl text-zinc-400 mb-12 leading-relaxed font-medium">
            An exclusive platform connecting serious renters with verified landlords directly. No agents. No endless scrolling. Just your next home.
        </p>
        
        <!-- Hero Search Bar -->
        <div class="w-full max-w-4xl bg-white/5 backdrop-blur-2xl border border-white/10 rounded-[2.5rem] p-4 md:p-6 shadow-2xl mb-12">
            <form action="/explore" method="GET" class="flex flex-col md:flex-row gap-4">
                <div class="flex-1 flex flex-col gap-2 text-left px-4">
                    <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest">Location</label>
                    <div class="relative">
                        <svg class="w-4 h-4 absolute left-0 top-1/2 -translate-y-1/2 text-blue-400" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/><circle cx="12" cy="10" r="3"/></svg>
                        <input type="text" name="search" placeholder="Nairobi, Westlands..." class="w-full bg-transparent border-none pl-6 py-2 text-sm font-bold text-white focus:outline-none placeholder:text-zinc-600">
                    </div>
                </div>
                <div class="w-full md:w-px h-px md:h-12 bg-white/10 self-center"></div>
                <div class="flex-1 flex flex-col gap-2 text-left px-4">
                    <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest">Price Range</label>
                    <select name="maxPrice" class="w-full bg-transparent border-none py-2 text-sm font-bold text-white focus:outline-none appearance-none cursor-pointer">
                        <option value="" class="bg-zinc-900">Any Price</option>
                        <option value="20000" class="bg-zinc-900">Under KSh 20k</option>
                        <option value="50000" class="bg-zinc-900">Under KSh 50k</option>
                        <option value="100000" class="bg-zinc-900">Under KSh 100k</option>
                    </select>
                </div>
                <div class="w-full md:w-px h-px md:h-12 bg-white/10 self-center"></div>
                <div class="flex-1 flex flex-col gap-2 text-left px-4">
                    <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest">Property Type</label>
                    <select name="type" class="w-full bg-transparent border-none py-2 text-sm font-bold text-white focus:outline-none appearance-none cursor-pointer">
                        <option value="" class="bg-zinc-900">All Types</option>
                        <option value="Apartment" class="bg-zinc-900">Apartment</option>
                        <option value="Studio" class="bg-zinc-900">Studio</option>
                        <option value="House" class="bg-zinc-900">House</option>
                    </select>
                </div>
                <button type="submit" class="bg-blue-600 text-white px-8 py-4 rounded-2xl font-black hover:bg-blue-500 transition-all flex items-center justify-center gap-2 shadow-xl shadow-blue-600/20">
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
                    Search
                </button>
            </form>
        </div>
    </section>

    <!-- Featured Listings -->
    <section class="py-32 bg-white/5 border-y border-white/5">
        <div class="max-w-7xl mx-auto px-6">
            <div class="flex justify-between items-end mb-16">
                <div>
                    <h3 class="text-[10px] uppercase tracking-[0.3em] text-zinc-500 font-black mb-4">Featured Sanctuaries</h3>
                    <h2 class="text-5xl font-black tracking-tighter">Hand-picked for you</h2>
                </div>
                <a href="/explore" class="text-zinc-400 hover:text-white font-bold underline">View all listings</a>
            </div>
            <div id="featured-grid" class="grid grid-cols-1 md:grid-cols-3 gap-8">
                <!-- JS Inject -->
            </div>
        </div>
    </section>

    <!-- How it Works -->
    <section class="py-32">
        <div class="max-w-7xl mx-auto px-6">
            <div class="text-center mb-20">
                <h3 class="text-[10px] uppercase tracking-[0.3em] text-zinc-500 font-black mb-4">The Process</h3>
                <h2 class="text-5xl font-black tracking-tighter">Simplified Discovery</h2>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-12">
                <div class="p-10 rounded-[3rem] bg-white/5 border border-white/10">
                    <div class="w-16 h-16 bg-blue-600 rounded-3xl flex items-center justify-center font-black text-2xl mb-8 shadow-xl shadow-blue-600/20">1</div>
                    <h4 class="text-2xl font-black mb-4">Explore</h4>
                    <p class="text-zinc-400 leading-relaxed font-medium">Browse through our curated list of verified sanctuaries in Kenya's top neighborhoods.</p>
                </div>
                <div class="p-10 rounded-[3rem] bg-white/5 border border-white/10">
                    <div class="w-16 h-16 bg-indigo-600 rounded-3xl flex items-center justify-center font-black text-2xl mb-8 shadow-xl shadow-indigo-600/20">2</div>
                    <h4 class="text-2xl font-black mb-4">Unlock</h4>
                    <p class="text-zinc-400 leading-relaxed font-medium">Pay a small verification fee to unlock the landlord's direct contact information instantly.</p>
                </div>
                <div class="p-10 rounded-[3rem] bg-white/5 border border-white/10">
                    <div class="w-16 h-16 bg-emerald-600 rounded-3xl flex items-center justify-center font-black text-2xl mb-8 shadow-xl shadow-emerald-600/20">3</div>
                    <h4 class="text-2xl font-black mb-4">Connect</h4>
                    <p class="text-zinc-400 leading-relaxed font-medium">Talk directly to the landlord and schedule a viewing. No middleman, no agents, no stress.</p>
                </div>
            </div>
        </div>
    </section>

    <!-- Neighborhoods -->
    <section class="pb-32">
        <div class="max-w-7xl mx-auto px-6 mb-12 text-center">
            <h3 class="text-[10px] uppercase tracking-[0.3em] text-zinc-500 font-black">Popular Neighborhoods</h3>
        </div>
        <div class="marquee-container">
            <div class="marquee-content flex gap-4 px-4">
                <div class="flex gap-4">
                    <a href="/explore?search=Westlands" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm text-white">Westlands</a>
                    <a href="/explore?search=Kilimani" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm text-white">Kilimani</a>
                    <a href="/explore?search=Karen" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm text-white">Karen</a>
                    <a href="/explore?search=Thika" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm text-white">Thika</a>
                </div>
                <div class="flex gap-4">
                    <a href="/explore?search=Westlands" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm text-white">Westlands</a>
                    <a href="/explore?search=Kilimani" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm text-white">Kilimani</a>
                    <a href="/explore?search=Karen" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm text-white">Karen</a>
                    <a href="/explore?search=Thika" class="px-10 py-5 bg-white/5 border border-white/10 rounded-3xl hover:bg-white hover:text-black transition-all font-black text-lg backdrop-blur-sm text-white">Thika</a>
                </div>
            </div>
        </div>
    </section>

    <!-- FAQ -->
    <section class="py-32 bg-white/5 border-y border-white/5">
        <div class="max-w-3xl mx-auto px-6">
            <h2 class="text-5xl font-black tracking-tighter mb-16 text-center">Common Questions</h2>
            <div class="space-y-6">
                <details class="group bg-black/40 border border-white/10 rounded-3xl p-8">
                    <summary class="list-none flex justify-between items-center cursor-pointer font-black text-xl text-white">
                        How do I know the listings are real?
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" class="group-open:rotate-180 transition-transform"><path d="m6 9 6 6 6-6"/></svg>
                    </summary>
                    <p class="mt-6 text-zinc-400 font-medium leading-relaxed">Every landlord on Nyumba undergoes a strict identity and property verification process. We manually review every listing before it goes live.</p>
                </details>
                <details class="group bg-black/40 border border-white/10 rounded-3xl p-8">
                    <summary class="list-none flex justify-between items-center cursor-pointer font-black text-xl text-white">
                        Why is there a verification fee?
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" class="group-open:rotate-180 transition-transform"><path d="m6 9 6 6 6-6"/></svg>
                    </summary>
                    <p class="mt-6 text-zinc-400 font-medium leading-relaxed">The fee ensures only serious renters are connecting with landlords. This prevents spam and keeps the platform exclusive and high-quality for everyone.</p>
                </details>
                <details class="group bg-black/40 border border-white/10 rounded-3xl p-8">
                    <summary class="list-none flex justify-between items-center cursor-pointer font-black text-xl text-white">
                        Can I get a refund?
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" class="group-open:rotate-180 transition-transform"><path d="m6 9 6 6 6-6"/></svg>
                    </summary>
                    <p class="mt-6 text-zinc-400 font-medium leading-relaxed">Yes. If a property you unlocked is no longer available or the listing is inaccurate, we provide a full refund of the verification fee. Just reach out to our support within 24 hours.</p>
                </details>
            </div>
        </div>
    </section>

    ${getFooter()}

    <script>
        const featured = ${JSON.stringify(featuredHouses)};
        const grid = document.getElementById('featured-grid');
        featured.forEach(house => {
            const card = document.createElement('div');
            card.className = 'bg-white/5 backdrop-blur-xl rounded-[2.5rem] overflow-hidden border border-white/10 hover:border-white/20 transition-all group';
            card.innerHTML = \`
                <div class="aspect-video overflow-hidden relative">
                    <img src="\${house.image_urls[0]}" alt="\${house.building_name}" class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-700">
                    <div class="absolute top-6 left-6 bg-blue-600 text-white px-4 py-1.5 rounded-full text-[10px] font-black tracking-widest uppercase">
                        \${house.location}
                    </div>
                </div>
                <div class="p-8">
                    <h3 class="text-2xl font-black tracking-tighter mb-2 text-white">\${house.building_name}</h3>
                    <div class="flex justify-between items-center pt-6 border-t border-white/5">
                        <span class="text-xl font-black text-white">KSh \${house.price.toLocaleString()}</span>
                        <a href="/property/\${house.id}" class="text-sm font-black text-blue-400 hover:text-blue-300 transition-colors">View Details</a>
                    </div>
                </div>
            \`;
            grid.appendChild(card);
        });
    </script>
</body>
</html>
`;

export const getExploreHTML = () => `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Explore Sanctuaries - Nyumba</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800;900&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Inter', sans-serif; background-color: #09090b; color: white; }
    </style>
</head>
<body>
    ${getHeader()}
    <main class="max-w-7xl mx-auto px-6 pt-32 pb-20">
        <div class="flex flex-col md:flex-row justify-between items-start md:items-end mb-12 gap-6">
            <div>
                <h2 class="text-5xl md:text-7xl font-black tracking-tighter leading-none text-white" id="explore-title">Available Sanctuaries</h2>
                <p class="text-zinc-500 mt-4 font-bold" id="filter-status">Showing all verified listings</p>
            </div>
            <a href="/explore" id="clear-filter" class="hidden text-sm text-zinc-400 hover:text-white underline font-bold">Clear filters</a>
        </div>

        <!-- Search & Filters -->
        <div class="flex flex-wrap gap-6 mb-16 p-8 bg-white/5 backdrop-blur-2xl rounded-[2.5rem] border border-white/10 shadow-2xl">
            <div class="flex-1 min-w-[300px] flex flex-col gap-3">
                <label class="text-[10px] font-black text-zinc-500 uppercase tracking-[0.2em]">Search Location or Building</label>
                <input type="text" id="search-input" placeholder="e.g. Westlands, Azure Heights" class="bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all">
            </div>
            <div class="flex flex-col gap-3">
                <label class="text-[10px] font-black text-zinc-500 uppercase tracking-[0.2em]">Max Price (KSh)</label>
                <input type="number" id="max-price" placeholder="Any" class="bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all w-44">
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
    ${getFooter()}
    <script>
        const urlParams = new URLSearchParams(window.location.search);
        const searchFilter = urlParams.get('search');
        let allHouses = [];

        if (searchFilter) {
            document.getElementById('search-input').value = searchFilter;
            document.getElementById('explore-title').innerText = 'Sanctuaries in ' + searchFilter;
            document.getElementById('filter-status').innerText = 'Showing listings for ' + searchFilter;
            document.getElementById('clear-filter').classList.remove('hidden');
        }

        function renderHouses(houses) {
            const grid = document.getElementById('houses-grid');
            grid.innerHTML = '';
            if (houses.length === 0) {
                grid.innerHTML = '<div class="col-span-full py-20 text-center text-zinc-500 font-bold">No sanctuaries found matching your criteria.</div>';
                return;
            }
            houses.forEach((house, index) => {
                const card = document.createElement('div');
                card.className = 'bg-white/5 backdrop-blur-xl rounded-[2.5rem] overflow-hidden border border-white/10 hover:border-white/20 transition-all group';
                card.innerHTML = \`
                    <a href="/property/\${house.id}" class="relative aspect-video overflow-hidden block">
                        <img src="\${house.image_urls[0]}" alt="\${house.building_name}" loading="lazy" class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-700">
                        <div class="absolute bottom-6 left-6">
                             <span class="bg-blue-600 text-white px-4 py-1.5 rounded-full text-[10px] font-black tracking-widest uppercase flex items-center gap-1.5">
                                <svg xmlns="http://www.w3.org/2000/svg" width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"><path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/><circle cx="12" cy="10" r="3"/></svg>
                                \${house.location}
                            </span>
                        </div>
                    </a>
                    <div class="p-8">
                        <div class="flex justify-between items-start mb-6">
                            <div>
                                <a href="/property/\${house.id}"><h3 class="text-3xl font-black tracking-tighter text-white hover:text-blue-400 transition-colors">\${house.building_name}</h3></a>
                                <p class="text-zinc-500 text-sm font-bold mt-1">Verified Sanctuary</p>
                            </div>
                            <div class="text-right">
                                <p class="text-[10px] font-black text-zinc-500 uppercase tracking-widest mb-1">Monthly</p>
                                <span class="text-2xl font-black text-white">KSh \${house.price.toLocaleString()}</span>
                            </div>
                        </div>
                        <div class="flex gap-6 pt-6 border-t border-white/5 mb-8">
                            <span class="text-sm text-zinc-400 font-bold flex items-center gap-2 bg-white/5 px-4 py-2 rounded-full">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-blue-400"><path d="M2 4v16"/><path d="M2 8h18a2 2 0 0 1 2 2v10"/><path d="M2 17h20"/><path d="M6 8v9"/></svg>
                                \${house.bedrooms} Bed
                            </span>
                            <span class="text-sm text-zinc-400 font-bold flex items-center gap-2 bg-white/5 px-4 py-2 rounded-full">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-emerald-400"><path d="M9 6 6.5 3.5a1.5 1.5 0 0 0-2.12 0 1.5 1.5 0 0 0 0 2.12L7 8"/><rect width="16" height="12" x="2" y="8" rx="2"/><path d="M7 12h.01"/><path d="M17 12h.01"/><path d="M12 12h.01"/></svg>
                                \${house.bathrooms} Bath
                            </span>
                        </div>
                        <button onclick="triggerPayment(\${house.id})" class="w-full bg-white text-black font-black py-4 rounded-2xl hover:bg-zinc-200 transition-all flex items-center justify-center gap-2 group/btn">
                            Unlock for KES 1,000
                            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" class="group-hover/btn:translate-x-1 transition-transform"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
                        </button>
                    </div>
                \`;
                grid.appendChild(card);
            });
        }

        function triggerPayment(houseId) {
            const btn = event.currentTarget;
            const originalText = btn.innerHTML;
            btn.disabled = true;
            btn.innerHTML = 'Processing...';
            fetch('/api/trigger-payment', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ houseId })
            })
            .then(res => res.json())
            .then(data => {
                alert(data.message);
            })
            .finally(() => {
                btn.disabled = false;
                btn.innerHTML = originalText;
            });
        }

        document.getElementById('apply-filters').addEventListener('click', () => {
            const search = document.getElementById('search-input').value;
            const maxPrice = document.getElementById('max-price').value;
            let url = '/api/houses?';
            if (search) url += 'search=' + search + '&';
            fetch(url)
                .then(res => res.json())
                .then(houses => {
                    let filtered = houses;
                    if (maxPrice) filtered = houses.filter(h => h.price <= maxPrice);
                    renderHouses(filtered);
                });
        });

        fetch('/api/houses' + (searchFilter ? '?search=' + searchFilter : ''))
            .then(res => res.json())
            .then(houses => {
                allHouses = houses;
                renderHouses(houses);
            });
    </script>
</body>
</html>
`;

export const getLandlordHTML = () => `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Landlord Portal - Nyumba</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800;900&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Inter', sans-serif; background-color: #09090b; color: white; }
    </style>
</head>
<body>
    ${getHeader()}
    <main class="max-w-7xl mx-auto px-6 pt-32 pb-20 flex flex-col md:flex-row gap-12">
        <aside class="w-full md:w-80">
            <div class="bg-white/5 border border-white/10 rounded-[2.5rem] p-8 sticky top-32">
                <h3 class="text-2xl font-black tracking-tighter mb-6 text-white">New Listing</h3>
                <div id="form-feedback" class="hidden mb-6 p-4 rounded-2xl text-sm font-bold"></div>
                <form id="add-house-form" class="space-y-5">
                    <div class="space-y-2">
                        <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Building Name</label>
                        <input type="text" name="building_name" placeholder="e.g. Azure Heights" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all" required>
                    </div>
                    <div class="space-y-2">
                        <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Location</label>
                        <input type="text" name="location" placeholder="e.g. Westlands" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all" required>
                    </div>
                    <div class="space-y-2">
                        <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Monthly Rent (KSh)</label>
                        <input type="number" name="price" placeholder="75000" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all" required>
                    </div>
                    <div class="grid grid-cols-2 gap-4">
                        <div class="space-y-2">
                            <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Beds</label>
                            <input type="number" name="bedrooms" placeholder="2" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all">
                        </div>
                        <div class="space-y-2">
                            <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Baths</label>
                            <input type="number" name="bathrooms" placeholder="2" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all">
                        </div>
                    </div>
                    <button type="submit" class="w-full bg-white text-black font-black py-5 rounded-2xl hover:bg-zinc-200 transition-all shadow-2xl shadow-white/5 mt-4">
                        Publish Sanctuary
                    </button>
                </form>
            </div>
        </aside>
        <div class="flex-1">
            <h2 class="text-6xl font-black tracking-tighter mb-12 text-white">Your <span class="text-blue-400">Portfolio</span></h2>
            <div class="grid grid-cols-1 gap-8">
                <div class="bg-white/5 border border-white/10 rounded-[3rem] p-8 flex flex-col md:flex-row gap-8 items-center">
                    <div class="w-full md:w-64 aspect-video rounded-3xl overflow-hidden">
                        <img src="https://picsum.photos/seed/apartment/800/600" class="w-full h-full object-cover">
                    </div>
                    <div class="flex-1">
                        <h3 class="text-3xl font-black tracking-tighter mb-2 text-white">Azure Heights</h3>
                        <p class="text-zinc-500 font-bold mb-6">Westlands • Verified Sanctuary</p>
                        <div class="flex gap-4">
                            <button class="bg-white/5 px-6 py-3 rounded-2xl font-bold hover:bg-white/10 transition-all text-white">Edit Details</button>
                            <button class="bg-rose-500/10 text-rose-500 px-6 py-3 rounded-2xl font-bold hover:bg-rose-500/20 transition-all">Unpublish</button>
                        </div>
                    </div>
                    <div class="text-center md:text-right">
                        <p class="text-[10px] font-black text-zinc-500 uppercase tracking-widest mb-1">Views</p>
                        <p class="text-4xl font-black text-emerald-400">2.4k</p>
                    </div>
                </div>
            </div>
        </div>
    </main>
    ${getFooter()}
    <script>
        document.getElementById('add-house-form').addEventListener('submit', function(e) {
            e.preventDefault();
            const form = e.target;
            const feedback = document.getElementById('form-feedback');
            const btn = form.querySelector('button[type="submit"]');
            btn.disabled = true;
            btn.innerText = 'Publishing...';
            const formData = new FormData(form);
            const data = Object.fromEntries(formData.entries());
            fetch('/api/add-house', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json', 'Accept': 'application/json' },
                body: JSON.stringify(data)
            })
            .then(res => res.json())
            .then(data => {
                feedback.innerText = data.message || 'Property listed successfully!';
                feedback.className = 'mb-6 p-4 rounded-2xl text-sm font-bold bg-emerald-500/10 text-emerald-500 border border-emerald-500/20';
                feedback.classList.remove('hidden');
                form.reset();
            })
            .finally(() => {
                btn.disabled = false;
                btn.innerText = 'Publish Sanctuary';
            });
        });
    </script>
</body>
</html>
`;

export const getAuthHTML = (type: 'login' | 'signup') => {
    const mode = type === 'login' ? 'Login' : 'Sign Up';
    return `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>${mode} - Nyumba</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800;900&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Inter', sans-serif; background-color: #09090b; color: white; }
    </style>
</head>
<body class="flex items-center justify-center min-h-screen p-6">
    <div class="w-full max-w-md bg-white/5 border border-white/10 rounded-[3rem] p-12 shadow-2xl">
        <a href="/" class="text-2xl font-black tracking-tighter mb-12 block text-center text-white">Nyumba.</a>
        <h2 class="text-4xl font-black tracking-tighter mb-8 text-center text-white">${mode}</h2>
        <form class="space-y-6">
            <div class="space-y-2">
                <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Email Address</label>
                <input type="email" placeholder="name@example.com" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all">
            </div>
            <div class="space-y-2">
                <label class="text-[10px] font-black text-zinc-500 uppercase tracking-widest ml-2">Password</label>
                <input type="password" placeholder="••••••••" class="w-full bg-black/40 border border-white/10 rounded-2xl px-5 py-4 text-sm text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 transition-all">
            </div>
            <button type="button" onclick="alert('Auth simulation: Success!')" class="w-full bg-white text-black font-black py-5 rounded-2xl hover:bg-zinc-200 transition-all shadow-2xl shadow-white/5 mt-4">
                ${mode}
            </button>
        </form>
        <p class="mt-8 text-center text-zinc-500 font-bold text-sm">
            ${mode === "Login" ? `Don't have an account? <a href="/signup" class="text-white underline">Sign Up</a>` : `Already have an account? <a href="/login" class="text-white underline">Login</a>`}
        </p>
    </div>
</body>
</html>
`;
};

export const getStaticHTML = (title: string, content: string) => `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>${title} - Nyumba</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800;900&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Inter', sans-serif; background-color: #09090b; color: white; }
    </style>
</head>
<body>
    ${getHeader()}
    <main class="max-w-3xl mx-auto px-6 pt-48 pb-32">
        <h1 class="text-6xl md:text-8xl font-black tracking-tighter mb-12 text-white">${title}</h1>
        <div class="prose prose-invert prose-xl">
            <div class="text-zinc-400 font-medium leading-relaxed">${content}</div>
        </div>
    </main>
    ${getFooter()}
</body>
</html>
`;

export const getPropertyDetailHTML = (house: any) => `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>${house.building_name} - Nyumba</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800;900&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Inter', sans-serif; background-color: #09090b; color: white; }
    </style>
</head>
<body>
    ${getHeader()}
    <main class="pt-32 pb-20">
        <!-- Image Gallery -->
        <section class="max-w-7xl mx-auto px-6 mb-12">
            <div class="grid grid-cols-1 md:grid-cols-4 grid-rows-2 gap-4 h-[600px]">
                <div class="md:col-span-2 md:row-span-2 rounded-[2.5rem] overflow-hidden border border-white/10">
                    <img src="${house.image_urls[0]}" class="w-full h-full object-cover">
                </div>
                <div class="rounded-[2rem] overflow-hidden border border-white/10">
                    <img src="https://picsum.photos/seed/${house.id}1/800/600" class="w-full h-full object-cover">
                </div>
                <div class="rounded-[2rem] overflow-hidden border border-white/10">
                    <img src="https://picsum.photos/seed/${house.id}2/800/600" class="w-full h-full object-cover">
                </div>
                <div class="rounded-[2rem] overflow-hidden border border-white/10">
                    <img src="https://picsum.photos/seed/${house.id}3/800/600" class="w-full h-full object-cover">
                </div>
                <div class="rounded-[2rem] overflow-hidden border border-white/10 relative">
                    <img src="https://picsum.photos/seed/${house.id}4/800/600" class="w-full h-full object-cover">
                    <div class="absolute inset-0 bg-black/60 flex items-center justify-center">
                        <span class="text-white font-black text-xl">+8 More</span>
                    </div>
                </div>
            </div>
        </section>

        <section class="max-w-7xl mx-auto px-6 grid grid-cols-1 lg:grid-cols-3 gap-12">
            <!-- Left Column: Details -->
            <div class="lg:col-span-2">
                <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-8 gap-4">
                    <div>
                        <div class="flex gap-2 mb-4">
                            <span class="bg-emerald-500 text-black px-3 py-1 rounded-full text-[10px] font-black tracking-widest uppercase">Verified</span>
                            <span class="bg-blue-600 text-white px-3 py-1 rounded-full text-[10px] font-black tracking-widest uppercase">${house.type}</span>
                        </div>
                        <h1 class="text-5xl md:text-7xl font-black tracking-tighter text-white">${house.building_name}</h1>
                        <p class="text-xl text-zinc-500 font-bold mt-2 flex items-center gap-2">
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/><circle cx="12" cy="10" r="3"/></svg>
                            ${house.location}
                        </p>
                    </div>
                </div>

                <div class="grid grid-cols-2 md:grid-cols-4 gap-6 py-10 border-y border-white/10 mb-10">
                    <div class="flex flex-col gap-1">
                        <span class="text-zinc-500 text-[10px] font-black uppercase tracking-widest">Bedrooms</span>
                        <span class="text-2xl font-black text-white">${house.bedrooms} Rooms</span>
                    </div>
                    <div class="flex flex-col gap-1">
                        <span class="text-zinc-500 text-[10px] font-black uppercase tracking-widest">Bathrooms</span>
                        <span class="text-2xl font-black text-white">${house.bathrooms} Baths</span>
                    </div>
                    <div class="flex flex-col gap-1">
                        <span class="text-zinc-500 text-[10px] font-black uppercase tracking-widest">Square Feet</span>
                        <span class="text-2xl font-black text-white">1,200 sqft</span>
                    </div>
                    <div class="flex flex-col gap-1">
                        <span class="text-zinc-500 text-[10px] font-black uppercase tracking-widest">Availability</span>
                        <span class="text-2xl font-black text-emerald-400">Immediate</span>
                    </div>
                </div>

                <div class="mb-12">
                    <h3 class="text-2xl font-black tracking-tighter text-white mb-6">Description</h3>
                    <p class="text-zinc-400 text-lg leading-relaxed font-medium">
                        ${house.description || 'This stunning sanctuary offers modern living at its finest. Located in the heart of ' + house.location + ', it features high-end finishes, ample natural light, and breathtaking views.'}
                    </p>
                </div>

                <div>
                    <h3 class="text-2xl font-black tracking-tighter text-white mb-6">Amenities</h3>
                    <div class="grid grid-cols-2 md:grid-cols-3 gap-4">
                        ${['24/7 Security', 'High Speed Lift', 'Borehole Water', 'Backup Generator', 'Gym', 'Swimming Pool', 'Parking', 'WiFi Ready'].map(amenity => `
                            <div class="flex items-center gap-3 p-4 bg-white/5 border border-white/5 rounded-2xl">
                                <svg class="w-5 h-5 text-emerald-500" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><path d="M20 6 9 17l-5-5"/></svg>
                                <span class="text-sm font-bold text-zinc-300">${amenity}</span>
                            </div>
                        `).join('')}
                    </div>
                </div>
            </div>

            <!-- Right Column: Pricing & Contact -->
            <div class="lg:col-span-1">
                <div class="sticky top-32 space-y-6">
                    <div class="bg-white/5 backdrop-blur-2xl border border-white/10 rounded-[2.5rem] p-10 shadow-2xl">
                        <div class="flex justify-between items-end mb-8">
                            <div>
                                <p class="text-[10px] font-black text-zinc-500 uppercase tracking-widest mb-1">Monthly Rent</p>
                                <span class="text-4xl font-black text-white">KSh ${house.price.toLocaleString()}</span>
                            </div>
                            <span class="text-zinc-500 font-bold">/ month</span>
                        </div>
                        
                        <div class="space-y-4 mb-8">
                            <div class="flex justify-between text-sm">
                                <span class="text-zinc-500 font-bold">Service Charge</span>
                                <span class="text-white font-black">Included</span>
                            </div>
                            <div class="flex justify-between text-sm">
                                <span class="text-zinc-500 font-bold">Security Deposit</span>
                                <span class="text-white font-black">1 Month</span>
                            </div>
                        </div>

                        <button onclick="triggerPayment(${house.id})" class="w-full bg-blue-600 text-white font-black py-5 rounded-2xl hover:bg-blue-500 transition-all shadow-xl shadow-blue-600/20 mb-4 flex items-center justify-center gap-2">
                            Unlock Landlord Contact
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
                        </button>
                        <p class="text-[10px] text-center text-zinc-500 font-black uppercase tracking-widest">Unlock for KES 1,000</p>
                    </div>

                    <div class="bg-white/5 border border-white/10 rounded-[2.5rem] p-8">
                        <div class="flex items-center gap-4 mb-6">
                            <div class="w-12 h-12 bg-zinc-800 rounded-full flex items-center justify-center text-white font-black">JD</div>
                            <div>
                                <h4 class="text-white font-black">John Doe</h4>
                                <p class="text-zinc-500 text-xs font-bold">Verified Landlord</p>
                            </div>
                        </div>
                        <button class="w-full border border-white/10 text-white font-black py-4 rounded-2xl hover:bg-white/5 transition-all">
                            Schedule a Viewing
                        </button>
                    </div>
                </div>
            </div>
        </section>
    </main>
    ${getFooter()}
    <script>
        function triggerPayment(houseId) {
            fetch('/api/trigger-payment', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ houseId })
            })
            .then(res => res.json())
            .then(data => alert(data.message));
        }
    </script>
</body>
</html>
`;
