package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"nyumba/models"
	"time"
)

// TriggerStkPush initiates the KES 1,000 payment gate
func TriggerStkPush(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get phone and house ID from request
	phone := r.FormValue("phone")
	houseID := r.FormValue("house_id")
	if phone == "" || houseID == "" {
		http.Error(w, "Phone and house_id required", http.StatusBadRequest)
		return
	}

	// 1. Get Access Token (Using your keys from Daraja)
	token, err := GetMpesaToken("COBGyH3d...", "ovklACIW...")
	if err != nil {
		http.Error(w, "M-Pesa Auth Failed", http.StatusInternalServerError)
		return
	}

	// 2. Prepare the STK Push Payload
	businessShortCode := "174379"
	passkey := "bfb279f..."
	timestamp := time.Now().Format("20060102150405")

	// ✅ FIX: Actually compute the password using Base64
	password := base64.StdEncoding.EncodeToString([]byte(businessShortCode + passkey + timestamp))

	payload := map[string]interface{}{
		"BusinessShortCode": businessShortCode,
		"Password":          password,
		"Timestamp":         timestamp,
		"TransactionType":   "CustomerPayBillOnline",
		"Amount":            1000,
		"PartyA":            phone,
		"PartyB":            businessShortCode,
		"PhoneNumber":       phone,
		"CallBackURL":       "https://your-app.render.com/api/mpesa/callback",
		"AccountReference":  "Nyumba-" + houseID,
		"TransactionDesc":   "Unlock House Details",
	}

	// 3. Send Request to Safaricom
	jsonData, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to encode request", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "STK Push Failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response
	var stkResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&stkResponse); err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	// Check for errors in response
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("M-Pesa Error: %v", stkResponse), resp.StatusCode)
		return
	}

	// Store CheckoutRequestID with house for later verification
	checkoutID, _ := stkResponse["CheckoutRequestID"].(string)
	fmt.Printf("STK Push sent. CheckoutID: %s\n", checkoutID)

	// TODO: Save checkoutID + houseID mapping to database/file

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":     "Payment Initiated. Check your phone!",
		"checkout_id": checkoutID,
	})
}

// MpesaCallbackHandler listens for payment confirmation from Safaricom
func MpesaCallbackHandler(w http.ResponseWriter, r *http.Request) {
	var callback models.MpesaCallback
	err := json.NewDecoder(r.Body).Decode(&callback)
	if err != nil {
		fmt.Println("Error decoding callback:", err)
		w.WriteHeader(http.StatusOK) // Still respond OK to Safaricom
		return
	}

	// 1. Check if the payment was successful (ResultCode 0 means success)
	if callback.Body.StkCallback.ResultCode == 0 {
		checkoutID := callback.Body.StkCallback.CheckoutRequestID
		fmt.Printf("✅ Payment Successful for CheckoutID: %s\n", checkoutID)

		// 2. Find the house associated with this CheckoutID and mark as paid
		// ✅ FIX: Use Houses from handlers package (already defined in handlers.go)
		for i, h := range Houses {
			// You'll need to store checkoutID when creating the house or in a separate map
			if h.CheckoutRequestID == checkoutID {
				Houses[i].IsPaid = true // Use IsPaid instead of IsBooked
				fmt.Printf("🔓 House %s is now unlocked for the user!\n", h.BuildingName)
				break
			}
		}
	} else {
		fmt.Printf("❌ Payment Failed: %s\n", callback.Body.StkCallback.ResultDesc)
	}

	// 3. Always respond to Safaricom with a success code
	w.WriteHeader(http.StatusOK)
}
