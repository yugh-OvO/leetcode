/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := calNum(l1)
	num2 := calNum(l2)
	return &ListNode{
		Val: num2,
	}
	resultNum := num1 + num2
	println(resultNum)
	ans := &ListNode{}
	ans = appendResult(ans, resultNum)

	return ans
}

func appendResult(ans *ListNode, num int) *ListNode {
	if num == 0 {
		return ans
	}
	// ans.Val = num % 10
	ans.Next = &ListNode{
		Val: num % 10,
	}
	return appendResult(ans, num/10)
}

func calNum(l *ListNode) int {
	num := 0
	for l != nil {
		num += num*10 + l.Val
		l = l.Next
	}
	return num
}

// @lc code=end

