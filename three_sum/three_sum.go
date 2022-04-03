package threesum


// BruteForceCount O(N) = N^3
func BruteForceCount(numbers []int) int {
	var count int
	length := len(numbers)

	for i := 0; i < length; i++ {
		for j := i + 1; j < count; j++ {
			for k := j + 1; k < count; k++ {
				if numbers[i]+numbers[j]+numbers[k] == 0 {
					count++
				}
			}
		}
	}

	return count
}  
