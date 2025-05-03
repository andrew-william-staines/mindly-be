package handlers

import (
	"encoding/json"
	"mindly-be/config"
	"mindly-be/models"
	"mindly-be/utils"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lib/pq"
)

func generateJWT(email string) (string, int64, error) {
	expiryTime := time.Now().Add(60 * time.Minute).Unix()
	claims := jwt.MapClaims{
		"email": email,
		"exp": expiryTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString,err := token.SignedString([]byte("your-secret-key"))
	if(err != nil) {
		return "",0,err;
	}

	return tokenString, expiryTime, nil
}

func RegisteHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, `{"message": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, `{"message": "Failed to hash password"}`, http.StatusInternalServerError)
		return
	}

	if !user.Name.Valid {
		user.Name.String = ""
	}

	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, hashedPwd)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" && strings.Contains(pqErr.Message, "email") {
				http.Error(w, `{"message": "Email already exists"}`, http.StatusConflict)
				return
			}
		}
		http.Error(w, `{"message": "Failed to save user"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func LoginHandler (w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if(err != nil) {
		http.Error(w, `{"message" : "Invalid Request Body"}`, http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	var hashedPwd string

	err = db.QueryRow("SELECT password FROM users WHERE email=$1", user.Email).Scan(&hashedPwd)
	if(err != nil) {
		http.Error(w, `{"message" : "User Not Found"}`, http.StatusUnauthorized)
		return
	}

	if(!utils.CheckPasswordHash(user.Password, hashedPwd)){
		http.Error(w, `{"message" : "Invalid Credentials"}`, http.StatusUnauthorized)
		return
	}

	token, expiry, err := generateJWT(user.Email);
	if err != nil {
		http.Error(w, `"message" : "Failed to create token"`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Conetnt-Type","application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token" : token,
		"expiry": expiry,
	})
}
