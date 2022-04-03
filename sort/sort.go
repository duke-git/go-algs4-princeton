package sort

import (
	"go-algs4-princeton/constraint"
)

func QuickSort[T constraint.Number](slice []T, lowIndex, highIndex int) []T {
	if lowIndex < highIndex {
		p := partition(slice, lowIndex, highIndex)
		QuickSort(slice, lowIndex, p-1)
		QuickSort(slice, p+1, highIndex)
	}

	return slice
}

// partition split slice into two parts
func partition[T constraint.Number](slice []T, lowIndex, highIndex int) int {
	p := slice[highIndex]
	i := lowIndex
	for j := lowIndex; j < highIndex; j++ {
		if slice[j] < p {
			swap(slice, i, j)
			i++
		}
	}

	swap(slice, i, highIndex)

	return i
}

func swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
