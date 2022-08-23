package problem

// 153. Find Minimum in Rotated Sorted Array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	res := nums[0]

	for l <= r {
		if nums[l] < nums[r] {
			if nums[l] < res {
				res = nums[l]
			}
		}

		mid := (l + r) / 2
		if nums[mid] < res {
			res = nums[mid]
		}

		if nums[mid] >= nums[l] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}
