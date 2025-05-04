package main

import (
	"log"
	"mindly-be/handlers"
	"mindly-be/middleware"
	"net/http"
)

func main() {
	http.HandleFunc("/register", middleware.CORSMiddleware(handlers.RegisteHandler))
	http.HandleFunc("/sign-otp", middleware.CORSMiddleware(handlers.VerifyHandler))
	http.HandleFunc("/login", middleware.CORSMiddleware(handlers.LoginHandler))
	http.HandleFunc("/forgot-otp", middleware.CORSMiddleware(handlers.OTPHandler))
	http.HandleFunc("/verify", middleware.CORSMiddleware(handlers.VerifyOTPHandler))
	http.HandleFunc("/resetPassword", middleware.CORSMiddleware(handlers.PasswordHandler))



	log.Println("Server is Running in port 4200")
	log.Fatal(http.ListenAndServe(":4200", nil))
}
