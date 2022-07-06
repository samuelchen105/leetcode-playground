package problem

// 49. Group Anagrams
// https://leetcode.com/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	result := make(map[[26]int][]string)
	for _, str := range strs {
		key := [26]int{}
		for _, r := range str {
			key[r-'a']++
		}
		result[key] = append(result[key], str)
	}

	ret := make([][]string, 0, len(result))
	for _, v := range result {
		ret = append(ret, v)
	}

	return ret
}
