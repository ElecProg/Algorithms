package datastructures

// LinkedList is a signly linked list
type LinkedList struct {
	first  *linkedListNode
	last   *linkedListNode
	length int
}

type linkedListNode struct {
	next  *linkedListNode
	value interface{}
}

// Empty see if the linked list is empty
func (list *LinkedList) Empty() bool {
	return list.length == 0
}

// Infront adds an element to the beginning of the list
func (list *LinkedList) Infront(value interface{}) {
	node := &linkedListNode{
		next:  list.first,
		value: value,
	}

	if list.Empty() {
		list.last = node
	}

	list.first = node
	list.length++
}

// Append an element to the end of the list
func (list *LinkedList) Append(value interface{}) {
	node := &linkedListNode{
		next:  nil,
		value: value,
	}

	if list.Empty() {
		list.first = node
	}

	list.last = node
	list.length++
}

// Contains sees if a value is contained within the list
func (list *LinkedList) Contains(value interface{}) bool {
	current := list.first

	for current != nil {
		if current.value == value {
			return true
		}

		current = current.next
	}

	return false
}
