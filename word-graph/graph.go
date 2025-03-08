package wordgraph

import (
	"fmt"
	"slices"
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

func (g WordGraph) Paths(a string, b string, max int) map[string]any {
	var getNext func(string, string, int) map[string]any
	getNext = func(end string, last string, length int) map[string]any {
		var dud map[string]any

		tail := make(map[string]any)
		nextWords := g[last]

		if slices.Contains(nextWords, end) {
			tail[last] = end
			return tail
		}

		if length == max - 1 {
			return dud
		}

		newTail := make(map[string]any)
		for _, word := range nextWords {
			childPath := getNext(b, word, length+1)

			if len(childPath) != 0 {
				newTail = merge(newTail, childPath)
			}
		}

		if len(newTail) == 0 {
			return dud
		}

		tail[last] = newTail
		return tail
	}

	paths := getNext(b, a, 1)

	return paths
}

func (g WordGraph) Print() {
	for k, v := range g {
		if len(v) > 0 {
			fmt.Printf("%v: %v\n", k, strings.Join(v, ", "))
		}
	}
}
