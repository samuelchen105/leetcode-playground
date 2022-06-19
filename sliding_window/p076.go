package problem

import "math"

// Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	letterMap := make(map[rune]int)
	for _, letter := range t {
		letterMap[letter]++
	}

	answer := []rune{}
	answerLen := math.MaxInt

	input := []rune(s)
	inputLen := len(input)

	letterCount := len(t)
	letterFoundMap := make(map[rune]int)
	letterFound := 0
	left, right := 0, 0

	for ; right < inputLen; right++ {
		// grow window
		letter := input[right]

		max, ok := letterMap[letter]
		if ok {
			letterFoundMap[letter]++
			if letterFoundMap[letter] <= max {
				letterFound++
			}

			// if found all, shrink window for minimum answer
			for letterFound == letterCount {
				if right-left < answerLen {
					answer = input[left : right+1]
					answerLen = right - left
				}

				if left >= right {
					break
				}

				victim := input[left]
				letterFoundMap[victim]--
				if letterFoundMap[victim] < letterMap[victim] {
					letterFound--
				}
				left++

				for {
					if _, ok := letterMap[input[left]]; ok {
						break
					}
					left++
				}
			}
		} else if letterFound == 0 {
			left++
		}
	}

	return string(answer)
}
