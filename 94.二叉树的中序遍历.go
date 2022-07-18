/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := inorder(root, make([]int, 0))
	return res
}

func inorder(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = inorder(node.Left, res)
	res = append(res, node.Val)
	res = inorder(node.Right, res)
	return res
}

// @lc code=end

