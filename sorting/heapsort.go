package sorting

import "github.com/ElecProg/Algorithms/datastructures"

// HeapSort sorts integers using a heap
func HeapSort(array []int) {
	queue := datastructures.PriorityQueue{}

	for _, val := range array {
		queue.Enqueue(val)
	}

	for i := len(array) - 1; i >= 0; i-- {
		array[i], _ = queue.Dequeue()
	}
}
