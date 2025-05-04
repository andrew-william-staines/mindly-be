package models

type OTPRequest struct {
	Email string `json:"email"`
}

type OTPVerify struct {
	Email string `json:"email"`
	OTP string `json:"otp"`
}

