/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	// moveTimes := 0
	// for i := 0; i < len(nums)-moveTimes; i++ {
	// 	if nums[i] == 0 {
	// 		nums = append(nums[:i], nums[i+1:]...)
	// 		nums = append(nums, 0)
	// 		moveTimes++
	// 		i--
	// 	}
	// }
	// 双指针
	//使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。

	// 右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。

	// 注意到以下性质：

	// 左指针左边均为非零数；

	// 右指针左边直到左指针处均为零。

	// 因此每次交换，都是将左指针的零与右指针的非零数交换，且非零数的相对顺序并未改变。

	left := 0
	right := 0
	n := len(nums)
	for right < n {
		if nums[right] != 0 {
			// 交换
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// @lc code=end

