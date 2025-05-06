package main
 
import (
	"fmt"
	"os"
)
 
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
 
func main() {
	var year int
	fmt.Print("Enter a year: ")
 
	// Read input and check for errors
	_, err := fmt.Scanln(&year)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
 
	if isLeapYear(year) {
		fmt.Printf("%d is a Leap Year ✅\n", year)
	} else {
		fmt.Printf("%d is NOT a Leap Year ❌\n", year)
	}
}
