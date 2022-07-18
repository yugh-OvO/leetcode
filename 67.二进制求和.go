/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	long := b
	short := a
	if len(a) > len(b) {
		long = a
		short = b
	}
	needAdd := 0 // 是否进位
	result := ""
	// 从右往左遍历
	for i := len(long) - 1; i >= 0; i-- {
		cur, _ := strconv.Atoi(string(long[i]))
		tmp := cur + needAdd
		if len(short) >= len(long)-i {
			shortCur, _ := strconv.Atoi(string(short[len(short)-len(long)+i]))
			tmp += shortCur
		}
		switch tmp {
		case 3:
			needAdd = 1
			result = "1" + result
		case 2:
			needAdd = 1
			result = "0" + result
		case 1:
			needAdd = 0
			result = "1" + result
		case 0:
			needAdd = 0
			result = "0" + result
		}
	}
	if needAdd == 1 {
		result = "1" + result
	}
	return result
}

// @lc code=end
