package main

import (
	"log"
	"mindly-be/handlers"
	"mindly-be/middleware"
	"net/http"
)

func main() {
	http.HandleFunc("/register", middleware.CORSMiddleware(handlers.RegisteHandler))
	http.HandleFunc("/login", middleware.CORSMiddleware(handlers.LoginHandler))

	log.Println("Server is Running in port 4200")
	log.Fatal(http.ListenAndServe(":4200", nil))
}
