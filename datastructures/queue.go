package datastructures

// Queue is a queue
type Queue struct {
	in  *queueNode
	out *queueNode
}

type queueNode struct {
	next  *queueNode
	value interface{}
}

// Empty see if queue is empty
func (queue Queue) Empty() bool {
	return queue.out == nil
}

// Enqueue value on queue
func (queue Queue) Enqueue(value interface{}) {
	node := &queueNode{
		next:  nil,
		value: value,
	}

	if queue.Empty() {
		queue.in = node
		queue.out = node
		return
	}

	queue.in.next = node
	queue.in = node
}

// Dequeue value from queue
func (queue Queue) Dequeue() (interface{}, bool) {
	if queue.Empty() {
		return nil, false
	}

	node := queue.out
	queue.out = node.next

	if queue.out == nil {
		// We just emptied the queue
		queue.in = nil
	}

	return node.value, true
}
