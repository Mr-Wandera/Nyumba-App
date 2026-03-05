package templates

// GetHTML serves a clean structure that Go won't reject
func GetHTML(isLoggedIn, currentUsername, myHubButton, landlordPanelDisplay string) string {
	return `<!DOCTYPE html><html><head><title>Nyumba</title></head>
	<body style="background: #0f172a; color: white; font-family: sans-serif; padding: 40px;">
		<h1 style="color: #6366f1;">Nyumba.</h1>
		<div id="results-area"></div>
	</body></html>`
}

// GetScripts returns basic JS to load your houses
func GetScripts(isLoggedIn bool, currentUsername string) string {
	return `<script>
		fetch('/houses').then(res => res.json()).then(data => {
			const container = document.getElementById('results-area');
			container.innerHTML = "";
			data.forEach(h => {
				const div = document.createElement('div');
				div.style.border = "1px solid #334155";
				div.style.padding = "20px";
				div.style.margin = "10px 0";
				div.innerHTML = "<h3>" + h.building_name + "</h3><p>" + h.location + "</p>";
				container.appendChild(div);
			});
		});
	</script>`
}

func GetLandingHTML() string {
	return `<!DOCTYPE html><html><body><h1>Welcome to Nyumba</h1><a href="/explore">View Houses</a></body></html>`
}