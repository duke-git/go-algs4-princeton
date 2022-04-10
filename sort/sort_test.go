package sort

import (
	"go-algs4-princeton/internal"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	assert := internal.NewAssert(t, "TestSelectionSort")

	numbers := []int{5, 4, 3, 2, 1}

	SelectionSort(numbers)

	assert.Equal([]int{1, 2, 3, 4, 5}, numbers)
}

func TestInsertionSort(t *testing.T) {
	assert := internal.NewAssert(t, "TestInsertionSort")

	numbers := []int{5, 4, 3, 2, 1}

	InsertionSort(numbers)

	assert.Equal([]int{1, 2, 3, 4, 5}, numbers)
}

func TestShellSort(t *testing.T) {
	assert := internal.NewAssert(t, "TestShellSort")

	numbers := []int{5, 4, 3, 2, 1}

	ShellSort(numbers)

	assert.Equal([]int{1, 2, 3, 4, 5}, numbers)
}

func TestShuffle(t *testing.T) {
	// assert := internal.NewAssert(t, "TestShuffle")

	numbers := []int{5, 4, 3, 2, 1}
	Shuffle(numbers)
	t.Log(numbers)
}
