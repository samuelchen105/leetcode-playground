package problem

// 217. Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	elemMap := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := elemMap[num]; ok {
			return true
		}
		elemMap[num] = struct{}{}
	}

	return false
}
