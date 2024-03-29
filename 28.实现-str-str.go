/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

package main

func main() {
	res := strStr("hello", "ll")
	println(res)
}

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		println(string(haystack[i : i+len(needle)-1]))
		if string(haystack[i:i+len(needle)]) == needle {
			return i
		}
	}
	return -1
}

// @lc code=end
