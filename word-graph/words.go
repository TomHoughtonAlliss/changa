package wordgraph

import (
	"bufio"
	"os"
	"strings"
)

// readRaw reads in the contents of a given file, assuming one string per line.
//
// It outputs each line as an element in a []string array, with line delimiters trimmed.
func readRaw(fName string, length int) []string {
	file, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if len(word) == 4 {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return words
}