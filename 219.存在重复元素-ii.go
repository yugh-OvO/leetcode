/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	pos := map[int]int{}
	for i, num := range nums {
		p, ok := pos[num]
		if ok && i-p <= k {
			return true
		}
		pos[num] = i
	}
	return false
}

// @lc code=end

