package stringproblems

// BoyerMoore implements BM substring search
func BoyerMoore(text string, pattern string) int {
	right := generateSkipTable(pattern)

	for i, skip := 0, 0; i <= len(text)-len(pattern); i, skip = i+skip, 0 {
		for s := len(pattern) - 1; s >= 0; s-- {
			if pattern[s] != text[i+s] {
				skip = s - right[text[i+s]]

				if skip < 1 {
					skip = 1
				}

				break
			}
		}

		if skip == 0 {
			return i
		}
	}

	return len(text)
}

func generateSkipTable(pattern string) []int {
	right := make([]int, chars)

	for c := 0; c < chars; c++ {
		right[c] = -1
	}

	for i := 0; i < len(pattern); i++ {
		right[pattern[i]] = i
	}

	return right
}
