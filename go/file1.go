package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	
	apiKey := os.Getenv("API_KEY") // The secret key should be stored in the environment variable "API_KEY"

	if apiKey == "" {
		// If the environment variable is not set, log an error and exit
		fmt.Println("API key not found. Please set the API_KEY environment variable.")
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Using API key: %s", apiKey)
		if err != nil {
			fmt.Println("Error writing to response:", err)
		}
	})

	fmt.Println("🚀 Server running on :8080")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
