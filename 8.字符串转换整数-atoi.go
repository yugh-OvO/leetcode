/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

package main

import "math"

func main() {
	res := myAtoi("3.14159")
	println(res)
}

// @lc code=start
func myAtoi(s string) int {
	ans := 0
	isMin := false
	relation := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	for i := 0; i < len(s); i++ {
		if i == 0 && string(s[i]) == " " {
			s = s[1:]
			i--
			continue
		}
		if string(s[i]) == " " {
			break
		}
		if string(s[i]) == "." {
			break
		}
		if ans == 0 && string(s[i]) == "-" {
			isMin = true
		}
		num, ok := relation[string(s[i])]
		if ok {
			ans = ans*10 + num
		}
	}
	if isMin {
		ans = 0 - ans
	}
	if ans > int(math.Pow(2, 31))-1 {
		return int(math.Pow(2, 31)) - 1
	}
	if ans < 0-int(math.Pow(2, 31)) {
		return 0 - int(math.Pow(2, 31))
	}
	return ans
}

// @lc code=end
