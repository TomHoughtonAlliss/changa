package wordgraph

import (
	"fmt"
	"slices"
	"strings"
)

// oneAway iterates over two strings and compares characters.
//
// It returns true if exactly one character differs between the two strings.
func oneAway(a, b string) bool {
	difs := 0

	for i := range a {
		if a[i] != b[i] {
			difs++
		}
	}

	return difs == 1
}

// merge combines the keys and values of two maps, and returns the result.
func merge(m, n map[string]any) map[string]any {
	for k, v := range n {
		m[k] = v
	}

	return m
}

func PrintPaths(p map[string]any) {
	paths := asArrays(p)
	for _, path := range paths {
		fmt.Println(strings.Join(path, " -> "))
	}
}

func asArrays(p map[string]any) [][]string {
	var getArrays func(any) [][]string
	getArrays = func(i any) [][]string {
		switch t := i.(type) {
		case string:
			return [][]string{{t}}

		case map[string]any:
			tails := make([][]string, 0)
			for k, v := range t {
				nextTails := getArrays(v)
				for _, t := range nextTails {
					tails = append(tails, slices.Insert(t, 0, k))
				}
			}

			return tails

		default:
			panic(fmt.Sprintf("encountered unexpected type %T", t))
		}
	}

	return getArrays(p)
}
