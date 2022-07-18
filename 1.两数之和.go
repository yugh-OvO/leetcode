/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	for k, v := range nums {
		for i, j := range nums {
			if k != i && v+j == target {
				return []int{k, i}
			}
		}
	}
	return []int{}
}

// @lc code=end

