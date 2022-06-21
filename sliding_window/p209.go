package problem

// Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/

func minSubArrayLen(target int, nums []int) int {
	ans := 0
	sum := 0
	left, right := 0, 0
	for right < len(nums) {
		sum += nums[right]

		for left <= right && sum >= target {
			newLen := right - left + 1
			if ans == 0 || newLen < ans {
				ans = newLen
			}

			sum -= nums[left]
			left++
		}

		right++
	}

	return ans
}
