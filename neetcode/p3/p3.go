package problem

// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	in := []rune(s)
	charSet := make(map[rune]struct{})
	l := 0
	res := 0

	for r := range in {
		charR := in[r]

		for inSet(charSet, charR) {
			delete(charSet, in[l])
			l++
		}

		charSet[charR] = struct{}{}

		windowSize := r - l + 1
		if windowSize > res {
			res = windowSize
		}
	}

	return res
}

func inSet(m map[rune]struct{}, key rune) bool {
	_, ok := m[key]
	return ok
}
