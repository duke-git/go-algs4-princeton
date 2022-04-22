package queue

import "go-algs4-princeton/constraint"

// PriorityHeapQueue
type PriorityHeapQueue[T constraint.Number] struct {
	items []T
	size  int
}

func NewPriorityHeapQueue[T constraint.Number](capacity int) *PriorityHeapQueue[T] {
	return &PriorityHeapQueue[T]{
		// items: make([]T, capacity+1, capacity+1),
		items: make([]T, capacity+1),
		size:  0,
	}
}

// Data return data
func (q *PriorityHeapQueue[T]) Data() []T {
	data := make([]T, q.size)
	for i := 1; i < q.size+1; i++ {
		data[i-1] = q.items[i]
	}

	return data
}

func (q *PriorityHeapQueue[T]) IsEmpty() bool {
	return q.size == 0
}

// Insert value into heap  O(N) = N * logN
func (q *PriorityHeapQueue[T]) Insert(val T) {
	q.size++
	q.items[q.size] = val
	q.swim(q.size)
}

// DeleteMax delete max value in heap O(N) = N * logN
func (q *PriorityHeapQueue[T]) DeleteMax() T {
	max := q.items[1]

	q.swap(1, q.size)
	q.size--

	q.sink(1)

	//set zero value of T
	var val T
	q.items[q.size+1] = val

	return max
}

// swim when child's key larger than parent's key, exchange them.
func (q *PriorityHeapQueue[T]) swim(index int) {
	for index > 1 && q.less(index/2, index) {
		q.swap(index, index/2)
		index = index / 2
	}
}

// sink when parent's key smaller than child's key, exchange parent's key with larger child's key.
func (q *PriorityHeapQueue[T]) sink(index int) {

	for 2*index <= q.size {
		j := 2 * index

		// get larger child node index
		if j < q.size && q.less(j, j+1) {
			j++
		}
		// if parent larger than child, stop
		if !(q.less(index, j)) {
			break
		}

		q.swap(index, j)
		index = j
	}
}

func (q *PriorityHeapQueue[T]) less(i, j int) bool {
	return q.items[i] < q.items[j]
}

func (q *PriorityHeapQueue[T]) swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}
