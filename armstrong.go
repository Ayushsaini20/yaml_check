package main

import (
    "fmt"
    "math"
)

func isArmstrong(n int) bool {
    original := n
    var result int
    digits := int(math.Log10(float64(n))) + 1

    for n > 0 {
        digit := n % 10
        result += int(math.Pow(float64(digit), float64(digits)))
        n /= 10
    }

    return result == original
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scanln(&num)

    if isArmstrong(num) {
        fmt.Printf("%d is an Armstrong number \n", num)
    } else {
        fmt.Printf("%d is NOT an Armstrong number \n", num)
    }
}
