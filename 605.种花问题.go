/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	ans := 0
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			ans = 1
		}
		return ans >= n
	}
	if flowerbed[1] == 0 && flowerbed[0] == 0 {
		flowerbed[0] = 1
		ans++
	}
	if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
		flowerbed[len(flowerbed)-1] = 1
		ans++
	}
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			ans++
		}
	}
	return ans >= n
}

// @lc code=end

