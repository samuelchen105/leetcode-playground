package problem

// 424. Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/

func characterReplacement(s string, k int) int {
	res := 0
	charMap := make(map[rune]int)
	in := []rune(s)

	l := 0
	for r := 0; r < len(in); r++ {
		charMap[in[r]]++

		for r-l+1-getMaxFreq(charMap) > k {
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

func getMaxFreq(charMap map[rune]int) int {
	ret := 0
	for _, v := range charMap {
		if v > ret {
			ret = v
		}
	}
	return ret
}
