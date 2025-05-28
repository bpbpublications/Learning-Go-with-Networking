package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user input from the query parameter.
		name := r.URL.Query().Get("name")

		// Use html/template to safely embed user input in the HTML response.
		tmpl := template.Must(template.New("hello").Parse("<h1>Hello, {{.}}!</h1>"))
		tmpl.Execute(w, name)
	})

	http.ListenAndServe(":8080", nil)
}
