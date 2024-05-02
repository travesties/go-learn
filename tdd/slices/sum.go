package slices

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numberSets ...[]int) []int {
	sums := make([]int, len(numberSets))

	for i, nums := range numberSets {
		sums[i] = Sum(nums)
	}

	return sums
}

func SumAllTails(numberSets ...[]int) []int {
	sums := make([]int, len(numberSets))

	for i, nums := range numberSets {
		if len(nums) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(nums[1:])
		}
	}

	return sums
}
