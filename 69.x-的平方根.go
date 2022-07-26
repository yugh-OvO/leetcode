/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	// 方法一：袖珍计算器算法，数学公式反推
	// if x == 0 {
	// 	return 0
	// }
	// ans := int(math.Exp(0.5 * math.Log(float64(x))))
	// if (ans+1)*(ans+1) <= x {
	// 	return ans + 1
	// }
	// return ans
	// 方法二：二分查找
	// 由于x平方根的整数部分ans是满足k²<=x的最大k值，因此我们可以对k进行二分查找，从而得到答案
	// 二分查找的下界为 0，上界可以粗略地设定为 x。在二分查找的每一步中，我们只需要比较中间元素 mid 的平方与 x 的大小关系，并通过比较的结果调整上下界的范围。由于我们所有的运算都是整数运算，不会存在误差，因此在得到最终的答案 ans 后，也就不需要再去尝试 ans+1 了。
	l := 0
	r := x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid + 1
			ans = mid
		}
	}
	return ans
}

// @lc code=end

