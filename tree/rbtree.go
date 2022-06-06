package tree

import "go-algs4-princeton/constraint"

type Color int

const (
	BLACK Color = iota //  0
	RED                // 1
)

// RBTree is left leaning red-black tree
type RBTree[T constraint.Comparable, U any] struct {
	root *Node[T, U]
}

func (t *RBTree[T, U]) isRed(node *Node[T, U]) bool {
	if node == nil {
		return false
	}

	return node.color == RED
}

func (t *RBTree[T, U]) rotateLeft(node *Node[T, U]) *Node[T, U] {
	if !(t.isRed(node.right)) {
		return nil
	}

	temp := node.right
	node.right = temp.left
	temp.left = node
	temp.color = node.color
	node.color = RED

	return temp
}

func (t *RBTree[T, U]) rotateRight(node *Node[T, U]) *Node[T, U] {
	if !(t.isRed(node.left)) {
		return nil
	}

	temp := node.left
	node.left = temp.right
	temp.right = node
	temp.color = node.color
	node.color = RED

	return temp
}

func (t *RBTree[T, U]) flipColors(node *Node[T, U]) {
	if t.isRed(node) {
		return
	}
	if !(t.isRed(node.left)) {
		return
	}
	if !(t.isRed(node.right)) {
		return
	}

	node.color = RED
	node.left.color = BLACK
	node.right.color = BLACK
}

func (t *RBTree[T, U]) Put(key T, val U) {
	t.root = t.put(t.root, key, val)
}

func (t *RBTree[T, U]) put(node *Node[T, U], key T, val U) *Node[T, U] {
	if node == nil {
		return NewNode(key, val)
	}

	if key < node.key {
		node.left = t.put(node.left, key, val)
	} else if key > node.key {
		node.right = t.put(node.right, key, val)
	} else if key == node.key {
		node.val = val
	}

	if t.isRed(node.right) && !t.isRed(node.left) {
		node = t.rotateLeft(node)
	}
	if t.isRed(node.left) && !t.isRed(node.left.left) {
		node = t.rotateRight(node)
	}

	if t.isRed(node.right) && t.isRed(node.left) {
		t.flipColors(node)
	}

	node.count = size(node.left) + size(node.right) + 1
	return node
}

func (t *RBTree[T, U]) Get(key T) U {
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

// func (t *RBTree[T, U]) Delete(key T) {
// 	t.root = delete(t.root, key)
// }

// func delete[T constraint.Comparable, U any](node *Node[T, U], key T) *Node[T, U] {
// 	if node == nil {
// 		return nil
// 	}

// 	if key < node.key {
// 		node.left = delete(node.left, key)
// 	} else if key > node.key {
// 		node.right = delete(node.right, key)
// 	} else {
// 		if node.right == nil {
// 			return node.left
// 		}
// 		if node.left == nil {
// 			return node.right
// 		}

// 		x := node
// 		// node = min(x.right)
// 		node.right = deleteMin(node.right)
// 		node.left = x.left
// 	}
// 	node.count = size(node.left) + size(node.right) + 1

// 	return node
// }

// func deleteMin[T constraint.Comparable, U any](node *Node[T, U]) *Node[T, U] {
// 	if node.left == nil {
// 		return node.right
// 	}

// 	node.left = deleteMin(node.left)
// 	node.count = size(node.left) + size(node.right) + 1
// 	return node
// }

// func (t *BSTree[T, U]) Floor(key T) T {
// 	node := floor(t.root, key)
// 	if node == nil {
// 		var key T
// 		return key
// 	}

// 	return node.key
// }

// func floor[T constraint.Comparable, U any](node *Node[T, U], key T) *Node[T, U] {
// 	if node == nil {
// 		return nil
// 	}

// 	if key == node.key {
// 		return node
// 	} else if key < node.key {
// 		return floor(node.left, key)
// 	}

// 	n := floor(node.right, key)
// 	if n != nil {
// 		return n
// 	}
// 	return node
// }

func (t *RBTree[T, U]) Size() int {
	return size(t.root)
}

// func size[T constraint.Comparable, U any](node *Node[T, U]) int {
// 	if node == nil {
// 		return 0
// 	}
// 	return node.count
// }
