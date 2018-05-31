package stringproblems

// RabinKarp implements RK Monte Carlo fingerprint search
func RabinKarp(text string, pattern string) int {
	// Calculate chars ^ (len(pattern) -1) % q
	c := int64(1)
	for i := 1; i <= len(pattern)-1; i++ {
		c = (chars * c) % q
	}

	patternHash := hash(pattern, len(pattern))

	textHash := hash(text, len(pattern))
	if textHash == patternHash {
		return 0
	}

	for i := len(pattern); i < len(text); i++ {
		textHash = (textHash + q - c*int64(text[i-len(pattern)])%q) % q
		textHash = (textHash*chars + int64(text[i])) % q

		if patternHash == textHash {
			return i - len(pattern) + 1
		}
	}

	return len(text)
}

func hash(key string, m int) int64 {
	h := int64(0)

	for j := 0; j < m; j++ {
		h = (chars*h + int64(key[j])) % q
	}

	return h
}
