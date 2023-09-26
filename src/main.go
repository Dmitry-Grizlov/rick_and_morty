package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize file server for serving static files
	fs := http.FileServer(http.Dir("./templates/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Initialize routes
	http.HandleFunc("/index", Index)

	// Start server
	log.Fatal(http.ListenAndServe(":8081", nil))
}
