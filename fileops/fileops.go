package fileops

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"math/rand"
	"os"
)

func ReadRandomLine(path string) string {
	word := readNthLine(rand.Intn(countLines(path)+1), path)
	//fmt.Println("Read: ", word)
	//fmt.Println("Read uppercase: ", strings.ToUpper(word))
	//return readNthLine(rand.Intn(countLines(path)+1), path)
	return word
}

func readNthLine(n int, path string) string {
	line := ""

	file, err := os.Open(path)
	// fatal error if file not found
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// // seek to n:th line
	scanner := bufio.NewScanner(file) // defaults to reading line by line

	i := 0
	for scanner.Scan() {
		if i == n {
			line = scanner.Text()
			//found = true
			break
		}
		i++
	}
	return line
}

// returns the count of '\n's in file
func countLines(path string) int {
	file, fileErr := os.Open(path)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer file.Close()

	buf := make([]byte, 32*1024) // 32k buffer (probably way overkill)
	count := 0
	lineSep := []byte{'\n'} // line separator (newline)

	r := bufio.NewReader(file)

	for {
		c, err := r.Read(buf) // read to buffer
		// count how many separators (newlines) in current buffer
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF: // if EOF reached, return count
			return count
		case err != nil:
			return -1
		}
	}
}
