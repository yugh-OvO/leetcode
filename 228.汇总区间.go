/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	res := []string{}
	start := nums[0]
	// end := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			startStr := strconv.Itoa(start)
			endStr := strconv.Itoa(nums[i-1])
			if startStr == endStr {
				res = append(res, startStr)
			} else {
				res = append(res, startStr+"->"+endStr)
			}
			start = nums[i]
		}
	}
	startStr := strconv.Itoa(start)
	endStr := strconv.Itoa(nums[len(nums)-1])
	if startStr == endStr {
		res = append(res, startStr)
	} else {
		res = append(res, startStr+"->"+endStr)
	}
	return res
}

// @lc code=end

