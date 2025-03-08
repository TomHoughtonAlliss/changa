package wordgraph

import (
	"bufio"
	"os"
	"strings"
)

// Read in the raw words.txt file - one word per line
func ReadRaw(fName string) []string {
	file, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return words
}