/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */

package main

func main() {
	res := validPalindrome("aba")
	println(res)
}

// @lc code=start
func validPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r && r < len(s) {
		if s[l] == s[r] {
			l++
			r--
			// s = s[l+1 : r+2] // 从两侧剥离字符会导致超时
			// println("s ", s)
		} else {
			return isValid(s[:l]+s[l+1:]) || isValid(s[:r]+s[r+1:])
		}
	}
	return true
}

func isValid(s string) bool {
	println("s is ", s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// @lc code=end
