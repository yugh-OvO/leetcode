/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	res := [][]int{}
	arr := []int{1}
	res = append(res, arr)
	for i := 1; i < rowIndex+1; i++ {
		arr = []int{1}
		for j := 1; j < i; j++ {
			arr = append(arr, res[i-1][j-1]+res[i-1][j])
		}
		arr = append(arr, 1)
		res = append(res, arr)
	}
	return arr
}

// @lc code=end

