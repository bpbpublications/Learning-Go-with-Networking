package main

import (
	"net/http"
	"github.com/ponzu-cms/ponzu/system/admin"
	"github.com/ponzu-cms/ponzu/system/item"
	"github.com/ponzu-cms/ponzu/system/router"
)

func main() {
	// Initialize Ponzu server
	router.SetHandler(http.DefaultServeMux)

	// Start Ponzu admin interface
	go admin.Start()

	// Start your web service
	http.ListenAndServe(":8080", nil)
}
