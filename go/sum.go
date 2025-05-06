package main

import (
    "fmt"
)

func sumOfDigits(n int) int {
    sum := 0
    for n != 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}

func main() {
    var number int
    fmt.Print("Enter a number: ")
    fmt.Scanln(&number)

    result := sumOfDigits(number)
    fmt.Printf("Sum of digits: %d\n", result)
}
