package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
)

// greetUser prints a personalized greeting with the current time
func greetUser(name string) {
    name = strings.Title(strings.TrimSpace(name)) // Capitalize and trim
    fmt.Printf("\nHello, %s!\n", name)
    fmt.Println("Welcome to the Go world 🚀")
    fmt.Println("Current time:", time.Now().Format("Mon Jan 2 15:04:05 2006"))
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name: ")
    inputName, _ := reader.ReadString('\n')

    greetUser(inputName)
}
