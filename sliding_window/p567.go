package problem

// Permutation in String
// https://leetcode.com/problems/permutation-in-string/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	rs1 := []rune(s1)
	rs2 := []rune(s2)

	mapS1 := make(map[rune]int)
	mapS2 := make(map[rune]int)
	for i := 0; i < len(s1); i++ {
		mapS1[rs1[i]]++
		mapS2[rs2[i]]++
	}

	found := 0
	for ru, co := range mapS2 {
		if _, ok := mapS1[ru]; ok {
			if co >= mapS1[ru] {
				found += mapS1[ru]
			} else {
				found += co
			}
		}
	}

	if found == len(rs1) {
		return true
	}

	l := 0
	for r := len(rs1); r < len(rs2); r++ {
		charL := rs2[l]
		charR := rs2[r]

		mapS2[charR]++
		if countS1, ok := mapS1[charR]; ok && mapS2[charR] <= countS1 {
			found++
		}

		mapS2[charL]--
		if countS1, ok := mapS1[charL]; ok && mapS2[charL] < countS1 {
			found--
		}

		if found == len(rs1) {
			return true
		}

		l++
	}

	return false
}
