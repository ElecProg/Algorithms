package datastructures

// Stack is a stack
type Stack struct {
	top *stackNode
}

type stackNode struct {
	next  *stackNode
	value interface{}
}

// Empty see if stack is empty
func (stack Stack) Empty() bool {
	return stack.top == nil
}

// Push value on stack
func (stack Stack) Push(value interface{}) {
	stack.top = &stackNode{
		next:  stack.top,
		value: value,
	}
}

// Pop value from stack
func (stack Stack) Pop() (interface{}, bool) {
	if stack.Empty() {
		return nil, false
	}

	node := stack.top
	stack.top = stack.top.next
	return node.value, true
}
