package main

import (
	"fmt"
)

func main() {
	// Declare variables to store user input
	var num1, num2 float64
	var operator string

	// Prompt user for input
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	// Perform calculation using switch statement
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		// Check for division by zero
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Error: Invalid operator. Please use +, -, *, or /.")
		return
	}

	// Display the result
	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
