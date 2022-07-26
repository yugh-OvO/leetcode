/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := []int{}
	for i := 0; i < len(nums1); i++ {
		find := false
		appended := false
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == nums1[i] {
				find = true
			}
			if find && nums2[j] > nums1[i] {
				ans = append(ans, nums2[j])
				appended = true
				break
			}
		}
		if !appended {
			ans = append(ans, -1)
		}
	}
	return ans
}

// @lc code=end

