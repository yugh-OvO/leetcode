/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

package main

import "math"

func main() {
	res := minWindow("a", "a")
	println(res)
}

// @lc code=start
func minWindow(s string, t string) string {
	// ans := s
	// relation := map[byte]int{}
	// hasFind := false
	// fmt.Printf("%v", relation)
	// num := len(t)
	// for i := 0; i < len(s); i++ {
	// 	num = len(t)
	// 	for i := 0; i < len(t); i++ {
	// 		relation[t[i]]++
	// 	}
	// 	println(num)
	// 	for j := i; j < len(s); j++ {
	// 		if relation[s[j]] > 0 {
	// 			relation[s[j]]--
	// 			num--
	// 			fmt.Printf("%v", relation)
	// 			if num == 0 {
	// 				if j-i+1 <= len(ans) {
	// 					hasFind = true
	// 					ans = s[i : j+1]
	// 				}
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	// println(ans)
	// if ans == s && !hasFind {
	// 	return ""
	// }
	// return ans
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

// @lc code=end
