/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		// 注意：s[i]取出的是byte类型，不能和string直接对比
		switch string(s[i]) {
		case "I":
			if i < len(s)-1 && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
				num -= 1
			} else {
				num += 1
			}
		case "V":
			num += 5
		case "X":
			if i < len(s)-1 && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
				num -= 10
			} else {
				num += 10
			}
		case "L":
			num += 50
		case "C":
			if i < len(s)-1 && (string(s[i+1]) == "D" || string(s[i+1]) == "M") {
				num -= 100
			} else {
				num += 100
			}
		case "D":
			num += 500
		case "M":
			num += 1000
		}
	}
	return num
}

// @lc code=end

