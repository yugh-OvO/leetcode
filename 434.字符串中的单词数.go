/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */

// @lc code=start
func countSegments(s string) int {
	// 找左侧有空格的字母
	res := 0
	s = " " + s
	for i := 1; i < len(s); i++ {
		if string(s[i]) != " " && string(s[i-1]) == " " {
			res++
		}
	}
	return res
}

// @lc code=end

