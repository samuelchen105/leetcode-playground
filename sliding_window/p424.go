package problem

// Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/

func characterReplacement(s string, k int) int {
	runeCount := make(map[rune]int)
	maxCount := 0
	maxLen := 0
	left, right := 0, 0
	in := []rune(s)
	for right < len(in) {
		runeCount[in[right]]++

		if runeCount[in[right]] > maxCount {
			maxCount = runeCount[in[right]]
		}

		for left < right && right-left+1-maxCount > k {
			runeCount[in[left]]--
			left++
		}

		windowSize := right - left + 1
		if windowSize > maxLen {
			maxLen = windowSize
		}

		right++
	}

	return maxLen
}
