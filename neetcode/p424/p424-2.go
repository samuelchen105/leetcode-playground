package problem

// 424. Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/

func characterReplacement2(s string, k int) int {
	res := 0
	charMap := make(map[rune]int)
	in := []rune(s)
	maxFreq := 0

	l := 0
	for r := 0; r < len(in); r++ {
		charMap[in[r]]++
		if charMap[in[r]] > maxFreq {
			maxFreq = charMap[in[r]]
		}

		for r-l+1-maxFreq > k {
			charMap[in[l]]--
			l++
		}

		windowSize := r - l + 1
		if windowSize > res {
			res = windowSize
		}
	}

	return res
}
