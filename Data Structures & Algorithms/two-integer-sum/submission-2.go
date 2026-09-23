func twoSum(nums []int, target int) []int {
    existed := make(map[int]int)
	for i, num := range nums {
		if v, exists := existed[target - num]; exists {
			return []int {v, i}
		}
		existed[num] = i
	}
	return nil
}
