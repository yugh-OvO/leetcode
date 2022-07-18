/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	// 动态规划
	// 第n级台阶：从第n-1级台阶爬1级或从第n-2级台阶爬2级
	// 递推公式： fn = fn-1 + fn-2
	if n == 1 {
		return 1
	}
	first := 1  // 状态1
	second := 2 // 状态2
	// 滚动数组
	for i := 3; i <= n; i++ {
		third := first + second
		first = second // 先把状态2赋值到状态1的位置
		second = third // 再把状态3赋值到状态2的位置
	}
	return second
}

// @lc code=end

