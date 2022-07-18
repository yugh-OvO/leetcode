/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	// 双指针
	data := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
	// left := 0
	right := len(s) - 1
	for left := 0; left < len(s)/2; left++ {
		if data[string(s[left])] {
			// 遍历右指针
			for !data[string(s[right])] {
				if right <= len(s)/2 {
					break
				}
				right--
			}
		}
	}
	return s
}

// @lc code=end

