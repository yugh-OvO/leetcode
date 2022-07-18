/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3 的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	for n > 0 {
		if n == 1 {
			return true
		}
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return false
}

// @lc code=end

