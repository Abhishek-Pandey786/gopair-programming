package main

import (
	"fmt"
	"strings"
)

func main() {
	// Declare a variable to store user input
	var input string

	// Prompt user for a string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	// Convert the input string to lowercase to ensure case insensitivity
	input = strings.ToLower(input)

	// Check if the string is a palindrome
	if isPalindrome(input) {
		fmt.Println(input, "is a palindrome.")
	} else {
		fmt.Println(input, "is not a palindrome.")
	}
}

// Function to check if a string is a palindrome
func isPalindrome(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false // Return false if characters don't match
		}
	}
	return true // Return true if the string is a palindrome
}
