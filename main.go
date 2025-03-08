package main

import (
	wordgraph "github.com/changa/word-graph"
)

func main() {
	g := wordgraph.NewGraph(4)

	p := g.Paths("kind", "wild", 3)

	wordgraph.PrintPaths(p)
}
