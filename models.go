package main

// This file holds the "Blueprints" for our data

type House struct {
	ID        int      `json:"id"`
	Location  string   `json:"location"`
	Price     float64  `json:"price"`
	Details   string   `json:"details"`
	Tags      []string `json:"tags"`
	Utilities float64  `json:"utilities"`
	ImageURL  string   `json:"image_url"`
	Phone     string   `json:"phone"`
	Owner     string   `json:"owner"`
	IsBooked  bool     `json:"is_booked"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
}
