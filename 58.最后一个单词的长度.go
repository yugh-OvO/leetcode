/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	// 反方向遍历字符串，找到第一个不是空的字符位置，和后续的第一个是空的字符位置
	min := 0
	max := -1
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			min = i
			break
		}
	}
	// case : "a"
	// 如果找不到max，则max为-1
	for i := min - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			max = i
			break
		}
	}
	return min - max
}

// @lc code=end

