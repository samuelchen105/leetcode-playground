package problem

// Permutation in String
// https://leetcode.com/problems/permutation-in-string/

func checkInclusion(s1 string, s2 string) bool {
	charMap := make(map[rune]int)
	for _, r := range s1 {
		charMap[r]++
	}
	threshold := len(s1)
	windowChar := make(map[rune]int)
	charFound := 0

	runes := []rune(s2)
	l := 0
	exceed := false
	for r, char := range runes {
		maxCount, ok := charMap[char]
		if !ok {
			l = r + 1
			charFound = 0
			windowChar = make(map[rune]int)
			exceed = false
			continue
		}

		for exceed && l <= r {
			charL := runes[l]
			windowChar[charL]--
			l++
			if windowChar[charL] >= charMap[charL] {
				exceed = false
			} else {
				charFound--
			}
		}

		windowChar[char]++
		if windowChar[char] <= maxCount {
			charFound++
		} else {
			exceed = true
		}

		if charFound == threshold {
			return true
		}
	}

	return false
}
