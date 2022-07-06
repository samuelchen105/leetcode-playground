package problem

// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	mapS := make(map[rune]int)
	for _, r := range s {
		mapS[r]++
	}
	count := len([]rune(s))

	for _, r := range t {
		if c, ok := mapS[r]; ok && c > 0 {
			mapS[r]--
			count--
		} else {
			return false
		}
	}

	return count == 0
}
