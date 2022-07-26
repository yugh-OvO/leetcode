/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	num := 0
	// 3,4,2,3
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if i == 0 || nums[i+1] >= nums[i-1] {
				nums[i] = nums[i+1]
			} else {
				nums[i+1] = nums[i]
			}
			num++
		}
		if num == 2 {
			return false
		}
	}
	return true
}

// @lc code=end

