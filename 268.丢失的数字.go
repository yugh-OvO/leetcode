/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	sort.Ints(nums)
	if nums[0] != 0 {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] == 2 {
			return nums[i] + 1
		}
	}
	return len(nums)
	// 数学方法，计算总和，减当前和，算出缺哪个数字
	// n := len(nums)
	// total := n * (n + 1) / 2
	// arrSum := 0
	// for _, num := range nums {
	// 	arrSum += num
	// }
	// return total - arrSum
}

// @lc code=end

