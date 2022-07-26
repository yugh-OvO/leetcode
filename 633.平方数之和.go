/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start
func judgeSquareSum(c int) bool {
	l := 0
	r := int(math.Sqrt(float64(c)))
	for l <= r {
		sum := l*l + r*r
		if sum > c {
			// 右指针左移
			r--
		} else if sum < c {
			l++
		} else {
			return true
		}
	}
	return false
}

// @lc code=end

