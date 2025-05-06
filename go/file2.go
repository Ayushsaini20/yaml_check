package main

import (
	"fmt"
	"net/http"
)

func main() {
	
	apiKey := "sk_live_51H8YoXXXXXXXSECRETXXXX" // hardcoded secret

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Using API key: %s", apiKey)
	})

	fmt.Println(" Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
