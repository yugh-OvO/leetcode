/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 */

// @lc code=start
func readBinaryWatch(turnedOn int) (ans []string) {
	for i := 0; i < 1024; i++ {
		h, m := i>>6, i&63 // 用位运算取出高 4 位和低 6 位
		if h < 12 && m < 60 && bits.OnesCount(uint(i)) == turnedOn {
			ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
		}
	}
	return
}

// @lc code=end

