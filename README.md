# Number Guessing Game

A simple command-line number guessing game written in Go, created as a solution for the [Number Guessing Game Project Challenge](https://roadmap.sh/projects/number-guessing-game).

## Description

This game generates a random number between 1 and 100 and challenges the player to guess it. The game features:

- Three difficulty levels:
  - Easy (10 chances)
  - Medium (5 chances)
  - Hard (3 chances)
- Helpful hints after each guess
- Option to quit the game at difficulty selection

## Requirements

- Go 1.23.4 or higher

## How to Run

1. Clone the repository
2. Navigate to the project directory
3. Run the game:
   ```bash
   go run main.go
   ```

## How to Play

1. Start the game
2. Select difficulty level (1-3) or type 'quit'/'q' to exit
3. Enter your guesses
4. The game will tell you if your guess is too high or too low
5. Try to guess the number before running out of chances!

## Game Rules

- The computer generates a random number between 1 and 100
- Player must guess the number within the allowed attempts based on difficulty
- After each guess, the game provides feedback:
  - If the guess is too high
  - If the guess is too low
  - If the guess is correct
