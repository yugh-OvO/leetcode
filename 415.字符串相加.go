/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

package main

import "strconv"

func main() {
	res := addStrings("11", "123")
	println("123", res)
}

// @lc code=start
func addStrings(num1 string, num2 string) string {
	short := num1
	long := num1
	if len(num1) > len(num2) {
		short = num2
	} else {
		long = num2
	}
	res := ""
	needAdd := 0 // 是否需要进位
	for i := 0; i < len(long); i++ {
		// 遍历较长字符串
		curNum := ""
		if len(short)-i-1 >= 0 {
			print("long", string(long[len(long)-i-1]))
			needAdd, curNum = calNum(string(long[len(long)-i-1]), string(short[len(short)-i-1]), needAdd)
		} else {
			needAdd, curNum = calNum(string(long[len(long)-i-1]), "0", needAdd)
		}
		println("333", curNum)
		res = curNum + res
	}
	if needAdd == 1 {
		res = "1" + res
	}
	return res
}

func calNum(num1 string, num2 string, needAdd int) (newNeedAdd int, res string) {
	println("num1", num1)
	println("2 is", num2)
	int1, _ := strconv.Atoi(num1)
	int2, _ := strconv.Atoi(num2)
	resInt := int1 + int2 + needAdd
	if resInt > 9 {
		newNeedAdd = 1
	}
	// println("3321", string(1))
	res = strconv.Itoa(resInt % 10)
	return
}

// @lc code=end
