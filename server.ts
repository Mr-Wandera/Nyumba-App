
import express from 'express';
import fs from 'fs';
import path from 'path';
import { 
    getLandingHTML, 
    getExploreHTML, 
    getLandlordHTML, 
    getAuthHTML, 
    getStaticHTML,
    getPropertyDetailHTML
} from './templates';

const app = express();
const PORT = 3000;

app.use(express.json());

// Helper to read houses
const getHouses = () => {
    try {
        const data = fs.readFileSync(path.join(process.cwd(), 'houses.json'), 'utf-8');
        return JSON.parse(data);
    } catch (err) {
        return [];
    }
};

// Helper to save houses
const saveHouses = (houses: any[]) => {
    fs.writeFileSync(path.join(process.cwd(), 'houses.json'), JSON.stringify(houses, null, 2));
};

// Pages
app.get('/', (req, res) => res.send(getLandingHTML()));
app.get('/explore', (req, res) => res.send(getExploreHTML()));
app.get('/landlord', (req, res) => res.send(getLandlordHTML()));
app.get('/login', (req, res) => res.send(getAuthHTML('login')));
app.get('/signup', (req, res) => res.send(getAuthHTML('signup')));

app.get('/property/:id', (req, res) => {
    const id = parseInt(req.params.id);
    const houses = getHouses();
    const house = houses.find((h: any) => h.id === id);
    if (!house) {
        return res.redirect('/explore');
    }
    res.send(getPropertyDetailHTML(house));
});

app.get('/about', (req, res) => res.send(getStaticHTML('About Us', `
    <p class="mb-4">Nyumba is Kenya's premier house-hunting platform, dedicated to making the process of finding a home transparent, secure, and efficient.</p>
    <p class="mb-4">We believe that everyone deserves a home they love, without the hassle of unreliable agents and hidden fees. By connecting tenants directly with verified landlords, we ensure a smoother transition for everyone involved.</p>
    <p>Our team works tirelessly to verify every listing, providing you with peace of mind as you search for your next sanctuary.</p>
`)));

app.get('/contact', (req, res) => res.send(getStaticHTML('Contact Us', `
    <p class="mb-6">Have questions or need assistance? We're here to help!</p>
    <div class="space-y-4">
        <div>
            <h3 class="font-bold text-gray-900">Email</h3>
            <p>support@nyumba.co.ke</p>
        </div>
        <div>
            <h3 class="font-bold text-gray-900">Phone</h3>
            <p>+254 700 000 000</p>
        </div>
        <div>
            <h3 class="font-bold text-gray-900">Office</h3>
            <p>Nairobi, Kenya</p>
        </div>
    </div>
`)));

// API Routes
app.get('/api/houses', (req, res) => {
    const { search, maxPrice } = req.query;
    let houses = getHouses();

    if (search) {
        const s = (search as string).toLowerCase();
        houses = houses.filter((h: any) => 
            h.building_name.toLowerCase().includes(s) || 
            h.location.toLowerCase().includes(s)
        );
    }

    if (maxPrice) {
        const price = parseInt(maxPrice as string);
        houses = houses.filter((h: any) => h.price <= price);
    }

    res.json(houses);
});

app.post('/api/add-house', (req, res) => {
    const houses = getHouses();
    const newHouse = {
        id: houses.length + 1,
        ...req.body,
        is_paid: false,
        image_urls: [`https://picsum.photos/seed/${Math.random()}/800/600`],
        bedrooms: Math.floor(Math.random() * 3) + 1,
        bathrooms: Math.floor(Math.random() * 2) + 1,
        landlord_phone: "+254700000000",
        description: "A newly listed property on Nyumba."
    };
    houses.push(newHouse);
    saveHouses(houses);
    res.status(201).json(newHouse);
});

app.post('/api/trigger-payment', (req, res) => {
    const { houseId } = req.body;
    // Simulate M-Pesa trigger
    setTimeout(() => {
        res.json({ 
            success: true, 
            message: `M-Pesa STK Push sent for House #${houseId}. Please check your phone to complete the KES 1,000 payment.` 
        });
    }, 1000);
});

app.listen(PORT, '0.0.0.0', () => {
    console.log(`Server running on http://localhost:${PORT}`);
});
