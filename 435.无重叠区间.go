/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	req := [][]int{}
	req = append(req, []int{1, 2})
	req = append(req, []int{2, 3})
	req = append(req, []int{3, 4})
	req = append(req, []int{1, 3})
	res := eraseOverlapIntervals(req)
	println(res)
}

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	ans := 0
	// 贪心，保留区间结尾小的元素
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	fmt.Printf("%v", intervals)
	prev := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prev {
			ans++
		} else {
			prev = intervals[i][1]
		}
	}
	return ans
}

// @lc code=end
