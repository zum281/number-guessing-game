package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	var chances, guess int
	var difficulty string
	var level string
	victory, validDifficulty := false, false

	number := rand.Intn(100)

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Println()

	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	fmt.Println("Type 'quit' or 'q' to exit the game.")

	for {

		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&difficulty)
		if err != nil {
			log.Fatal(err)
		}
		switch difficulty {
		case "1":
			chances = 10
			level = "Easy"
			validDifficulty = true
		case "2":
			chances = 5
			level = "Medium"
			validDifficulty = true
		case "3":
			chances = 3
			level = "Hard"
			validDifficulty = true
		case "quit", "q":
			fmt.Println("Thank you for playing the game!")
			return
		default:
			fmt.Println("Invalid choice! Please select a valid difficulty level.")
			validDifficulty = false
		}

		if validDifficulty {
			break
		}
	}

	fmt.Printf("Great! You have selected the %s difficulty level\n", level)
	fmt.Println("Let's start the game!")

	for i := 0; i < chances; i++ {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanln(&guess)
		if err != nil {
			log.Fatal(err)
		}

		if guess == number {
			victory = true
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", i+1)
			break
		}

		if guess < number {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
		}

		if guess > number {
			fmt.Printf("Incorrect! The number is less than %d.\n", guess)
		}
	}

	if !victory {
		fmt.Printf("Sorry! You have exhausted all your chances. The correct number was %d.\n", number)
	}

}
