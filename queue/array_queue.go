package queue

type ArrayQueue struct {
	items  []string
	length int
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		items:  make([]string, 0, capacity),
		length: 0,
	}
}
