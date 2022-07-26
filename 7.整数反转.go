/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	ans := 0
	isMin := false
	if x < 0 {
		isMin = true
		x = 0 - x
	}
	for x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if ans > int(math.Pow(2, 31))-1 {
		return 0
	}
	if isMin {
		ans = 0 - ans
	}
	return ans
}

// @lc code=end

