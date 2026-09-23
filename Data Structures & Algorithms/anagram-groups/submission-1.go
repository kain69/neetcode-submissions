func groupAnagrams(strs []string) [][]string {
	group := make(map[[26]int][]string)
	for _, str := range strs {
		freq := [26]int{}
		for _, ch := range str {
			freq[ch-'a'] += 1
		}
		group[freq] = append(group[freq], str)
	}

	var result [][]string
	for _, arr := range group {
		result = append(result, arr)
	}
	return result
}
