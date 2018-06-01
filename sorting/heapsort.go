package sorting

import "github.com/ElecProg/Algorithms/datastructures"

// HeapSort sorts integers using a heap
func HeapSort(array []int) {
	queue := datastructures.BuildPriorityQueue(array)

	for i := len(array) - 1; i >= 0; i-- {
		array[i], _ = queue.Dequeue()
	}
}
