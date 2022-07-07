package problem

import "sort"

// 15. 3Sum
// https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	ret := [][]int{}

	sort.Ints(nums)

	for i, n := range nums {
		if i > 0 && nums[i-1] == n {
			continue
		}

		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := n + nums[l] + nums[r]
			if sum == 0 {
				ret = append(ret, []int{n, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return ret
}
