package main

import "sync"

// --- CONSTANTS ---
const (
	userFile   = "users.json"
	houseFile  = "uploads/houses.json"
	CookieName = "session_token"
)

// --- DATA STRUCTURES ---

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type House struct {
	ID           int      `json:"id"`
	BuildingName string   `json:"building_name"` // <-- NEW FIELD
	Location     string   `json:"location"`
	Price        float64  `json:"price"`
	Type         string   `json:"type"`
	Utilities    float64  `json:"utilities"`
	Details      string   `json:"details"`
	ImageURLs    []string `json:"image_urls"`
	Phone        string   `json:"phone"`
	Owner        string   `json:"owner"`
	IsBooked     bool     `json:"is_booked"`
	TenantPhone  string   `json:"tenant_phone"`
	MapURL       string   `json:"map_url"`
}

// --- GLOBAL MEMORY ---
var (
	users  []User
	houses []House
	mutex  sync.Mutex
)
