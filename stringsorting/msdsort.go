package stringsorting

// MSDSort sorts an array of strings
func MSDSort(array []string) {
	aux := make([]string, len(array))
	msdsort(array, 0, aux, 0, len(array)-1)
}

func msdsort(array []string, idx int, aux []string, lo int, hi int) {
	if hi <= lo {
		return
	}

	count := countingsort(array, idx, aux, lo, hi)

	// count[0] is the amount of strings that did not have an index idx
	// thus for our further sorting we start at index lo + count[0]
	for i := 0; i < len(count)-1; i++ {
		msdsort(array, idx+1, aux, lo+count[i], lo+count[i+1]-1)
	}
}
