/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	big := x
	small := y
	if y > x {
		big = y
		small = x
	}
	ans := 0
	for big > 0 {
		if big%2 != small%2 {
			ans++
		}
		big /= 2
		small /= 2
	}
	return ans
}

// @lc code=end

