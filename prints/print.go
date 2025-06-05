package prints

import (
	"fmt"
	"slices"
)

// Prints the word with only the guessed letters.
func Word(word string, guessed []rune) {
	for _, c := range word {
		if slices.Contains(guessed, c) {
			fmt.Print(string(c))
		} else {
			fmt.Print("_")
		}
	}
	fmt.Println()
}

// Prints the letters player has quessed.
func Guessed(guessed []rune) {
	fmt.Print("guessed: ")
	for _, c := range guessed {
		fmt.Print(string(c))
	}
	fmt.Println()
}

func Lives(lives int) {
	fmt.Println("lives: ", lives)
}
