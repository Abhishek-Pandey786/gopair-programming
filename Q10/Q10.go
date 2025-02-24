package main

import (
	"fmt"
)

func main() {
	// Declare a variable to store user input
	var num, sum int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	// Ensure the number is positive
	if num < 0 {
		fmt.Println("Please enter a positive number.")
		return
	}

	// Loop to extract and sum digits
	originalNum := num // Store original number for display
	for num > 0 {
		digit := num % 10  // Get the last digit
		sum += digit       // Add it to sum
		num /= 10          // Remove the last digit
	}

	// Display the result
	fmt.Printf("Sum of digits of %d is: %d\n", originalNum, sum)
}
