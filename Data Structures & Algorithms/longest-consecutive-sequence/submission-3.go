func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	result := 0
	for k, _ := range set {
		if _, exists := set[k-1]; !exists {		
			curr := k
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
	}
	return result
}