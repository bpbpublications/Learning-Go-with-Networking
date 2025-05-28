package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
	"time"
	"github.com/hashicorp/vault/api"
	"crypto/bcrypt"
	"database/sql"
	"github.com/spf13/viper"
)

var (
	db           *sql.DB
	sessionStore sync.Map
)

func init() {

	// Initialize MySQL database connection
	var err error


	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		panic(err)
	}

	// Initialize HashiCorp Vault client
	vaultClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	vaultToken := viper.GetString("vault.token")
	vaultClient.SetToken(vaultToken)

	secretPath := viper.GetString("vault.secretPath") 
	fmt.Println("Reading value from Vault")

	mysqlSecret, err := vaultClient.Logical().Read(secretPath)
	if err != nil {
		panic(err)
	}


	mysqlUsername := mysqlSecret.Data["data"].(map[string]interface{})["username"].(string)
	mysqlPassword := mysqlSecret.Data["data"].(map[string]interface{})["password"].(string)

	// Connect to MySQL database
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/dbname", mysqlUsername, mysqlPassword))
	if err != nil {
		panic(err)
	}

	// Check database connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Retrieve hashed password from MySQL database
	storedPassword, err := getPasswordFromDatabase(username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Compare the provided password with the stored hashed password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Authentication successful, generate a session token
	sessionToken := generateSessionToken()
	storeSessionToken(sessionToken, username)

	// Set the session token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour), // Example: session expires in 24 hours
		Path:     "/",
		Secure:   true, // Enable for HTTPS
		HttpOnly: true,
	})

	// Redirect to the dashboard
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Check the user's session
	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Validate the session token
	username, ok := getSessionToken(sessionToken.Value)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Authentication successful, serve the dashboard page
	fmt.Fprintf(w, "Welcome to the dashboard, %s!", username)
}

func getPasswordFromDatabase(username string) (string, error) {
	// Query the hashed password from the MySQL database
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&hashedPassword)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}
	return hashedPassword, nil
}

func generateSessionToken() string {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(token)
}

func storeSessionToken(sessionToken, username string) {
	// In a real-world scenario, store the session token in a secure data store
	// Here, we're using an in-memory map for demonstration purposes
	sessionStore.Store(sessionToken, username)
}

func getSessionToken(sessionToken string) (string, bool) {
	// In a real-world scenario, retrieve the username associated with the session token
	// Here, we're using an in-memory map for demonstration purposes
	if username, ok := sessionStore.Load(sessionToken); ok {
		return username.(string), true
	}
	return "", false
}
