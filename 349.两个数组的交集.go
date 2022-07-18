/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	relation := map[int]bool{}
	joined := map[int]bool{}
	for i := 0; i < len(nums1); i++ {
		relation[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		if relation[nums2[i]] && !joined[nums2[i]] {
			res = append(res, nums2[i])
			joined[nums2[i]] = true
		}
	}
	return res
}

// @lc code=end

