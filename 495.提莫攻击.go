/*
 * @lc app=leetcode.cn id=495 lang=golang
 *
 * [495] 提莫攻击
 */

// @lc code=start
func findPoisonedDuration(timeSeries []int, duration int) int {
	// 把每次中毒的时间叠加，减去重叠的时间
	times := len(timeSeries) * duration
	for i := 0; i < len(timeSeries)-1; i++ {
		// 求重叠时间，如果一次中毒后，在 t + duration - 1 时间内又中箭，发生时间叠加
		if timeSeries[i]+duration-1 >= timeSeries[i+1] {
			times -= timeSeries[i] + duration - timeSeries[i+1]
		}
	}
	return times
}

// @lc code=end

