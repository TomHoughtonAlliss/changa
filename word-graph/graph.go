package wordgraph

import (
	"fmt"
	"strings"
)

type WordGraph map[string][]string

func NewGraph(length int) WordGraph {
	words := readRaw("./word-graph/words.txt", length)

	g := make(WordGraph)

	for _, w := range words {
		g.insertWord(w, words)
	}

	return g
}

func (g WordGraph) insertWord(w string, words []string) {
	adjacent := make([]string, 0)
	for _, cmpWord := range words {
		if oneAway(w, cmpWord) {
			adjacent = append(adjacent, cmpWord)
		}
	}

	g[w] = adjacent
}

func (g WordGraph) Print() {
	for k, v := range g {
		if len(v) > 0 {
			fmt.Printf("%v: %v\n", k, strings.Join(v, ", "))
		}
	}
}
