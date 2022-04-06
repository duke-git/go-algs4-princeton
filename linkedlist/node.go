package linkedlist

type Node struct {
	Item string
	Next *Node
}

func NewNode(item string) *Node {
	return &Node{Item: item, Next: nil}
}
