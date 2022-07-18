/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(n int) bool {
	n = cal(n)
	return n == 1
}

func cal(n int) int {
	if n == 0 {
		return 0
	}
	if n%5 == 0 {
		n = cal(n / 5)
	}
	if n%3 == 0 {
		n = cal(n / 3)
	}
	if n%2 == 0 {
		n = cal(n / 2)
	}
	return n
}

// @lc code=end

