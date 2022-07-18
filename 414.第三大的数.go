/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	sort.Ints(nums)
	times := 0
	max := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < max {
			max = nums[i]
			times++
		}
		if times == 2 {
			return max
		}
	}
	return nums[len(nums)-1]
}

// @lc code=end

