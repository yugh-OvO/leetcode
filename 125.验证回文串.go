/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	// 考察点在数据预处理，要去除字符串数字以外的内容
	good := ""
	for i := 0; i < len(s); i++ {
		if isGood(s[i]) {
			good += string(s[i])
		}
	}
	// 把字符串转成小写
	good = strings.ToLower(good)
	for i := 0; i < len(good)/2; i++ {
		if good[i] != good[len(good)-i-1] {
			return false
		}
	}
	return true
}

func isGood(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

// @lc code=end

