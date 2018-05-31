package stringsorting

// Returns the value of the character at index idx, if idx is too large, returns -1
func getchr(str string, idx int) int {
	if idx >= len(str) {
		return -1
	}

	return int(str[idx])
}
