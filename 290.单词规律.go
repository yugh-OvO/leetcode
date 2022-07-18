/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	if len(arr) != len(pattern) {
		return false
	}
	relation := map[string]string{}
	// 这种找规律的问题要注意反向的规律也需要验证，比如
	// pattern: abba
	// s: dog dog dog dog
	relation2 := map[string]string{}
	for i := 0; i < len(arr); i++ {
		str, ok := relation[string(pattern[i])]
		if ok && str != arr[i] {
			return false
		}
		str2, ok := relation2[arr[i]]
		if ok && str2 != string(pattern[i]) {
			return false
		}
		relation[string(pattern[i])] = arr[i]
		relation2[arr[i]] = string(pattern[i])
	}
	return true
}

// @lc code=end

