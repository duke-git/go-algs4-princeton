package tree

import "go-algs4-princeton/constraint"

type BSTree[T constraint.Comparable, U any] struct {
	root *Node[T, U]
}

func (t *BSTree[T, U]) Put(key T, val U) {

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
