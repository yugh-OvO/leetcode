/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		if numbers[l]+numbers[r] > target {
			r--
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			break
		}
	}
	return []int{l + 1, r + 1}
}

// @lc code=end

