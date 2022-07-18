/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

package main

import "math"

func main() {
	res := isPalindrome(11)
	println(res)
}

// @lc code=start
func isPalindrome(x int) bool {
	// x 121
	// 从两端向中间收束
	if x < 0 {
		// 排除负数
		return false
	}
	length := 0
	a := x
	if a > 0 {
		for a > 0 {
			a /= 10
			length++
		}
		// println("length", length)
		for i := 0; i <= length/2; i++ {
			// println("循环", dumpNum(x, length-i), dumpNum(x, i))
			if dumpNum(x, length-i-1) != dumpNum(x, i) {
				return false
			}
		}
	}
	return true
}

// 指定位数返回对应位置数字
func dumpNum(x, pos int) int {
	// println(11 % 1)
	println(pos)
	println(x / int(math.Pow10(pos)) % 10)
	return (x / int(math.Pow10(pos)) % 10)
}

// @lc code=end
