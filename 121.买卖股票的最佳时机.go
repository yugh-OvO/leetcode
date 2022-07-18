/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 动态规划
	// 当前盈利
	// 最大盈利
	cur := 0
	max := 0
	// 如果当前盈利为负数，放弃掉，已prices[i]作为买入点
	for i := 1; i < len(prices); i++ {
		cur += prices[i] - prices[i-1]
		if cur < 0 {
			cur = 0
		}
		if max < cur {
			max = cur
		}
	}
	return max
}

// @lc code=end

