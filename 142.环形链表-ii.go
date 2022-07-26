/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	hasFind := false
	for fast != nil && fast.Next != nil {
		if !hasFind {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
				fast = head
				hasFind = true
				if fast == slow {
					return fast
				}
			}
		} else {
			fast = fast.Next
			slow = slow.Next
			if fast == slow {
				return fast
			}
		}
	}
	return nil
}

// @lc code=end

