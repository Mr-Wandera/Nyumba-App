package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type House struct {
	ID                int      `json:"id"`
	BuildingName      string   `json:"building_name"`
	Location          string   `json:"location"`
	Type              string   `json:"type"`
	Price             float64  `json:"price"`
	Utilities         float64  `json:"utilities"`
	Details           string   `json:"details"`
	ImageURLs         []string `json:"image_urls"`
	Phone             string   `json:"phone"`
	Owner             string   `json:"owner"`
	IsBooked          bool     `json:"is_booked"`
	TenantPhone       string   `json:"tenant_phone"`
	MapURL            string   `json:"map_url"`
	CheckoutRequestID string   `json:"checkout_request_id"`
	IsPaid            bool     `json:"is_paid"`
	MapLink           string   `json:"map_link"` // Fixes the Ln 47 handlers error
}

type MpesaCallback struct {
	Body struct {
		StkCallback struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode        int    `json:"ResultCode"`
			ResultDesc        string `json:"ResultDesc"`
		} `json:"stkCallback"`
	} `json:"Body"`
}
