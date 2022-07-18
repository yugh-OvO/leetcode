/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	relation := map[byte]int{}
	for i := 0; i < len(s); i++ {
		relation[s[i]]++
	}
	for j := 0; j < len(t); j++ {
		if relation[t[j]] == 0 {
			return t[j]
		}
		relation[t[j]]--
	}
	return t[0]
}

// @lc code=end

