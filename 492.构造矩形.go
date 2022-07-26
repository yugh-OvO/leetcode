/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */

// @lc code=start
func constructRectangle(area int) []int {
	res := []int{}
	w := 1
	l := 1
	min := area
	for w <= area/w {
		if area%w != 0 {
			w++
			continue
		}
		l = area / w
		if l-w < min {
			min = l - w
			res = []int{l, w}
		}
		w++
	}
	return res
}

// @lc code=end

