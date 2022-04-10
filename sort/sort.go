package sort

import (
	"go-algs4-princeton/constraint"
	"math/rand"
	"time"
)

// SelectionSort O(N) = 1/2 * N^2, n times of swap
func SelectionSort[T constraint.Number](slice []T) []T {
	for i := 0; i < len(slice); i++ {
		min := i

		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		swap(slice, i, min)
	}

	return slice
}

// InsertionSort O(N) = 1/4 * N^2, 1/4 * N^2 times of swap
// for partical sorted data, InsertionSort takes n times
func InsertionSort[T constraint.Number](slice []T) []T {
	for i := 0; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if slice[j] < slice[j-1] {
				swap(slice, j, j-1)
			} else {
				break
			}
		}

	}

	return slice
}

// ShellSort O(N) = N^2/3
func ShellSort[T constraint.Number](slice []T) []T {
	size := len(slice)

	h := 1
	for h < size/3 {
		h = 3*h + 1 //1, 4, 13
	}

	for h >= 1 {
		for i := h; i < size; i++ {
			for j := i; j >= h && slice[j] < slice[j-h]; j = j - h {
				swap(slice, j, j-h)
			}
		}
		h = h / 3
	}

	return slice
}

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


func Shuffle(slice []int)  {
	for i := 0; i < len(slice); i++ {
		r := RandInt(0, i+1)
		swap(slice, i, r)
	}
}
func RandInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
