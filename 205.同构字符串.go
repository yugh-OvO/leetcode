/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	relation := map[string]string{}
	// 忽略了t字符串是baba，s是badc的情况，要用t再验证一遍s是否同构
	relation2 := map[string]string{}
	for i := 0; i < len(s); i++ {
		res, ok := relation[string(s[i])]
		if ok && res != string(t[i]) {
			return false
		}
		res2, ok := relation2[string(t[i])]
		if ok && res2 != string(s[i]) {
			return false
		}
		relation[string(s[i])] = string(t[i])
		relation2[string(t[i])] = string(s[i])
	}
	return true
}

// @lc code=end

