/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	relation := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		relation[nums[i]] = true
	}
	for i := 0; i < len(nums); i++ {
		if !relation[i+1] {
			res = append(res, i+1)
		}
	}
	return res
}

// @lc code=end

