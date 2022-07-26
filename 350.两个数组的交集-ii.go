/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	relation := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		relation[nums1[i]]++
	}
	for j := 0; j < len(nums2); j++ {
		if relation[nums2[j]] > 0 {
			res = append(res, nums2[j])
			relation[nums2[j]]--
		}
	}
	return res
}

// @lc code=end

