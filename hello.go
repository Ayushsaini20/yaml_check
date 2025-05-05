package main

import (
    "fmt"
    "time"
)

func main() {
    name := "Gopher"
    fmt.Printf("Hello, %s!\n", name)
    
    currentTime := time.Now()
    fmt.Println("Current time is:", currentTime.Format("Mon Jan 2 15:04:05 2006"))
}
