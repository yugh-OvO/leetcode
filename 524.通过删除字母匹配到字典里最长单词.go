/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lc code=start
func findLongestWord(s string, dictionary []string) string {
	tmp := ""
	for i := 0; i < len(dictionary); i++ {
		if judge(dictionary[i], s) {
			if len(tmp) < len(dictionary[i]) || (len(tmp) == len(dictionary[i]) && judgeFont(dictionary[i], tmp)) {
				tmp = dictionary[i]
			}
		}
	}
	return tmp
}

func judge(a, b string) bool {
	for i := 0; i < len(b); i++ {
		if b[i] == a[0] {
			if len(a) == 1 {
				return true
			}
			a = a[1:]
		}
	}
	return false
}

func judgeFont(a, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return false
		} else if a[i] < b[i] {
			return true
		}
	}
	return false
}

// @lc code=end

