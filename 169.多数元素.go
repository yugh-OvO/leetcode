/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	half := len(nums) / 2
	for {
		tmp := nums[0]
		num := 1
		for j := 1; j < len(nums); j++ {
			if nums[j] == tmp {
				num++
				if num > half {
					return tmp
				}
				nums = append(nums[:j], nums[j+1:]...)
				j--
			}
		}
		nums = nums[1:]
		if len(nums) == 0 {
			return tmp
		}
	}
	return 0
}

// @lc code=end

