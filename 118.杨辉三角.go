/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := [][]int{}
	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		arr := []int{1}
		for j := 1; j < i; j++ {
			arr = append(arr, res[i-1][j-1]+res[i-1][j])
		}
		arr = append(arr, 1)
		res = append(res, arr)
	}
	return res
}

// @lc code=end

