package templates

import (
	"encoding/json"
	"fmt"
	"nyumba/models" // Standardizing to the central model package
)



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

func getFooter() string {
	return `
	<footer class="bg-zinc-950 border-t border-white/5 py-20">
		<div class="max-w-7xl mx-auto px-6 grid grid-cols-1 md:grid-cols-4 gap-12">
			<div class="col-span-1 md:col-span-2">
				<h2 class="text-3xl font-black tracking-tighter mb-6">Nyumba.</h2>
				<p class="text-zinc-500 max-w-sm leading-relaxed mb-8">Kenya's premier sanctuary discovery platform. Connecting serious renters with verified landlords directly.</p>
				<div class="flex gap-4">
					<a href="#" class="p-3 bg-white/5 rounded-2xl hover:bg-white/10 transition-colors"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"/></svg></a>
					<a href="#" class="p-3 bg-white/5 rounded-2xl hover:bg-white/10 transition-colors"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="20" height="20" x="2" y="2" rx="5" ry="5"/><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"/><line x1="17.5" x2="17.51" y1="6.5" y2="6.5"/></svg></a>
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
	</footer>`
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
					<div class="absolute top-6 left-6 bg-blue-600 text-white px-4 py-1.5 rounded-full text-[10px] font-black tracking-widest uppercase">
						${house.location}
					</div>
				</div>
				<div class="p-8">
					<h3 class="text-2xl font-black tracking-tighter mb-2">${house.building_name}</h3>
					<div class="flex justify-between items-center pt-6 border-t border-white/5">
						<span class="text-xl font-black">KSh ${house.price.toLocaleString()}</span>
						<a href="/explore" class="text-sm font-black text-blue-400 hover:text-blue-300 transition-colors">View Details</a>
					</div>
				</div>
			%s;
			grid.appendChild(card);
		});
	</script>
</body>
</html>`, getHeader(), getFooter(), featuredJSON, "`", "`")
}

func GetExploreHTML(houses []models.House) string {
	return fmt.Sprintf(`
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
	%s
	<main class="max-w-7xl mx-auto px-6 pt-32 pb-20">
		<h2 class="text-5xl md:text-7xl font-black tracking-tighter leading-none">Available Sanctuaries</h2>
		<div id="houses-grid" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 mt-16"></div>
	</main>
	%s
</body>
</html>`, getHeader(), getFooter())
}

func GetLandlordHTML() string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Landlord Portal - Nyumba</title>
	<script src="https://cdn.tailwindcss.com"></script>
</head>
<body>
	%s
	<main class="max-w-7xl mx-auto px-6 pt-32 pb-20">
		<h2 class="text-6xl font-black tracking-tighter mb-12">Landlord Dashboard</h2>
	</main>
	%s
</body>
</html>`, getHeader(), getFooter())
}

