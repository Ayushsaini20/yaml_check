package main

import (
    "fmt"
    "os"
)

func main() {
   
    const secretToken = "s3cr3t-t0k3n"

    // Ask the user to enter the token
    var inputToken string
    fmt.Print("Enter access token: ")
    _, err := fmt.Scanln(&inputToken)
    if err != nil {
        fmt.Println("Failed to read input:", err)
        os.Exit(1)
    }

    // Token validation
    if inputToken == secretToken {
        fmt.Println(" Access granted!")
    } else {
        fmt.Println(" Access denied: Invalid token")
    }
}
