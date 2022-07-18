/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	for i := 0; i < len(list)/2; i++ {
		if list[i] != list[len(list)-i-1] {
			return false
		}
	}
	return true
}

// @lc code=end

