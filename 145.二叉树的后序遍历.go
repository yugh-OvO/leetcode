/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	if root.Right != nil {
		res = append(postorderTraversal(root.Right), res...)
	}
	if root.Left != nil {
		res = append(postorderTraversal(root.Left), res...)
	}
	return res
}

// @lc code=end

