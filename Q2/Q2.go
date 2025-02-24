package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator to get different numbers on each run
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	targetNumber := rand.Intn(100) + 1

	// Declare a variable to store the user's guess
	var userGuess int

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100. Try to guess it!")

	for {
		// Prompt the user for a guess
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&userGuess)

		// Check if the guess is correct, too high, or too low
		if userGuess > targetNumber {
			fmt.Println("Too high! Try again.")
		} else if userGuess < targetNumber {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Congratulations! You guessed the correct number:", targetNumber)
			break // Exit the loop when the correct number is guessed
		}
	}
}
