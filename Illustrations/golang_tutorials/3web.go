package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) { // localhost:8000
    fmt.Fprintf(w, "Index")
}

func about_handler(w http.ResponseWriter, r *http.Request) { // localhost:8000/about
    fmt.Fprintf(w, "About")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)
	// ListenAndServe starts an HTTP server with a given address and handler. 
	//  The handler is usually nil, which means to use DefaultServeMux. 
	// Handle and HandleFunc add handlers to DefaultServeMux
	// https://golang.org/pkg/net/http/
}