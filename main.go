package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"unicode"

	"github.com/pyrdelic/super-hangout/fileops"
	"github.com/pyrdelic/super-hangout/prints"
	"golang.org/x/term"
)

func main() {
	word := fileops.ReadRandomLine("wordlist.txt")
	lives := 10
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

	// TODO: inf game loop, break with ESC, game over or victory

	for {
		// print
		prints.Lives(lives)
		prints.Word(word, guessed)
		prints.Guessed(guessed)
		fmt.Println()

		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}
		//fmt.Println("rune inputted: ", rune(char))

		if char == 27 { // input: ESC
			fmt.Println("Exiting.")
			break
		}
		if slices.Contains(guessed, unicode.ToUpper(char)) {
			fmt.Println("Already guessed")
			continue
		}
		if lives <= 0 {
			fmt.Println("-- GAME OVER --")
			break
		}

		guessed = append(guessed, unicode.ToUpper(char))

		// TODO: lives logic (not in word -> lives--)
		lives--
	}

}
