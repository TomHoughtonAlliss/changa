package wordgraph

import "fmt"

func Graph() {
	words := readRaw("./word-graph/words.txt")

	fmt.Println(words)
}
