/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小操作次数使数组元素相等
 */

// @lc code=start
func minMoves(nums []int) int {
	// 反推，每次把最大的数字减1，直至变成最小的数字
	// 所以总和为每个数字减最小数字的累加总和
	res := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		res += nums[i] - nums[0]
	}
	return res
}

// @lc code=end

