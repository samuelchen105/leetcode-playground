package problem

import (
	"math"
)

// 76. Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	mapT := make(map[rune]int)
	for _, r := range t {
		mapT[r]++
	}

	res := struct {
		start  int
		end    int
		length int
	}{
		start:  0,
		end:    -1,
		length: math.MaxInt,
	}

	mapS := make(map[rune]int)
	have := 0
	in := []rune(s)
	l := 0
	for r := 0; r < len(in); r++ {
		mapS[in[r]]++

		if ct, ok := mapT[in[r]]; ok && mapS[in[r]] == ct {
			have++
		}

		for l <= r && have == len(mapT) {
			windowSize := r - l + 1
			if windowSize < res.length {
				res.length = windowSize
				res.start = l
				res.end = r
			}

			if ct, ok := mapT[in[l]]; ok {
				mapS[in[l]]--
				if mapS[in[l]]+1 == ct {
					have--
				}
			}

			l++
		}
	}

	return string(in[res.start : res.end+1])
}
