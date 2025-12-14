package main

import (
	"fmt"      // Provides formatted I/O operations (e.g., Printf, Fprintf)
	"log"      // Provides logging utilities for errors and debugging
	"net/http" // Core Go package for building HTTP servers
)

// hellohandler handles requests made to /hello
func hellohandler(w http.ResponseWriter, r *http.Request) {

	// Reject any request that is not exactly /hello
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Only allow GET requests; reject others (POST, PUT, etc.)
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Write response to the client
	fmt.Fprintf(w, "Hello!")
}
func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
func main() {

	// FileServer serves static files from the ./static directory
	// Example: ./static/index.html becomes available at http://localhost:8080/index.html
	fileserver := http.FileServer(http.Dir("./static"))

	// Handle the root path with the file server
	http.Handle("/", fileserver)

	// Register handler for /form (expects formhandler to be defined elsewhere)
	http.HandleFunc("/form", formhandler)

	// Register handler for /hello
	http.HandleFunc("/hello", hellohandler)

	// Log startup
	fmt.Printf("Starting server at port 8080\n")

	// Start HTTP server on port 8080
	// If the server fails to start, log.Fatal will print the error and exit
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
