package internal

import (
	"fmt"
	"log"
	"os"
)

func PickDifficulty() int {
	var level, difficulty string
	var chances int
	validDifficulty := false

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
			os.Exit(0)
		default:
			fmt.Println("Invalid choice! Please select a valid difficulty level.")
			validDifficulty = false
		}

		if validDifficulty {
			break
		}
	}

	fmt.Printf("Great! You have selected the %s difficulty level\n", level)
	return chances
}
