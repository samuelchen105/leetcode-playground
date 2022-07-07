package problem

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	ret := make([]int, len(nums))

	ret[0] = 1
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		ret[i] = pre
		pre *= nums[i]
	}

	pre = 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= pre
		pre *= nums[i]
	}

	return ret
}