func GetAuthHTML(mode string) string {
	isLogin := mode == "Login"
	
	// Determine form fields based on mode
	var formFields string
	if isLogin {
		formFields = `
			<div class="space-y-4">
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Email</label>
					<input type="email" name="email" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="your@email.com">
				</div>
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Password</label>
					<input type="password" name="password" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="••••••••">
				</div>
			</div>
			<button type="submit" class="w-full bg-white text-black font-black py-4 rounded-xl hover:scale-[1.02] transition-transform mt-6">Sign In</button>
			<p class="text-center text-zinc-500 text-sm mt-6 font-medium">Don't have an account? <a href="/signup" class="text-blue-400 hover:text-blue-300 font-bold">Sign up</a></p>
		`
	} else {
		formFields = `
			<div class="space-y-4">
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Full Name</label>
					<input type="text" name="name" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="Abdul Wandera">
				</div>
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Email</label>
					<input type="email" name="email" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="your@email.com">
				</div>
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Phone</label>
					<input type="tel" name="phone" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="+254 700 000 000">
				</div>
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Password</label>
					<input type="password" name="password" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="••••••••">
				</div>
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Account Type</label>
					<select name="role" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white focus:outline-none focus:border-blue-500 transition-colors">
						<option value="" disabled selected class="bg-zinc-900">Select account type</option>
						<option value="renter" class="bg-zinc-900">Looking for a home (Renter)</option>
						<option value="landlord" class="bg-zinc-900">Listing properties (Landlord)</option>
					</select>
				</div>
			</div>
			<button type="submit" class="w-full bg-white text-black font-black py-4 rounded-xl hover:scale-[1.02] transition-transform mt-6">Create Account</button>
			<p class="text-center text-zinc-500 text-sm mt-6 font-medium">Already have an account? <a href="/login" class="text-blue-400 hover:text-blue-300 font-bold">Log in</a></p>
		`
	}

	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>%s - Nyumba</title>
	<script src="https://cdn.tailwindcss.com"></script>
	<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800;900&display=swap" rel="stylesheet">
	<style>
		body { font-family: 'Inter', sans-serif; }
	</style>
</head>
<body class="bg-[#09090b] text-white flex items-center justify-center min-h-screen p-6 bg-mesh">
	<div class="w-full max-w-md">
		<!-- Logo -->
		<div class="text-center mb-8">
			<a href="/" class="text-3xl font-black tracking-tighter">Nyumba.</a>
			<p class="text-zinc-500 text-sm mt-2 font-medium">Find your sanctuary</p>
		</div>
		
		<!-- Form Card -->
		<div class="bg-white/5 border border-white/10 rounded-[2rem] p-8 md:p-10 backdrop-blur-xl">
			<h2 class="text-3xl font-black tracking-tighter mb-2 text-center">%s</h2>
			<p class="text-zinc-500 text-center mb-8 text-sm font-medium">%s</p>
			
			<form action="/%s" method="POST" class="space-y-4">
				%s
			</form>
			
			<!-- Social Login Divider -->
			<div class="relative my-8">
				<div class="absolute inset-0 flex items-center">
					<div class="w-full border-t border-white/10"></div>
				</div>
				<div class="relative flex justify-center text-xs uppercase">
					<span class="bg-[#09090b] px-2 text-zinc-500 font-bold tracking-widest">Or continue with</span>
				</div>
			</div>
			
			<!-- Google Button -->
			<button type="button" class="w-full bg-white/5 border border-white/10 rounded-xl py-3 px-4 flex items-center justify-center gap-3 hover:bg-white/10 transition-colors font-bold text-sm">
				<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
					<path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
					<path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
					<path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
					<path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
				</svg>
				Google
			</button>
		</div>
		
		<!-- Footer -->
		<p class="text-center text-zinc-600 text-xs mt-8">
			By %s, you agree to our <a href="#" class="text-zinc-400 hover:text-white transition-colors">Terms</a> and <a href="#" class="text-zinc-400 hover:text-white transition-colors">Privacy Policy</a>
		</p>
	</div>
	
	<style>
		.bg-mesh {
			background-image: 
				radial-gradient(at 0%% 0%%, rgba(30, 58, 138, 0.15) 0px, transparent 50%%),
				radial-gradient(at 100%% 0%%, rgba(20, 184, 166, 0.1) 0px, transparent 50%%),
				radial-gradient(at 100%% 100%%, rgba(30, 58, 138, 0.1) 0px, transparent 50%%);
		}
	</style>
</body>
</html>`, mode, mode, 
		map[bool]string{true: "Welcome back to your sanctuary", false: "Join Kenya's premier housing platform"}[isLogin],
		map[bool]string{true: "login", false: "signup"}[isLogin],
		formFields,
		map[bool]string{true: "signing in", false: "signing up"}[isLogin])
}

func GetStaticHTML(title, content string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>%s - Nyumba</title>
	<script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-[#09090b] text-white">
	%s
	<main class="max-w-3xl mx-auto px-6 pt-48 pb-32">
		<h1 class="text-6xl md:text-8xl font-black tracking-tighter mb-12">%s</h1>
		<p class="text-zinc-400 font-medium leading-relaxed">%s</p>
	</main>
	%s
</body>
</html>`, title, getHeader(), title, content, getFooter())
}