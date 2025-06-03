package prints

import (
	"fmt"
)

// check wheter a rune exists in a slice
func inSlice(x rune, guessed []rune) bool {
	//fmt.Println(x, guessed)
	for _, c := range guessed {
		//fmt.Println("loopchar: ", c)
		if x == c {
			//fmt.Println("FOUND!")
			return true
		}
	}
	return false
}

// Prints the word with only the guessed letters.
func Word(word string, guessed []rune) {
	for _, c := range word {
		if inSlice(c, guessed) {
			fmt.Print(string(c))
		} else {
			fmt.Print("_")
		}
	}
	fmt.Println()
}

// Prits the letters player has quessed.
func Guessed(guessed []rune) {
	fmt.Print("guessed: ")
	for _, c := range guessed {
		fmt.Print(string(c))
	}
	fmt.Println()
}
