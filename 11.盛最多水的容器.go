/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	// 双指针，较小数字的指针向中间移动
	// 感觉这个移动有点博弈论的味了，每次都移动自己最差的一边，虽然可能变得更差，但是总比不动（或者减小）强，动最差的部分可能找到更好的结果，但是动另一边总会更差或者不变，兄弟们，这不是题，这是人生，逃离舒适圈！！（这解释我觉得无敌了，哈哈哈）
	ans := 0
	l := 0
	r := len(height) - 1
	for l < r {
		if min(height[r], height[l])*(r-l) > ans {
			ans = min(height[r], height[l]) * (r - l)
		}
		if height[r] > height[l] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

