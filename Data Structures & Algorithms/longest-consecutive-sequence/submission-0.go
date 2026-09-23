func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	maxLen := 0

	for _, num := range nums {
		if _, ok := m[num-1]; !ok {
			currentNum := num
			currentLen := 1
			
			for m[currentNum+1] {
				currentNum++
				currentLen++
			}

			if currentLen > maxLen {
				maxLen = currentLen
			}
		}
	}

	return maxLen
}