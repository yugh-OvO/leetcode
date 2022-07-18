/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	tmp := ""
	for i := 0; i < len(s)/2+1; i++ {
		tmp += string(s[i])
		judge := ""
		if len(s)%len(tmp) != 0 {
			continue
		}
		if len(s) == len(tmp) {
			break
		}
		for j := 0; j < len(s)/len(tmp); j++ {
			judge += tmp
		}
		if judge == s {
			return true
		}
	}
	return false
}

// @lc code=end

