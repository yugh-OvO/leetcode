/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(s string) []int {
	pos := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if pos[s[i]] < i {
			pos[s[i]] = i
		}
	}
	ans := []int{}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if pos[s[i]] > end {
			end = pos[s[i]]
		}
		if i == end {
			ans = append(ans, end-start+1)
			start = i + 1
		}
	}
	return ans
}

// @lc code=end

