package models

func GetLandingHTML(houses []House) string {
	return ""
}

func GetExploreHTML(houses []House) string {
	return ""
}

type House struct {
	ID            int      `json:"id"`
	BuildingName  string   `json:"building_name"`
	Location      string   `json:"location"`
	MapLink       string   `json:"map_link"`
	Price         float64  `json:"price"`
	Deposit       float64  `json:"deposit"`
	Type          string   `json:"type"`
	ImageURLs     []string `json:"image_urls"`
	IsPaid        bool     `json:"is_paid"`
	Bedrooms      int      `json:"bedrooms"`
	Bathrooms     int      `json:"bathrooms"`
	LandlordPhone string   `json:"landlord_phone"`
	Description   string   `json:"description"`
}
