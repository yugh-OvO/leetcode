/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	// 从后向前遍历，如果最后一位是9，向上进1
	needAdd := true // 是否需要进位
	for i := len(digits) - 1; i >= 0; i-- {
		if needAdd {
			digits[i]++
			if digits[i] == 10 {
				digits[i] = 0
				needAdd = true
			} else {
				needAdd = false
			}
		}
	}
	// case [9]，需要向左侧加1
	if needAdd {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end

