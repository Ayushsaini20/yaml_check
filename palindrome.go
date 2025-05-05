package main

import (
    "fmt"
    "strings"
)

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    n := len(s)
    for i := 0; i < n/2; i++ {
        if s[i] != s[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var input string
    fmt.Print("Enter a word: ")
    fmt.Scanln(&input)

    if isPalindrome(input) {
        fmt.Printf("%s is a palindrome ✅\n", input)
    } else {
        fmt.Printf("%s is not a palindrome ❌\n", input)
    }
}
