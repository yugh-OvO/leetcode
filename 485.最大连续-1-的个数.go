/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续 1 的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	tmp := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp++
		} else {
			tmp = 0
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

// @lc code=end

