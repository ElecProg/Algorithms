package stringsorting

// CountingSort sorts byte arrays by character on index i in ~ 10n
func CountingSort(array [][]byte, idx int) {
	aux := make([][]byte, len(array))
	count := make([]int, 257)

	// Make copy of array and build count array
	for i, str := range array {
		aux[i] = str
		count[str[idx]+1]++
	}

	// Build start index list
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// Sort
	for i := 0; i < len(array); i++ {
		str := aux[i]
		ch := str[idx]
		array[count[ch]] = str
		count[ch]++
	}
}
