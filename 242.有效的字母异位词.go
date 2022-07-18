/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	pos1 := map[string]int{}
	pos2 := map[string]int{}
	for i := 0; i < len(s); i++ {
		_, ok := pos1[string(s[i])]
		if ok {
			pos1[string(s[i])]++
		} else {
			pos1[string(s[i])] = 1
		}
		_, ok = pos2[string(t[i])]
		if ok {
			pos2[string(t[i])]++
		} else {
			pos2[string(t[i])] = 1
		}
	}
	for k, v := range pos1 {
		if v != pos2[k] {
			return false
		}
	}
	return true
}

// @lc code=end

