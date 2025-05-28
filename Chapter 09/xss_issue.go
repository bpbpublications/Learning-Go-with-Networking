package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user input from the query parameter.
		name := r.URL.Query().Get("name")

		// Display the user input in the HTML response without proper escaping.
		htmlResponse := fmt.Sprintf("<h1>Hello, %s!</h1>", name)
		w.Write([]byte(htmlResponse))
	})

	http.ListenAndServe(":8080", nil)
}
