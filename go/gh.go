package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    // ⚠️ Hardcoded token for testing only (do NOT commit in real projects)
    token := "ghp_yourHardcodedGitHubTokenHere"

    req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
    if err != nil {
        fmt.Println("Request creation failed:", err)
        return
    }

    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Accept", "application/vnd.github+json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Request failed:", err)
        return
    }

    // Safely close the response body and handle any error
    defer func() {
        if cerr := resp.Body.Close(); cerr != nil {
            fmt.Println("Error closing response body:", cerr)
        }
    }()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("GitHub API request failed: %s\n", resp.Status)
        return
    }

    var data map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    fmt.Println("✅ Authenticated GitHub user:", data["login"])
}
