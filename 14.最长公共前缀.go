/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := ""
	// 取第一个字符串的第一个字符，对比所有字符，看是否是公共前缀，如果不是，退出循环
	for i := 0; i < len(strs[0]); i++ {
		// 对比的字符 string(strs[0][i])
		pattern := string(strs[0][i])
		noMatch := false
		for _, str := range strs {
			if len(str) <= i || string(str[i]) != pattern {
				noMatch = true
				break
			}
		}
		if noMatch {
			break
		} else {
			prefix += pattern
		}
	}
	return prefix
}

// @lc code=end

