package problem

// 347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	bucket := make([][]int, len(nums)+1)
	for n, c := range count {
		bucket[c] = append(bucket[c], n)
	}

	ret := make([]int, 0, k)
	for i := len(bucket) - 1; i > 0; i-- {
		for _, n := range bucket[i] {
			if len(ret) == k {
				return ret
			}
			ret = append(ret, n)
		}
	}

	return ret
}
