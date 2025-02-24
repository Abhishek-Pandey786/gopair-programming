package main

import (
	"fmt"
)

func main() {
	// Declare a variable to store the year
	var year int

	fmt.Print("Enter a year: ")
	fmt.Scanln(&year)

	// Check if the year is a leap year using control flow
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is not a leap year.")
	}
}
