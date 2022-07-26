/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(n int) []int {
	result := []int{}
	for i := 0; i <= n; i++ {
		result = append(result, cal(i))
	}
	return result
}

func cal(n int) (num int) {
	for n > 0 {
		if n%2 == 1 {
			num++
		}
		n /= 2
	}
	return
}

// @lc code=end

