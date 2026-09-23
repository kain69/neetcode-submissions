func isAnagram(s string, t string) bool {
	chars := make(map[rune]int, len(s))
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
	if len(chars) > 0 {
		return false
	}
	return true
}
