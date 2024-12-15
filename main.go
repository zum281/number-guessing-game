package main

import (
	"zusk/number-guess/internal"
)

func main() {
	var chances int

	internal.Intro()
	chances = internal.PickDifficulty()
	internal.GameLoop(chances)

}
