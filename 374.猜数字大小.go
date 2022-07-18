/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left := 0
	right := n
	for left <= right {
		mid := (left + right) / 2
		if guess(mid) == -1 {
			// 右标取中间值向左一位
			right = mid - 1
		} else if guess(mid) == 1 {
			// 左标取中间值向右一位
			left = mid + 1
		} else {
			return mid
		}
	}
	return n
}

// @lc code=end

