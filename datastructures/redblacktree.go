package datastructures

// RBTree is a red black tree
type RBTree struct {
	root *rbTreeNode
}

type rbTreeNode struct {
	key   int
	value interface{}
	left  *rbTreeNode
	right *rbTreeNode
	red   bool
}

// Empty check if tree is empty
func (tree *RBTree) Empty() bool {
	return tree.root == nil
}

// Put adds key and value to tree
func (tree *RBTree) Put(key int, value interface{}) {
	tree.root = tree.put(tree.root, key, value)
	tree.root.red = false
}

// Put adds key and value to tree
func (tree *RBTree) put(node *rbTreeNode, key int, value interface{}) *rbTreeNode {
	if node == nil {
		node := rbTreeNode{
			key:   key,
			value: value,
			red:   true,
		}

		return &node
	}

	if key < node.key {
		node.left = tree.put(node.left, key, value)

	} else if key > node.key {
		node.right = tree.put(node.right, key, value)

	} else {
		node.value = value
	}

	if isRed(node.right) && !isRed(node.left) {
		node = rotateLeft(node)
	}

	if isRed(node.right) && isRed(node.left.left) {
		node = rotateRight(node)
	}

	if isRed(node.left) && isRed(node.right) {
		flipColors(node)
	}

	return node
}

func isRed(node *rbTreeNode) bool {
	if node == nil {
		return false
	}

	return node.red
}

func flipColors(node *rbTreeNode) {
	node.red = true
	node.left.red = false
	node.right.red = false
}

func rotateLeft(node *rbTreeNode) *rbTreeNode {
	r := node.right
	node.right = r.left
	r.left = node
	r.red = node.red
	node.red = true
	return r
}

func rotateRight(node *rbTreeNode) *rbTreeNode {
	l := node.left
	node.left = l.right
	l.right = node
	l.red = node.red
	node.red = true
	return l
}
