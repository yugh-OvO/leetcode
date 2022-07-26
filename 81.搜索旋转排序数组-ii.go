/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			nums = append(nums[i+1:], nums[:i+1]...)
		}
	}
	fmt.Printf("%v", nums)
	// 二分法查找
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}

// @lc code=end

