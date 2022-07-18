/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 以head为key，存map，以判断是否成环
	seen := map[*ListNode]bool{}
	for head != nil {
		if seen[head] {
			return true
		}
		seen[head] = true
		head = head.Next
	}
	return false
}

// @lc code=end

