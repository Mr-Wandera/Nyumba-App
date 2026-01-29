package main

type House struct {
	ID          int      `json:"id"`
	Location    string   `json:"location"`
	Price       float64  `json:"price"`
	Type        string   `json:"type"`
	Details     string   `json:"details"`
	Tags        []string `json:"tags"`
	Utilities   float64  `json:"utilities"`
	ImageURLs   []string `json:"image_urls"`
	Phone       string   `json:"phone"`
	Owner       string   `json:"owner"`
	IsBooked    bool     `json:"is_booked"`
	TenantPhone string   `json:"tenant_phone"`
	MapURL      string   `json:"map_url"` // 👈 NEW: Stores the Google Maps Link
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
}
