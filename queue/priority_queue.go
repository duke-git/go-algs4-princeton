package queue

// unordered max priority queue
type UnorderedMaxPQ[K comparable] struct {
	items []K
	size  int
}

func NewUnorderedMaxPQ[K comparable](capacity int) *UnorderedMaxPQ[K] {
	return &UnorderedMaxPQ[K]{
		items: make([]K, capacity, capacity),
		size:  0,
	}
}

func (q *UnorderedMaxPQ[K]) IsEmpty() bool {
	return q.size == 0
}

func (q *UnorderedMaxPQ[K]) Insert(val K) {
	q.items[q.size] = val
	q.size++
}

func (q *UnorderedMaxPQ[K]) DeleteMax() K {
	var max int

	for i := 0; i < q.size; i++ {
		if max < i {
			max = i
		}
	}

	q.items[max], q.items[q.size-1] = q.items[q.size-1], q.items[max]

	q.size--

	return q.items[q.size]
}
