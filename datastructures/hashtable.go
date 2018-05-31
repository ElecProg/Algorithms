package datastructures

// This file implements a hashtable using linear probing.
// Note: thombstones are represented by a hastableNode with nil key and value.
// Todo: Currently we probe until the end of the array, it would be better to wrap around.

// Hashable types are types that can be used as a key in a hash table
type Hashable interface {
	Hash() int
}

// Hashtable is a hash table
type Hashtable struct {
	table    []*hashtableNode
	elements int
}

type hashtableNode struct {
	key   Hashable
	value interface{}
}

// Empty check if hastable is empty
func (table *Hashtable) Empty() bool {
	return table.elements == 0
}

// Get tries to find a key in the hash table
func (table *Hashtable) Get(key Hashable) (value interface{}, ok bool) {
	if key == nil || table.elements == 0 {
		return nil, false
	}

	for i := key.Hash() % table.elements; i < len(table.table); i++ {
		node := table.table[i]

		if node == nil {
			// Element is not in the hashtable
			return nil, false
		}

		if node.key == key {
			return node.value, true
		}
	}

	return nil, false
}

// Put adds key and value to the hash table
func (table *Hashtable) Put(key Hashable, value interface{}) (ok bool) {
	if key == nil {
		return false
	}

	node := &hashtableNode{
		key:   key,
		value: value,
	}

	table.add(node)
	return true
}

func (table *Hashtable) add(node *hashtableNode) {
	if node == nil || node.key == nil {
		// Can't add nil node nor node with nil key (thombstone)
		return
	}

	if table.elements == 0 {
		table.table = make([]*hashtableNode, 2)
	}

	if table.elements == len(table.table)/2 {
		table.grow()
	}

	for {
		for i := node.key.Hash() % table.elements; i < len(table.table); i++ {
			if table.table[i] == nil {
				table.table[i] = node
				table.elements++
				return
			}
		}

		// No space found
		table.grow()
	}
}

func (table *Hashtable) grow() {
	oldtable := table.table
	table.table = make([]*hashtableNode, len(oldtable)*2)
	table.elements = 0

	for _, node := range oldtable {
		table.add(node)
	}
}

// Remove key from the hash table
func (table *Hashtable) Remove(key Hashable) (ok bool) {
	if key == nil || table.elements == 0 {
		return false
	}

	for i := key.Hash() % table.elements; i < len(table.table); i++ {
		node := table.table[i]

		if node == nil {
			// Element is not in the hash table
			return false
		}

		if node.key == key {
			node.key = nil
			node.value = nil
			table.elements--

			if table.elements <= len(table.table)/4 {
				table.shrink()
			}

			return true
		}
	}

	return false
}

func (table *Hashtable) shrink() {
	oldtable := table.table
	table.table = make([]*hashtableNode, len(oldtable)/2)
	table.elements = 0

	for _, node := range oldtable {
		table.add(node)
	}
}
