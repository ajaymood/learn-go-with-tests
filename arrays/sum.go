package arrays

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// take in any number of slices
func SumAll(numberSlices ...[]int) []int {
	// slice size should equal number of input slices
	// resultSliceSize := len(numberSlices)
	// sums := make([]int, resultSliceSize)
	var sums []int
	// a for loop for the number of slices
	for _, numbers := range numberSlices {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numberSlices ...[]int) []int {
	if len(numberSlices) < 1 {
		return []int{}
	}
	var sums []int
	// a for loop for the number of slices
	for _, numbers := range numberSlices {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}

	return sums
}
