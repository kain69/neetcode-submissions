func isPalindrome(s string) bool {
	runes := []rune(s)
	l, r := 0, len(runes)-1

	for l < r {
		for l < r && !isLetterOrDigit(runes[l]) {
			l++
		}
		for r > l && !isLetterOrDigit(runes[r]) {
			r--
		}
		if unicode.ToLower(runes[l]) != unicode.ToLower(runes[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func isLetterOrDigit(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}