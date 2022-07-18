/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */

// @lc code=start
func findWords(words []string) []string {
	l1 := map[string]bool{
		"Q": true,
		"W": true,
		"E": true,
		"R": true,
		"T": true,
		"Y": true,
		"U": true,
		"I": true,
		"O": true,
		"P": true,
	}
	l2 := map[string]bool{
		"A": true,
		"S": true,
		"D": true,
		"F": true,
		"G": true,
		"H": true,
		"J": true,
		"K": true,
		"L": true,
	}
	l3 := map[string]bool{
		"Z": true,
		"X": true,
		"C": true,
		"V": true,
		"B": true,
		"N": true,
		"M": true,
	}
	result := []string{}
	for _, word := range words {
		isTrue := true
		for i := 1; i < len(word); i++ {
			if l1[strings.ToUpper(string(word[0]))] && !l1[strings.ToUpper(string(word[i]))] {
				isTrue = false
				break
			}
			if l2[strings.ToUpper(string(word[0]))] && !l2[strings.ToUpper(string(word[i]))] {
				isTrue = false
				break
			}
			if l3[strings.ToUpper(string(word[0]))] && !l3[strings.ToUpper(string(word[i]))] {
				isTrue = false
				break
			}
		}
		if isTrue {
			result = append(result, word)
		}
	}
	return result
}

// @lc code=end

