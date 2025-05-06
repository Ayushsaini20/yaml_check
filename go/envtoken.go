package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

func main() {
    token := os.Getenv("GITHUB_TOKEN")
    if token == "" {
        fmt.Println("GITHUB_TOKEN environment variable is not set")
        os.Exit(1)
    }

    req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
    if err != nil {
        fmt.Println("Request creation failed:", err)
        os.Exit(1)
    }

    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Accept", "application/vnd.github+json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Request failed:", err)
        os.Exit(1)
    }

    
    defer func() {
        if err := resp.Body.Close(); err != nil {
            fmt.Println("Error closing response body:", err)
        }
    }()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("GitHub API request failed with status: %s\n", resp.Status)
        os.Exit(1)
    }

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        fmt.Println("Failed to decode JSON:", err)
        os.Exit(1)
    }

    fmt.Println("✅ GitHub user login:", result["login"])
}
