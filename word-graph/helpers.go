package wordgraph

func oneAway(a, b string) bool {
	difs := 0

	for i := range a {
		if a[i] != b[i] {
			difs++
		}
	}

	return difs == 1
}
