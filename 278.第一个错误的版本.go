/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right { // 循环直至区间左右端点相同
		mid := left + (right-left)/2 // 防止计算时溢出
		if isBadVersion(mid) {
			right = mid // 答案在区间 [left, mid] 中
		} else {
			left = mid + 1 // 答案在区间 [mid+1, right] 中
		}
	}
	// 此时有 left == right，区间缩为一个点，即为答案
	return left
}

// @lc code=end

