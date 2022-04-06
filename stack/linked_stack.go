package stack

import "go-algs4-princeton/linkedlist"

type LinkedStack struct {
	First *linkedlist.Node
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{First: nil}
}

func (s *LinkedStack) IsEmpty() bool {
	return s.First == nil
}

func (s *LinkedStack) Push(item string) {
	oldFirst := s.First
	first := linkedlist.NewNode(item)
	first.Next = oldFirst
}

func (s *LinkedStack) Pop() string {
	item := s.First.Item
	s.First = s.First.Next
	return item
}
