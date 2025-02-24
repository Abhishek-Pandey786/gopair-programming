package main

import (
	"fmt"
)

func main() {
	// Declare a variable to store N
	var N int

	fmt.Print("Enter a number (N): ")
	fmt.Scanln(&N)

	// Initialize sum to 0
	sum := 0

	// Loop from 1 to N
	for i := 1; i <= N; i++ {
		// Check if the number is even
		if i%2 == 0 {
			sum += i // Add even number to sum
		}
	}

	// Print the sum of even numbers
	fmt.Println("Sum of all even numbers from 1 to", N, "is:", sum)
}
