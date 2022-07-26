/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	if len(nums) == 0 {
		return ans
	}
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			for i := mid; i >= 0; i-- {
				if nums[i] != target {
					break
				}
				ans[0] = i
			}
			for i := mid; i < len(nums); i++ {
				if nums[i] != target {
					break
				}
				ans[1] = i
			}
			break
		}
	}
	return ans
}

// @lc code=end

