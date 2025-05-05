package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    fmt.Println("Hello, World! 👋")

    sum := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", sum)
}
