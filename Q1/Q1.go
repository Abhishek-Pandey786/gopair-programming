package main

import (
	"fmt"
)

func main() {
	// Declare three variables for user input
	var num1, num2, num3 int

	// Prompt the user to enter three numbers
	fmt.Println("Enter three numbers:")

	// Take input using Scanln
	fmt.Scanln(&num1, &num2, &num3)

	// Determine the largest number using conditional statements
	var largest int // Normal declaration

	if num1 >= num2 && num1 >= num3 {
		largest = num1
	} else if num2 >= num1 && num2 >= num3 {
		largest = num2
	} else {
		largest = num3
	}

	// Print the largest number
	fmt.Println("The largest number is:", largest)
}
