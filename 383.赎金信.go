/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	relation := map[byte]int{}
	for i := 0; i < len(magazine); i++ {
		relation[magazine[i]] = relation[magazine[i]] + 1
	}
	for j := 0; j < len(ransomNote); j++ {
		if relation[ransomNote[j]] == 0 {
			return false
		}
		relation[ransomNote[j]] = relation[ransomNote[j]] - 1
	}
	return true
}

// @lc code=end

