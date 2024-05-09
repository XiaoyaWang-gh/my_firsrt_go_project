package utils

func PrefixSum(nums ...int) []int {
	sums := make([]int, len(nums))

	if len(sums) == 0 {
		return sums
	}
	sums[0] = nums[0]
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	return sums
}
