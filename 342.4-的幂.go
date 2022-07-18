/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	for n > 0 {
		if n == 1 {
			return true
		}
		if n%4 != 0 {
			return false
		}
		n /= 4
	}
	return false
}

// @lc code=end

