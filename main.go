package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/pyrdelic/super-hangout/fileops"
	"github.com/pyrdelic/super-hangout/prints"
	"golang.org/x/term"
)

func main() {
	word := fileops.ReadRandomLine("wordlist.txt")
	lives := 10
	blanksLeft := utf8.RuneCountInString(word)
	guessed := make([]rune, len(word))
	//guessedLetters := ""

	prints.Word(word, guessed[:])
	prints.Guessed(guessed[:])

	reader := bufio.NewReader(os.Stdin)

	// switch stdin into 'raw' mode
	prevState, termErr := term.MakeRaw(int(os.Stdin.Fd()))
	if termErr != nil {
		fmt.Println(termErr)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), prevState)

	// TODO: game over, victory

	for {
		// print
		prints.Lives(lives)
		prints.Word(word, guessed)
		prints.Guessed(guessed)
		fmt.Println("blanksLeft: ", blanksLeft)
		fmt.Println()

		// Read input, force to uppercase
		char, _, err := reader.ReadRune()
		char = unicode.ToUpper(char)

		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}
		//fmt.Println("rune inputted: ", rune(char))

		if char == 27 { // input: ESC
			fmt.Println("Exiting.")
			break
		}
		if slices.Contains(guessed, char) {
			fmt.Println("Already guessed")
			continue
		}
		if strings.Contains(word, string(char)) {
			blanksLeft = blanksLeft - strings.Count(word, string(char))
		} else {
			lives--
		}

		// game over condition
		if lives <= 0 {
			fmt.Println("-- GAME OVER --")
			break
		}
		// victory condition
		if blanksLeft <= 0 {
			fmt.Println(word)
			fmt.Println("VICTORY!")
			break
		}

		guessed = append(guessed, char)
		// TODO: Win-condition
		// TODO: lives logic (not in word -> lives--)
		//lives--
	}

}
