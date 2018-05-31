package stringsorting

// CountingSort sorts byte arrays by character on index i in ~ 11n
// strings without a value a index i are considered to be first in order.
func CountingSort(array []string, idx int) {
	aux := make([]string, len(array))

	countingsort(array, idx, aux, 0, len(array)-1)
}

// Returns the value of the character at index idx, if idx is too large, returns -1
func getchr(str string, idx int) int {
	if idx >= len(str) {
		return -1
	}

	return int(str[idx])
}

// Internal counting sort returns the counting array
func countingsort(array []string, idx int, aux []string, lo int, hi int) []int {
	count := make([]int, 256+2)

	// Make copy of array and build count array
	for i := lo; i <= hi; i++ {
		str := array[i]
		aux[i] = str
		count[getchr(str, idx)+2]++
	}

	// Build start index list
	for i := 2; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// Sort
	for i := lo; i <= hi; i++ {
		str := aux[i]
		ch := getchr(str, idx)
		array[lo+count[ch+1]] = str
		count[ch+1]++
	}

	return count
}
