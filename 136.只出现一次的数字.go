/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		hasFind := false
		for j := 1; j < len(nums); j++ {
			if nums[j] == tmp {
				nums = append(nums[:j], nums[j+1:]...)
				hasFind = true
			}
		}
		if hasFind {
			nums = nums[i+1:]
			i--
		} else {
			return tmp
		}
	}
	return nums[0]
}

// @lc code=end

