/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

package main

func main() {
	res := longestPalindrome("abb")
	print(res)
}

// @lc code=start
func longestPalindrome(s string) string {
	ans := ""
	if len(s) == 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		// 右指针从右向左，判断是否是回文子串
		for r := len(s) - 1; r >= i; r-- {
			if isRight(s[i:r+1]) && len(s[i:r+1]) > len(ans) {
				ans = s[i : r+1]
			}
		}
	}
	return ans
}

func isRight(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end
