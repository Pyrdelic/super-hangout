package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pyrdelic/super-hangout/fileops"
	"github.com/pyrdelic/super-hangout/prints"
	"golang.org/x/term"
)

func main() {
	word := fileops.ReadRandomLine("kotus_sanat.txt")
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

	for lives > 0 {
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}
		guessed = append(guessed, char)
		prints.Word(word, guessed)
		prints.Guessed(guessed)
		lives--
	}

}
