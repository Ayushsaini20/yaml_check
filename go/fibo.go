package main
 
import (
	"fmt"
	"os"
)
 
func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}
 
func main() {
	var terms int
	fmt.Print("Enter the number of terms: ")
	if _, err := fmt.Scanln(&terms); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
 
	if terms <= 0 {
		fmt.Println("Please enter a positive integer.")
	} else {
		fmt.Printf("Fibonacci sequence with %d terms:\n", terms)
		fibonacci(terms)
	}
}
