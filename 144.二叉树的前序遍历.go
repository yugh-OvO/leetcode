/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	res = append(res, root.Val)
	if root.Left != nil {
		res = append(res, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, preorderTraversal(root.Right)...)
	}
	return res
}

// @lc code=end

