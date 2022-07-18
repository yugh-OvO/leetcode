/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	// 一层一层剥离相邻的括号，直至最终如果不是空字符串，返回false
	s = stripStr(s)
	return s == ""
}

func stripStr(s string) string {
	matchMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	newStr := s
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && matchMap[string(s[i])] == string(s[i+1]) {
			// 去除中间的括号
			newStr = string(s[:i]) + string(s[i+2:])
		}
	}
	if newStr != s {
		newStr = stripStr(newStr)
	}
	return newStr
}

// @lc code=end

