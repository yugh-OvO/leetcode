/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
func longestPalindrome(s string) int {
	relation := map[byte]bool{}
	res := 0
	for i := 0; i < len(s); i++ {
		if relation[s[i]] {
			res += 2
			relation[s[i]] = false
		} else {
			relation[s[i]] = true
		}
	}
	if res < len(s) {
		res += 1
	}
	return res
}

// @lc code=end

