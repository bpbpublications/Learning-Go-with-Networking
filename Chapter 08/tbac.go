package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	// Secret key used to sign JWT tokens
	secretKey = []byte("your_secret_key")
)

// User represents a user in the system
type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

// Claims represents the JWT claims
type Claims struct {
	User
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token for the given user
func GenerateJWT(user User) (string, error) {
	// Set expiration time for the token
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create claims
	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Middleware function to validate JWT token
func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If token is valid, continue with the next handler
		next.ServeHTTP(w, r)
	}
}

// ProtectedHandler is a handler for protected resources
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are authorized to access this resource!")
}

// LoginHandler is a handler for generating JWT token upon successful authentication
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// In a real-world scenario, perform user authentication here
	user := User{Username: "user123", Role: "admin"}

	// Generate JWT token
	tokenString, err := GenerateJWT(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the token as response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, tokenString)
}

func main() {
	// Define HTTP routes
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/protected", ValidateToken(ProtectedHandler))

	// Start the server
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
