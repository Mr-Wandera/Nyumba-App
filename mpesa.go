package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// --- CONFIGURATION (SANDBOX) ---
const (
	ConsumerKey    = "COBGyH3dHvYrVjLKG0Znfh8RR1yAPeVbZ6hZitAwgvquIqhL"
	ConsumerSecret = "ovklACIWd4ZMihM4Vv28TAwgEBG8MywaI5FOnHahzIPXAG16CTCikL2RSSqT4cog"
	ShortCode      = "174379" // Safaricom Test Paybill
	Passkey        = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
	BaseURL        = "https://sandbox.safaricom.co.ke"
)

// --- STRUCTS ---
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type STKPushBody struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

// --- FUNCTIONS ---

// 1. Get the "Access Token" (The Badge)
func getAccessToken() (string, error) {
	url := BaseURL + "/oauth/v1/generate?grant_type=client_credentials"
	req, _ := http.NewRequest("GET", url, nil)

	// Create the Basic Auth header (Key:Secret)
	auth := ConsumerKey + ":" + ConsumerSecret
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+encodedAuth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result AccessTokenResponse
	json.Unmarshal(body, &result)

	return result.AccessToken, nil
}

// 2. Trigger the STK Push (The Popup)
func initiateSTKPush(phoneNumber string, amount string) (string, error) {
	// A. Get Token
	token, err := getAccessToken()
	if err != nil {
		return "", err
	}

	// B. Create Timestamp & Password
	timestamp := time.Now().Format("20060102150405")
	password := base64.StdEncoding.EncodeToString([]byte(ShortCode + Passkey + timestamp))

	// C. Create Request Body
	// NOTE: We are using your Render URL for the callback!
	reqBody := STKPushBody{
		BusinessShortCode: ShortCode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            amount,
		PartyA:            phoneNumber, // Your phone
		PartyB:            ShortCode,   // Where money goes
		PhoneNumber:       phoneNumber,
		CallBackURL:       "https://nyumba-app.onrender.com/callback", // 👈 IMPORTANT
		AccountReference:  "NyumbaApp",
		TransactionDesc:   "Rent Payment",
	}

	jsonBody, _ := json.Marshal(reqBody)

	// D. Send Request
	url := BaseURL + "/mpesa/stkpush/v1/processrequest"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBytes, _ := ioutil.ReadAll(resp.Body)
	return string(responseBytes), nil
}
