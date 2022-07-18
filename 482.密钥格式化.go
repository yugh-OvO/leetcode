/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */

// @lc code=start
func licenseKeyFormatting(s string, k int) string {
	// 去破折号
	s = strings.Replace(s, "-", "", -1)
	s = strings.ToUpper(s)
	// 从右向左处理，第k个加破折号
	for i := len(s) - k; i >= 1; i -= k {
		s = string(s[:i]) + "-" + string(s[i:])
	}
	return s
}

// @lc code=end

