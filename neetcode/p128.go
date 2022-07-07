package problem

// 128. Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/

func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, n := range nums {
		numSet[n] = struct{}{}
	}

	maxLen := 0
	for n := range numSet {
		if _, ok := numSet[n-1]; ok {
			continue
		}

		length := 1
		for {
			if _, ok := numSet[n+length]; !ok {
				break
			}
			length++
		}

		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}
