package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Enter the number of Fibonacci terms: ")
	fmt.Scanln(&n)

	// Handle invalid input
	if n <= 0 {
		fmt.Println("Please enter a positive integer.")
		return
	}

	// Initialize the first two Fibonacci numbers
	first, second := 0, 1

	// Print Fibonacci sequence
	fmt.Println("Fibonacci sequence:")

	for i := 0; i < n; i++ {
		fmt.Print(first, " ") // Print the current Fibonacci number
		next := first + second // Calculate the next Fibonacci number
		first, second = second, next // Update first and second for next iteration
	}
	fmt.Println()
}
