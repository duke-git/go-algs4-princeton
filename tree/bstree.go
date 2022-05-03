package tree

import "go-algs4-princeton/constraint"

type BSTree[T constraint.Comparable, U any] struct {
	root *Node[T, U]
}

func (t *BSTree[T, U]) Put(key T, val U) {
	t.root = put(t.root, key, val)
}

func put[T constraint.Comparable, U any](node *Node[T, U], key T, val U) *Node[T, U] {
	if node == nil {
		return NewNode(key, val)
	}

	if key < node.key {
		node.left = put(node.left, key, val)
	} else if key > node.key {
		node.right = put(node.right, key, val)
	} else if key == node.key {
		node.val = val
	}

	return node
}

func (t *BSTree[T, U]) Get(key T) U {
	current := t.root

	for current != nil {
		if key < current.key {
			current = current.left
		} else if key > current.key {
			current = current.left
		} else {
			return current.val
		}
	}
	var val U
	return val
}

func (t *BSTree[T, U]) Delete(key T) {

}
