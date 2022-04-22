package queue

import (
	"go-algs4-princeton/internal"
	"testing"
)

func TestPriorityHeapQueue(t *testing.T) {
	assert := internal.NewAssert(t, "TestPriorityHeapQueue")

	phq := NewPriorityHeapQueue[int](10)

	for i := 1; i < 11; i++ {
		phq.Insert(i)
	}

	t.Log(phq.Data())

	expected := []int{10, 9, 6, 7, 8, 2, 5, 1, 4, 3}
	assert.Equal(expected, phq.Data())
}

func TestPriorityHeapQueue_DeleteMax(t *testing.T) {
	assert := internal.NewAssert(t, "TestPriorityHeapQueue_DeleteMax")

	phq := NewPriorityHeapQueue[int](10)

	for i := 1; i < 11; i++ {
		phq.Insert(i)
	}

	assert.Equal(10, phq.DeleteMax())

	expected := []int{9, 8, 6, 7, 3, 2, 5, 1, 4}
	assert.Equal(expected, phq.Data())
	t.Log(phq.Data())
}
