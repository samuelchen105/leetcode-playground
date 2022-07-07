package problem

// 125. Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
	rs := []rune(s)
	l := 0
	r := len(rs) - 1
	for l < r {
		for l < r && !isAlphaNum(rs[l]) {
			l++
		}

		for l < r && !isAlphaNum(rs[r]) {
			r--
		}

		if toLower(rs[l]) != toLower(rs[r]) {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlphaNum(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return 'a' + r - 'A'
	}
	return r
}
