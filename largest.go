package main

import (
    "fmt"
)

func main() {
    var a, b, c int
    fmt.Print("Enter first number: ")
    fmt.Scanln(&a)
    fmt.Print("Enter second number: ")
    fmt.Scanln(&b)
    fmt.Print("Enter third number: ")
    fmt.Scanln(&c)

    max := a
    if b > max {
        max = b
    }
    if c > max {
        max = c
    }

    fmt.Printf("The largest number is: %d\n", max)
}
