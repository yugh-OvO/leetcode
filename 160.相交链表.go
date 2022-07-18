/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	seen := map[*ListNode]bool{}
	for headA != nil {
		seen[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if seen[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// @lc code=end

