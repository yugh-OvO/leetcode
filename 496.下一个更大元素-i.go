/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	relation := map[int]int{}
	tmp := -1
	for k := len(nums2) - 1; k >= 0; k-- {
		// 从右往左找，如果数字比tmp大，则重置tmp，设置该数字的value为-1
		if nums2[k] > tmp {
			tmp = nums2[k]
			relation[nums2[k]] = -1
		} else {
			relation[nums2[k]] = tmp
		}
	}
	result := []int{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {

			}
		}
	}
	return result
}

// @lc code=end

