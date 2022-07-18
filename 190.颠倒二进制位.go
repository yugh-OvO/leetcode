/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	i := 0
	for num > 0 {
		if num%2 == 1 {
			res += uint32(math.Pow(2, float64(32-i-1)))
		}
		num /= 2
		i++
	}
	return res
}

// @lc code=end

