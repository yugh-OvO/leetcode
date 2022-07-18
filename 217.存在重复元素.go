/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	// hash判断效率更高，复杂度 O(N)
	hash := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if hash[nums[i]] {
			return true
		}
		hash[nums[i]] = true
	}
	return false
}

// @lc code=end

