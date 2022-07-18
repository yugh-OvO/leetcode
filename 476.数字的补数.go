/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start
func findComplement(num int) int {
	pos := 0
	res := 0
	for num > 0 {
		if num%2 == 0 {
			res += int(math.Pow(2, float64(pos)))
		}
		pos++
		num /= 2
	}
	return res
}

// @lc code=end

