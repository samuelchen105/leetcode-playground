package problem

// Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	result := 0
	start := 0
	m := map[rune]int{}

	for i, r := range s {
		if idx, ok := m[r]; ok {
			start = max(start, idx+1)
		}
		m[r] = i

		l := i - start + 1
		if l > result {
			result = l
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
