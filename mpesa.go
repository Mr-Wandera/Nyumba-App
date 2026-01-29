package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// --- CONFIGURATION (Sandbox Credentials) ---
const (
	consumerKey    = "y4514nMN2a7A2e23Kk75" // Replace if you have your own
	consumerSecret = "9aB8c7D6e5F4g3H2"     // Replace if you have your own
	shortCode      = "174379"               // Sandbox Shortcode
	passkey        = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
	mpesaAuthURL   = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	mpesaPushURL   = "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	callbackURL    = "https://nyumba-app.onrender.com/callback" // Your Live URL
)

// 1. INITIATE STK PUSH (The function handlers.go is looking for!)
func initiateSTKPush(phoneNumber, amount string) error {
	token, err := getAccessToken()
	if err != nil {
		return err
	}

	timestamp := time.Now().Format("20060102150405")
	password := base64.StdEncoding.EncodeToString([]byte(shortCode + passkey + timestamp))

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	// Safaricom expects standard formatted numbers (2547...)
	// In a real app, we would use a helper to format this properly.

	payload := map[string]string{
		"BusinessShortCode": shortCode,
		"Password":          password,
		"Timestamp":         timestamp,
		"TransactionType":   "CustomerPayBillOnline",
		"Amount":            amount,
		"PartyA":            phoneNumber,
		"PartyB":            shortCode,
		"PhoneNumber":       phoneNumber,
		"CallBackURL":       callbackURL,
		"AccountReference":  "NyumbaApp",
		"TransactionDesc":   "Viewing Fee",
	}

	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", mpesaPushURL, bytes.NewBuffer(jsonPayload))

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to send STK: %s", string(body))
	}

	return nil
}

// 2. HELPER: GET ACCESS TOKEN
func getAccessToken() (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", mpesaAuthURL, nil)

	auth := base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret))
	req.Header.Add("Authorization", "Basic "+auth)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to get token: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["access_token"].(string), nil
}
