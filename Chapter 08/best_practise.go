package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define a secret key used for signing JWT tokens
var jwtKey = []byte("secret_key")

// Credentials struct to represent login credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims struct to represent JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Handler for the login route
func login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	creds.Username = r.Form.Get("username")
	creds.Password = r.Form.Get("password")

	// Dummy authentication check (replace with your own logic)
	if creds.Username != "user" || creds.Password != "password" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create JWT token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set token in response header
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
}

// Handler for the protected route
func protected(w http.ResponseWriter, r *http.Request) {
	// Get token from request cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse and validate JWT token
	tokenString := cookie.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Serve protected content
	fmt.Fprintf(w, "Welcome, %s!", claims.Username)
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/protected", protected)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
