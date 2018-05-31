package sorting

// QuickSort ~ 1.39 n lg(n)
func QuickSort(array []int) {
	quicksort(array, 0, len(array)-1)
}

func quicksort(array []int, lo int, hi int) {
	if hi <= lo {
		return
	}

	p := array[lo]

	l := lo + 1
	r := hi

	for {
		for array[l] < p {
			if l == hi {
				break
			}

			l++
		}

		for array[r] > p {
			if r == lo {
				break
			}

			r--
		}

		if l >= r {
			break
		}

		array[l], array[r] = array[r], array[l]
	}

	array[lo], array[r] = array[r], array[lo]

	quicksort(array, lo, r-1)
	quicksort(array, r+1, hi)
}
