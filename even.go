package main

import (
    "fmt"
)

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scanln(&num)

    if num%2 == 0 {
        fmt.Printf("%d Even \n", num)
    } else {
        fmt.Printf("%d Odd \n", num)
    }
}
