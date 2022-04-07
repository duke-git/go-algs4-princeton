package queue

import "go-algs4-princeton/linkedlist"

type LinkedQueue struct {
	First *linkedlist.Node
	Last  *linkedlist.Node
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{First: nil, Last: nil}
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.First == nil
}

func (q *LinkedQueue) Enqueue(item string) {
	oldLast := q.Last
	last := linkedlist.NewNode(item)
	last.Next = nil
	if q.IsEmpty() {
		q.First = last
	} else {
		oldLast.Next = last
	}
}

func (q *LinkedQueue) Dequeue() string {
	if q.IsEmpty() {
		return ""
	}

	item := q.First.Item
	q.First = q.First.Next
	return item
}
