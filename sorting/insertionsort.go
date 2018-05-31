package sorting

// InsertionSort ~ n² / 4
func InsertionSort(array []int) {
	for k := 1; k < len(array); k++ {
		for i := k - 1; i >= 0 && array[i] > array[k]; i, k = i-1, k-1 {
			array[k], array[i] = array[i], array[k]
		}
	}
}
