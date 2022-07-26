/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	// 其实是求不重叠的数组数量
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	fmt.Printf("%v", points)
	prev := points[0][1]
	num := 0
	for i := 1; i < len(points); i++ {
		if points[i][0] <= prev {
			num++
		} else {
			prev = points[i][1]
		}
	}
	return len(points) - num
}

// @lc code=end

