package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    secret := rand.Intn(100) + 1 // random number between 1 and 100
    var guess int

    fmt.Println("🎯 I'm thinking of a number between 1 and 100. Can you guess it?")

    for {
        fmt.Print("Enter your guess: ")
        fmt.Scanln(&guess)

        if guess < secret {
            fmt.Println("Too low! Try again.")
        } else if guess > secret {
            fmt.Println("Too high! Try again.")
        } else {
            fmt.Println("🎉 Correct! You guessed the number!")
            break
        }
    }
}
