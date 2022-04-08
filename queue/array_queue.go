package queue

type ArrayQueue struct {
	items    []string
	head     int
	tail     int
	capacity int
	size     int
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		items:    make([]string, 0, capacity),
		head:     0,
		tail:     0,
		capacity: capacity,
		size:     0,
	}
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *ArrayQueue) Enqueue(item string) bool {
	if q.head == 0 && q.tail == q.capacity { //队列满
		return false
	} else if q.head != 0 && q.tail == q.capacity { //扩容
		for i := q.head; i < q.tail; i++ {
			q.items[i-q.head] = q.items[i]
		}
		q.head = 0
		q.tail = q.tail - q.head
	}

	q.items[q.tail] = item
	q.tail++
	q.size++
	return true
}

func (q *ArrayQueue) Dequeue() (string, bool) {
	if q.head == q.tail {
		return "", false
	}
	item := q.items[q.head]
	q.head++
	return item, true
}
