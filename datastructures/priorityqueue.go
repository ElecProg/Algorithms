package datastructures

// PriorityQueue is a max priority queue of integers
type PriorityQueue struct {
	heap []int
	size int
}

// BuildPriorityQueue creates a priority queue from an array
func BuildPriorityQueue(array []int) PriorityQueue {
	queue := PriorityQueue{
		heap: make([]int, len(array)*2),
		size: len(array),
	}

	// Copy array to heap
	for i, val := range array {
		queue.heap[i+1] = val
	}

	// Construct heap
	for i := queue.size; i > 0; i-- {
		queue.swim(i)
	}

	return queue
}

// Empty checks if the priority queue is empty
func (queue *PriorityQueue) Empty() bool {
	return queue.size == 0
}

// Enqueue adds element to queue
func (queue *PriorityQueue) Enqueue(value int) {
	// Grow queue
	if queue.size+2 >= len(queue.heap) {
		newHeap := make([]int, 2*queue.size+2)

		// Copy to new heap
		for i, val := range queue.heap {
			newHeap[i] = val
		}

		queue.heap = newHeap
	}

	queue.size++
	queue.heap[queue.size] = value
	queue.swim(queue.size)
}

// Dequeue removes the top element from the queue
func (queue *PriorityQueue) Dequeue() (value int, ok bool) {
	if queue.Empty() {
		return 0, false
	}

	// Shrink queue
	if queue.size <= len(queue.heap)/4 {
		newHeap := make([]int, len(queue.heap)/2)

		// Copy to new heap
		for i := 0; i <= queue.size; i++ {
			newHeap[i] = queue.heap[i]
		}

		queue.heap = newHeap
	}

	result := queue.heap[1]
	queue.heap[1] = queue.heap[queue.size]
	queue.size--
	queue.sink(1)

	return result, true
}

func (queue *PriorityQueue) swim(k int) {
	for k > 1 && queue.heap[k/2] < queue.heap[k] {
		queue.heap[k/2], queue.heap[k] = queue.heap[k], queue.heap[k/2]
		k = k / 2
	}
}

func (queue *PriorityQueue) sink(k int) {
	for 2*k <= queue.size {
		j := 2 * k

		if j < queue.size && queue.heap[j] < queue.heap[j+1] {
			j++
		}

		if queue.heap[j] < queue.heap[k] {
			break
		}

		queue.heap[j], queue.heap[k] = queue.heap[k], queue.heap[j]
		k = j
	}
}
