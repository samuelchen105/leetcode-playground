package problem

// 11. Container With Most Water
// https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
	ret := 0
	l := 0
	r := len(height) - 1
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		if area > ret {
			ret = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
