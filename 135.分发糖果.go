/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

package main

func main() {
	res := candy([]int{1, 2, 87, 87, 87, 2, 1})
	println(res)
}

// @lc code=start
func candy(ratings []int) int {
	// 先给每个孩子都设置为1，
	ans := 0
	candy := []int{}
	for i := 0; i < len(ratings); i++ {
		candy = append(candy, 1)
	}
	for i := 1; i < len(candy); i++ {
		// 从左向右遍历，如果右侧比左侧大，右侧比左侧+1
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		}
	}
	// 还要倒着来一遍
	for i := len(candy) - 1; i >= 1; i-- {
		// 从左向右遍历，如果左侧比右侧大，左侧取左侧和右侧+1的最大值
		if ratings[i-1] > ratings[i] {
			candy[i-1] = max(candy[i-1], candy[i]+1)
		}
	}
	for i := 0; i < len(candy); i++ {
		// println(candy[i])
		ans += candy[i]
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
