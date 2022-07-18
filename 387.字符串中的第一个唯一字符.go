/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	relation := map[byte]int{}
	for i := 0; i < len(s); i++ {
		relation[s[i]] = relation[s[i]] + 1
	}
	for j := 0; j < len(s); j++ {
		if relation[s[j]] == 1 {
			return j
		}
	}
	return -1
}

// @lc code=end

