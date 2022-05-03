package tree

import (
	"go-algs4-princeton/constraint"
)

type Node[T constraint.Comparable, U any] struct {
	key   T
	val   U
	left  *Node[T, U]
	right *Node[T, U]
}

func NewNode[T constraint.Comparable, U any](key T, val U) *Node[T, U] {
	return &Node[T, U]{key, val, nil, nil}
}
