package wordgraph

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
