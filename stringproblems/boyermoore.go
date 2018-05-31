package stringproblems

// BoyerMoore implements BM substring search
func BoyerMoore(text string, pattern string) int {
	st := generateSkipTable(pattern)

	for i, skip := 0, 0; i <= len(text)-len(pattern); i, skip = i+skip, 0 {
		for j := len(pattern) - 1; j >= 0; j-- {
			if pattern[j] != text[i+j] {
				skip = j - st[text[i+j]]

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
	st := make([]int, chars)

	for c := 0; c < chars; c++ {
		st[c] = -1
	}

	for i := 0; i < len(pattern); i++ {
		st[pattern[i]] = i
	}

	return st
}
