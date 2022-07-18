/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	i := 1
	for {
		n -= i
		if n < 0 {
			i = i - 1
			break
		} else if n == 0 {
			break
		}
		i++
	}
	return i
}

// @lc code=end

