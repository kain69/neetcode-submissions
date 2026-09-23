import "slices"

func productExceptSelf(nums []int) []int {
	prefix := slices.Repeat([]int{1}, len(nums)) // 1  1  2  8
	suffix := slices.Repeat([]int{1}, len(nums)) // 48 24 6  1
	result := make([]int, len(nums))

	for i := 1; i <= len(nums) - 1; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := 0; i <= len(nums) - 1; i++ {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}