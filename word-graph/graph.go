package wordgraph

import "fmt"

func Graph() {
	words := readRaw("./word-graph/words.txt", 4)

	fmt.Println(words)
}
