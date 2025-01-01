package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create an HTTP server to handle dynamic search requests
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// Get the search term from the query parameter
		searchTerm := r.URL.Query().Get("search_term")
		if searchTerm == "" {
			http.Error(w, "Search term is required", http.StatusBadRequest)
			return
		}

		// Handle successful search
		fmt.Fprintf(w, "Search term: %s", searchTerm)
	})

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
