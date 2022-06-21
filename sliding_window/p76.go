package problem

// Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	mapT := make(map[rune]int)
	for _, r := range t {
		mapT[r]++
	}

	answer := struct {
		length int
		start  int
		end    int
	}{
		length: -1,
	}

	type Indice struct {
		idx int
		val rune
	}

	filteredS := []Indice{}
	for i, r := range s {
		if _, ok := mapT[r]; ok {
			filteredS = append(filteredS, Indice{i, r})
		}
	}

	foundMap := make(map[rune]int)
	found := 0
	left, right := 0, 0

	for right < len(filteredS) {
		indiceR := filteredS[right]

		foundMap[indiceR.val]++
		if foundMap[indiceR.val] <= mapT[indiceR.val] {
			found++
		}

		for left <= right && found == len(t) {
			indiceL := filteredS[left]

			newAnswerLen := indiceR.idx - indiceL.idx + 1
			if answer.length == -1 || newAnswerLen < answer.length {
				answer.length = newAnswerLen
				answer.start = indiceL.idx
				answer.end = indiceR.idx
			}

			foundMap[indiceL.val]--
			if foundMap[indiceL.val] < mapT[indiceL.val] {
				found--
			}
			left++
		}
		right++
	}

	if answer.length == -1 {
		return ""
	}
	return string([]rune(s)[answer.start : answer.end+1])
}
