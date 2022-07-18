/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// 动态规划
	// 若前一个元素大于0，则将其加到当前元素上
	// 当前数组中最大元素即最大子数组和
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
	// 贪心
	// 若当前指针所指元素之前的和小于0，则丢弃当前元素之前的数列
	// 有两个值：当前和，最大和
	// max := nums[0]
	// curMax := nums[0]
	// for i := 1; i < len(nums); i++ {
	// 	if nums[i] < curMax+nums[i] {
	// 		curMax = curMax + nums[i]
	// 	} else {
	// 		curMax = nums[i]
	// 	}
	// 	if curMax > max {
	// 		max = curMax
	// 	}
	// }
	// return max
}

// @lc code=end

