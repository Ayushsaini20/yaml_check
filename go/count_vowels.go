package main

import (
    "fmt"
    "strings"
)

func countVowels(input string) int {
    input = strings.ToLower(input)
    count := 0
    for _, char := range input {
        if strings.ContainsRune("aeiou", char) {
            count++
        }
    }
    return count
}

func main() {
    var text string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&text)

    vowels := countVowels(text)
    fmt.Printf("The string contains %d vowel(s).\n", vowels)
}
