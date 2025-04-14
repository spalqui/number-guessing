package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	MaxNumber = 100

	EasyDifficulty = iota
	MediumDifficulty
	HardDifficulty
)

func main() {
	printIntroAndRules()

	for {
		maxAttempts := getDifficultyLevel()

		fmt.Println("Let's start the game!")

		correctNumber := generateRandomNumber()

		playGame(maxAttempts, correctNumber)

		if userWantsToQuit() {
			break
		}
	}

	fmt.Println("The end of the game!")
}

func printIntroAndRules() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
}

func getDifficultyLevel() int {
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scan(&choice)

	return mapDifficultyToMaxAttempts(choice)
}

func mapDifficultyToMaxAttempts(difficulty int) int {
	switch difficulty {
	case EasyDifficulty:
		fmt.Println("Great! You've selected the Easy difficulty level.")
		return 10
	case MediumDifficulty:
		fmt.Println("Great! You've selected the Medium difficulty level.")
		return 5
	case HardDifficulty:
		fmt.Println("Great! You've selected the Hard difficulty level.")
		return 3
	default:
		fmt.Println("Oops! Invalid choice. Please select a valid difficulty level between (1-3).")
		fmt.Println("")
		return getDifficultyLevel()
	}
}

func generateRandomNumber() int {
	return rand.Intn(MaxNumber + 1)
}

func playGame(maxAttempts, correctNumber int) {
	fmt.Printf("You have %d chances to guess the correct number.\n", maxAttempts)

	for i := 0; i < maxAttempts; i++ {
		guess := getUserGuess()

		if isGuessCorrect(guess, correctNumber, i) {
			return
		}
	}
	fmt.Printf("Sorry! You've used all your chances. The correct number was %d.\n", correctNumber)
}

func getUserGuess() int {
	fmt.Println("")
	fmt.Print("Enter your guess: ")

	var guess int
	fmt.Scan(&guess)
	return guess
}

func isGuessCorrect(guess, correctNumber, attempt int) bool {
	if guess < correctNumber {
		fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
		return false
	}

	if guess > correctNumber {
		fmt.Printf("Incorrect! The number is less than %d.\n", guess)
		return false
	}

	fmt.Printf("Congratulations! You've guessed the correct number in %d attempts.\n", attempt)
	return true
}

func userWantsToQuit() bool {
	fmt.Print("Do you want to play again? (y/n): ")
	var playAgain string
	fmt.Scan(&playAgain)

	if strings.ToLower(playAgain) == "y" {
		fmt.Println("Great! Let's play again!")
		return false
	}

	fmt.Println("Thank you for playing! Goodbye!")
	return true
}
