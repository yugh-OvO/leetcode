/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	ans := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			ans += prices[i+1] - prices[i]
		}
	}
	return ans
}

// @lc code=end

