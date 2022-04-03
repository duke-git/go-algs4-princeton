package threesum

import (
	"go-algs4-princeton/internal"
	"testing"
)

func TestBruteForceCount(t *testing.T) {
	assert := internal.NewAssert(t, "TestBruteForceCount")
	input := []int{30, -40, -20, -10, 40, 0, 10, 5}

	count := BruteForceCount(input)
	assert.Equal(4, count)
}

func BenchmarkBruteForceCount(b *testing.B) {
	input := []int{30, -40, -20, -10, 40, 0, 10, 5}
	var count int

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count = BruteForceCount(input)
	}

	b.StopTimer()

	b.Log(count)
}

func TestSortBasedCount(t *testing.T) {
	assert := internal.NewAssert(t, "TestSortBasedCount")
	input := []int{30, -40, -20, -10, 40, 0, 10, 5}

	count := SortBasedCount(input)
	assert.Equal(4, count)
}

func BenchmarkSortBasedCount(b *testing.B) {
	input := []int{30, -40, -20, -10, 40, 0, 10, 5}
	var count int

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count = SortBasedCount(input)
	}

	b.StopTimer()

	b.Log(count)
}
