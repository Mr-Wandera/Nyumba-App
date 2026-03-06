package models

type House struct {
	ID            int      `json:"id"`
	BuildingName  string   `json:"building_name"`
	Location      string   `json:"location"`
	MapLink       string   `json:"map_link"`
	Price         float64  `json:"price"`
	ImageURLs     []string `json:"image_urls"`
	IsPaid        bool     `json:"is_paid"`
	LandlordPhone string   `json:"landlord_phone"`
	Bedrooms      int      `json:"bedrooms"`
	Bathrooms     int      `json:"bathrooms"`
}
