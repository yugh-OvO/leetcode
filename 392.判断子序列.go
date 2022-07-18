/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	for i := 0; i < len(t) && len(s) > 0; i++ {
		if s[0] == t[i] {
			s = s[1:]
		}
	}
	if len(s) == 0 {
		return true
	}
	return false
}

// @lc code=end

