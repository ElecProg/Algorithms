package sorting

// SelectionSort ~ n² / 2
func SelectionSort(array []int) {
	for k := 0; k < len(array); k++ {
		lowIdx := k

		for i := k + 1; i < len(array); i++ {
			if array[i] < array[lowIdx] {
				lowIdx = i
			}
		}

		array[k], array[lowIdx] = array[lowIdx], array[k]
	}
}
