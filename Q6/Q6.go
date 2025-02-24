package main

import "fmt"

func main() {
	// Declare a variable to store user input
	var decimal int

	// Prompt the user for input
	fmt.Print("Enter a decimal number: ")
	fmt.Scanln(&decimal)

	// Edge case: If input is 0, binary is also 0
	if decimal == 0 {
		fmt.Println("Binary representation: 0")
		return
	}

	// Convert decimal to binary using a loop
	binary := ""
	num := decimal // Store original number for reference

	for num > 0 {
		remainder := num % 2             // Get remainder (0 or 1)
		binary = fmt.Sprintf("%d%s", remainder, binary) // Prepend remainder to binary string
		num /= 2                          // Divide number by 2
	}

	// Print the binary result
	fmt.Println("Binary representation of", decimal, "is:", binary)
}
