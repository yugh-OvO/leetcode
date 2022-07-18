/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

package main

func main() {
	convertToTitle(52)
}

// @lc code=start
func convertToTitle(columnNumber int) string {
	str, _ := convert("", columnNumber)
	return str
}

func convert(str string, number int) (string, int) {
	cols := "ZABCDEFGHIJKLMNOPQRSTUVWXY"
	print(number)
	println(string(cols[number%26]))
	str = string(cols[number%26]) + str
	if number > 26 {
		str, number = convert(str, (number-1)/26)
	}
	return str, number
}

// @lc code=end
