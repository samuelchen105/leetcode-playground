package problem

// 1. Two Sum
// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)
	for i, n := range nums {
		if prevIdx, ok := prevMap[target-n]; ok {
			return []int{prevIdx, i}
		}
		prevMap[n] = i
	}

	return []int{}
}
