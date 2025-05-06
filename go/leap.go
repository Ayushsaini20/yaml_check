package main

import (
    "fmt"
)

func isLeapYear(year int) bool {
    return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
    var year int
    fmt.Print("Enter a year: ")
    fmt.Scanln(&year)

    if isLeapYear(year) {
        fmt.Printf("%d is a Leap Year ✅\n",
