package threesum

import (
	"go-algs4-princeton/constraint"
	"go-algs4-princeton/sort"
)

// BruteForceCount O(N) = N^3
func BruteForceCount(numbers []int) int {
	var count int
	length := len(numbers)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if numbers[i]+numbers[j]+numbers[k] == 0 {
					count++
				}
			}
		}
	}

	return count
}

// SortBasedCount first sort, then count. O(N) = N^2 * logN
func SortBasedCount(numbers []int) int {
	length := len(numbers)
	sortedNumbers := sort.QuickSort(numbers, 0, length-1)
	var count int

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			target := 0 - (numbers[i] + numbers[j])
			index := BinarySearch(sortedNumbers, target, 0, length-1)
			if index > j {
				count++
			}
		}
	}

	return count

}

func BinarySearch[T constraint.Number](slice []T, target T, lowIndex, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex

	var midIndex int
	for startIndex <= endIndex {
		midIndex = int(startIndex + (endIndex-startIndex)/2)
		if slice[midIndex] > target {
			endIndex = midIndex - 1
		} else if slice[midIndex] < target {
			startIndex = midIndex + 1
		} else {
			return midIndex
		}
	}
	return -1
}
