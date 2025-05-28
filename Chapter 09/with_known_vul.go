// main.go
package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux" // Intentionally using an outdated version
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleRequest)

	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, OWASP Demo!")
}
