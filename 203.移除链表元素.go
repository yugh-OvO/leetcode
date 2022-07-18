/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// 也可以用迭代的方法删除链表中所有节点值等于特定值的节点。

	// 用 \textit{temp}temp 表示当前节点。如果 \textit{temp}temp 的下一个节点不为空且下一个节点的节点值等于给定的 \textit{val}val，则需要删除下一个节点。删除下一个节点可以通过以下做法实现：

	// \textit{temp}.\textit{next} = \textit{temp}.\textit{next}.\textit{next}
	// temp.next=temp.next.next

	// 如果 \textit{temp}temp 的下一个节点的节点值不等于给定的 \textit{val}val，则保留下一个节点，将 \textit{temp}temp 移动到下一个节点即可。

	// 当 \textit{temp}temp 的下一个节点为空时，链表遍历结束，此时所有节点值等于 \textit{val}val 的节点都被删除。

	// 具体实现方面，由于链表的头节点 \textit{head}head 有可能需要被删除，因此创建哑节点 \textit{dummyHead}dummyHead，令 \textit{dummyHead}.\textit{next} = \textit{head}dummyHead.next=head，初始化 \textit{temp}=\textit{dummyHead}temp=dummyHead，然后遍历链表进行删除操作。最终返回 \textit{dummyHead}.\textit{next}dummyHead.next 即为删除操作后的头节点。
	dummyHead := &ListNode{Next: head}
	for tmp := dummyHead; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummyHead.Next
}

// @lc code=end

