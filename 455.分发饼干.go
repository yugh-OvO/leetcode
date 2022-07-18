/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	// 用最小的饼干去找最小胃口的小孩
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	for i := 0; i < len(s) && len(g) > 0; i++ {
		if s[i] >= g[0] {
			g = g[1:]
			res++
		}
	}
	return res
}

// @lc code=end

