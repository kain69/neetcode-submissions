func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	slice := make([]int, 0, len(nums))
	for k, _ := range set {
		if _, exists := set[k-1]; !exists {
			slice = append(slice, k)
		} 
	}

	result := 0
	for _, num := range slice {
		curr := num
		temp := 1
		for {
			if _, exists := set[curr + 1]; exists {
				temp++
				curr++
			} else {
				if result < temp {
					result = temp
				}
				break
			}
		}
	}
	return result
}