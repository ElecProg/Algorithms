package stringproblems

// KnuthMorrisPratt implements KMP substring search
func KnuthMorrisPratt(text string, pattern string) int {
	dfa := generateDFA(pattern)

	s, i := 0, 0
	for ; s < len(pattern) && i < len(text); i++ {
		s = dfa[text[i]][s]
	}

	if s == len(pattern) {
		return i - len(pattern)
	}

	return len(text)
}

func generateDFA(pattern string) [][]int {
	dfa := make([][]int, chars)
	for i := range dfa {
		dfa[i] = make([]int, len(pattern))
	}

	dfa[pattern[0]][0] = 1

	for x, i := 0, 1; i < len(pattern); i++ {
		for c := 0; c < chars; c++ {
			dfa[c][i] = dfa[c][x]
		}

		dfa[pattern[i]][i] = i + 1
		x = dfa[pattern[i]][x]
	}

	return dfa
}
