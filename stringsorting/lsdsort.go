package stringsorting

// LSDSort sorts an array of strings on the first k characters
func LSDSort(array []string, k int) {
	aux := make([]string, len(array))

	for d := k - 1; d >= 0; d-- {
		countingsort(array, d, aux, 0, len(array)-1)
	}
}
