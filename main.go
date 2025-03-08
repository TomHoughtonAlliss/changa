package main

import (
	wordgraph "github.com/changa/word-graph"
)

func main() {
	g := wordgraph.NewGraph(4)

	p := g.Paths("GLEN", "CURT", 8)

	wordgraph.PrintPaths(p)
}
