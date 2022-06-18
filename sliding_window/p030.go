package problem

// Substring with Concatenation of All Words
// https://leetcode.com/problems/substring-with-concatenation-of-all-words/

func findSubstring(s string, words []string) []int {
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	sw := &SlidingWindow{
		str:     s,
		wordMap: wordMap,
		wordLen: len(words[0]),
		wordNum: len(words),
	}

	answer := []int{}
	for i := 0; i < sw.wordLen; i++ {
		answer = append(answer, sw.find(i)...)
	}

	return answer
}

type SlidingWindow struct {
	str     string
	wordMap map[string]int
	wordLen int
	wordNum int
}

func (s *SlidingWindow) find(left int) []int {
	answer := []int{}
	wordFound := make(map[string]int)
	foundNum := 0
	exceeded := false

	n := len(s.str) - s.wordLen
	for right := left; right <= n; right += s.wordLen {
		next := right + s.wordLen
		substr := s.str[right:next]

		maxCount, desired := s.wordMap[substr]

		if !desired {
			// not desired, reset
			left = next
			wordFound = make(map[string]int)
			foundNum = 0
			exceeded = false

			continue
		}

		// shrink window if found or exceeded in previous run
		for foundNum >= s.wordNum || exceeded {
			leftMost := s.str[left : left+s.wordLen]
			left += s.wordLen
			wordFound[leftMost]--

			if wordFound[leftMost] >= s.wordMap[leftMost] {
				exceeded = false
			} else {
				foundNum--
			}
		}

		wordFound[substr]++
		if wordFound[substr] <= maxCount {
			foundNum++
		} else {
			// exceeded
			exceeded = true
		}

		if foundNum == s.wordNum {
			// found answer
			answer = append(answer, left)
		}

	}

	return answer
}
