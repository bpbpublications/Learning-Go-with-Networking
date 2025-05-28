package main

import (
	"fmt"
	"net/http"
)

var users = map[string]string{
	"alice": "password123",
	"bob":   "letmein",
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

	// Check if the username exists and the password matches (insecurely)
	if storedPassword, ok := users[username]; ok && storedPassword == password {
		// Successful login, redirect to the dashboard
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	// Authentication failed, show an error message
	http.Error(w, "Invalid username or password", http.StatusUnauthorized)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// In a real application, you would check if the user is authenticated here
	// and then serve the dashboard page.
	fmt.Fprintln(w, "Welcome to the dashboard!")
}
