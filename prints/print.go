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

func AsciiArt(lives int) {
	asciiArt :=
		`
	
	
	
	
	
	
	
	`
	switch lives {
	case 0:
		asciiArt = ` ______
 |/   |
 |    O ~ game over
 |   /|\
 |   / \
 |
 ______
/      \`
	case 1:
		asciiArt = ` ______
 |/   |
 |    O
 |
 |
 |
 ______
/      \`
	case 2:
		asciiArt = ` ______
 |/
 |
 |
 |
 |
 ______
/      \`
	case 3:
		asciiArt = `
 |/
 |
 |
 |
 |
 ______
/      \`
	case 4:
		asciiArt = ` 
 |
 |
 |
 |
 |
 ______
/      \`
	case 5:
		asciiArt = `






 ______
/      \`
	default:
		break
	}
	fmt.Println(asciiArt)
}
