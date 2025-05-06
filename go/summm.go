package main
 
import (
    "fmt"
    "os"
)
 
func main() {
    var number int
    fmt.Print("Enter a number: ")
 
    // Read input and check for errors
    _, err := fmt.Scanln(&number)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
        os.Exit(1)
    }
 
    // Calculate sum of numbers from 1 to number
    sum := 0
    for i := 1; i <= number; i++ {
        sum += i
    }
 
    fmt.Printf("Sum of numbers from 1 to %d is: %d\n", number, sum)
}
