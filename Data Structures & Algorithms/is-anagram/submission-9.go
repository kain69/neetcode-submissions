func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	chars := make(map[rune]int)
	for _, ch := range s {
		chars[ch] += 1
	}

	for _, ch := range t {
		if chars[ch] == 0 {
			return false
		}
		chars[ch] -= 1
		if chars[ch] == 0 {
			delete(chars, ch)
		}
	}
	return true
}
