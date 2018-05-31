package sorting

// MergeSort ~ n lg n
func MergeSort(array []int) {
	aux := make([]int, len(array))

	mergesort(array, aux, 0, len(array)-1)
}

func mergesort(array []int, aux []int, lo int, hi int) {
	if hi <= lo {
		return
	}

	mid := (lo + hi) / 2
	mergesort(array, aux, lo, mid)
	mergesort(array, aux, mid+1, hi)

	for i := lo; i <= hi; i++ {
		aux[i] = array[i]
	}

	l := lo
	r := mid + 1

	for i := lo; i <= hi; i++ {
		if l > mid {
			array[i] = aux[r]
			r++

		} else if r > hi {
			array[i] = aux[l]
			l++

		} else if aux[l] <= aux[r] {
			array[i] = aux[l]
			l++

		} else {
			array[i] = aux[r]
			r++
		}
	}
}
