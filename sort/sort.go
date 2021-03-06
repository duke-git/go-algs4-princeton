package sort

import (
	"go-algs4-princeton/constraint"
	"math/rand"
	"time"
)

// SelectionSort O(N) = 1/2 * N^2, n times of swap
func SelectionSort[T constraint.Number](slice []T) {
	for i := 0; i < len(slice); i++ {
		min := i

		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		swap(slice, i, min)
	}
}

// InsertionSort O(N) = 1/4 * N^2, 1/4 * N^2 times of swap
// for partical sorted data, InsertionSort takes n times
func InsertionSort[T constraint.Number](slice []T) {
	for i := 0; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if slice[j] < slice[j-1] {
				swap(slice, j, j-1)
			} else {
				break
			}
		}

	}
}

// ShellSort O(N) = N^2/3
func ShellSort[T constraint.Number](slice []T) {
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
}

// MergeSort O(N) = N * logN
func MergeSort[T constraint.Number](slice []T) {
	s := make([]T, len(slice))
	mergeSort(slice, s, 0, len(slice)-1)
}

func mergeSort[T constraint.Number](partSortedSlice []T, slice []T, low, high int) {
	if high <= low {
		return
	}
	mid := low + (high-low)/2
	mergeSort(partSortedSlice, slice, low, mid)
	mergeSort(partSortedSlice, slice, mid+1, high)
	merge(partSortedSlice, slice, low, mid, high)
}

func merge[T constraint.Number](partSortedSlice []T, slice []T, low, mid, high int) {
	for k := low; k <= high; k++ {
		slice[k] = partSortedSlice[k]
	}

	i := low
	j := mid + 1

	for k := low; k <= high; k++ {
		if i > mid {
			partSortedSlice[k] = slice[j]
			j++
		} else if j > high {
			partSortedSlice[k] = slice[i]
			i++
		} else if slice[j] < slice[i] {
			partSortedSlice[k] = slice[j]
			j++
		} else {
			partSortedSlice[k] = slice[i]
			i++
		}
	}

}

// QuickSort O(N) = N * logN
func QuickSort[T constraint.Number](slice []T) {
	// quickSort2(slice, 0, len(slice)-1)
	threeWayQuickSort(slice, 0, len(slice)-1)
}

func quickSort1[T constraint.Number](slice []T, low, high int) {
	if low > high {
		return
	}

	i := low
	j := high

	for i != j {
		for slice[j] >= slice[low] && j > i {
			j--
		}
		for slice[i] <= slice[low] && j > i {
			i++
		}

		if i < j {
			swap(slice, i, j)
		}
	}

	swap(slice, low, i)

	quickSort1(slice, low, i-1)
	quickSort1(slice, i+1, high)
}

func quickSort2[T constraint.Number](slice []T, low, high int) {
	if high <= low {
		return
	}
	j := partition(slice, low, high)
	quickSort2(slice, low, j-1)
	quickSort2(slice, j+1, high)
}

// // partition split slice into two parts
func partition[T constraint.Number](slice []T, low, high int) int {
	i := low
	highVal := slice[high]

	for j := low; j < high; j++ {
		if slice[j] < highVal {
			swap(slice, i, j)
			i++
		}
	}

	swap(slice, i, high)

	return i
}

func threeWayQuickSort[T constraint.Number](slice []T, low, high int) {
	if high <= low {
		return
	}

	lt := low
	gt := high
	val := slice[low]

	i := low

	for i <= gt {
		if slice[i] < val {
			swap(slice, lt, i)
			lt++
			i++
		} else if slice[i] > val {
			swap(slice, i, gt)
			gt--
		} else {
			i++
		}
	}

	threeWayQuickSort(slice, low, lt-1)
	threeWayQuickSort(slice, gt+1, high)
}

func swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func Shuffle(slice []int) {
	for i := 0; i < len(slice); i++ {
		r := RandInt(0, i+1)
		swap(slice, i, r)
	}
}
func RandInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
