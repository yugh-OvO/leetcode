/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		newHead = &ListNode{
			Next: newHead,
			Val:  head.Val,
		}
		head = head.Next
	}
	return newHead
}

// @lc code=end

