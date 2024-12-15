package internal

import (
	"fmt"
	"log"
	"math/rand"
)

func GameLoop(chances int) {
	var guess int

	victory := false
	number := rand.Intn(100)

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
