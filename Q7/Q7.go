package main

import (
	"fmt"
)

// Recursive function to calculate factorial
func factorial(n int) int {
	// Base case: If n is 0 or 1, return 1
	if n == 0 || n == 1 {
		return 1
	}
	// Recursive case: n * factorial(n-1)
	return n * factorial(n-1)
}

func main() {
	// Declare a variable to store user input
	var num int

	// Prompt user for input
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	// Handle negative input
	if num < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
		return
	}

	// Call the recursive function and display the result
	fmt.Printf("Factorial of %d is: %d\n", num, factorial(num))
}
