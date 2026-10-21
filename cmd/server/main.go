package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "method not allowed")
		return
	} 
	fmt.Fprintln(w, "This server is health!")
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "method not allowed")
		return 
	}
	
	w.Header().Set("Content-Type", "application/json")
	
	response := map[string]string {
		"name": "URL shortener API",
		"status": "running",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL Shortener API")
	})

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/info", infoHandler)


	fmt.Println("Server running on :8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}