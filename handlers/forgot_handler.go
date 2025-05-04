package handlers

import (
	"encoding/json"
	"log"
	"mindly-be/config"
	"mindly-be/models"
	"mindly-be/utils"
	"net/http"
)

func OTPHandler(w http.ResponseWriter, r *http.Request) {
	var req models.OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message": "Invalid request"}`, http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	var email string
	errors := db.QueryRow("SELECT email FROM users WHERE email=$1", req.Email).Scan(&email)
	if errors != nil {
		log.Println("Error:", errors)
		http.Error(w, `{"message" : "User Not Found"}`, http.StatusUnauthorized)
		return
	}

	otp := utils.GenerateNumericOTP(6)
	err := utils.SendForgotMailOTP(req.Email, otp)
	if err != nil {
		http.Error(w, `{"message": "Failed to send OTP"}`, http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(`INSERT INTO otp (email, otp) 
                  VALUES ($1, $2) 
                  ON CONFLICT (email) 
                  DO UPDATE SET otp = EXCLUDED.otp, created_at = CURRENT_TIMESTAMP`, req.Email, otp)
	if err != nil {
		http.Error(w, `{"message": "Failed to save OTP"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "OTP sent"}`))
}

func VerifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	var req models.OTPVerify
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message": "Invalid request"}`, http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	if err := utils.VerifyOTP(db, req.Email, req.OTP); err != nil {
		if err.Error() == "user not found" {
			http.Error(w, `{"message": "User Not Found"}`, http.StatusUnauthorized)
		} else {
			http.Error(w, `{"message": "Incorrect Verification Code"}`, http.StatusUnauthorized)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "OTP verified successfully"}`))
}


func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	var req models.OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message": "Invalid request"}`, http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	otp := utils.GenerateNumericOTP(6)
	err := utils.SendEmailOTP(req.Email, otp)
	if err != nil {
		http.Error(w, `{"message": "Failed to send OTP"}`, http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(`INSERT INTO otp (email, otp) 
                  VALUES ($1, $2) 
                  ON CONFLICT (email) 
                  DO UPDATE SET otp = EXCLUDED.otp, created_at = CURRENT_TIMESTAMP`, req.Email, otp)
	if err != nil {
		http.Error(w, `{"message": "Failed to save OTP"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "OTP sent"}`))
}
