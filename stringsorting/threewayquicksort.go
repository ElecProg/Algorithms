package stringsorting

// ThreeWayQuickSort sorts an array of strings using three way partitioning
func ThreeWayQuickSort(array []string) {
	threewayquicksort(array, 0, 0, len(array)-1)
}

func threewayquicksort(array []string, idx int, lo int, hi int) {
	if hi <= lo {
		return
	}

	p := getchr(array[lo], idx)

	l := lo
	r := hi
	i := l + 1

	for i <= r {
		chr := getchr(array[i], idx)

		if chr < p {
			array[l], array[i] = array[i], array[l]
			l++
			i++

		} else if chr > p {
			array[r], array[i] = array[i], array[r]
			r--

		} else {
			i++
		}
	}

	threewayquicksort(array, idx, lo, l-1)

	if p >= 0 {
		// If the pivot wasn't a string without index idx:
		threewayquicksort(array, idx+1, l, r)
	}

	threewayquicksort(array, idx, r+1, hi)
}
