/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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

func sumOfLeftLeaves(root *TreeNode) int {
	return cal(root, 0)
}

func cal(root *TreeNode, num int) int {
	if root == nil {
		return num
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		num += root.Left.Val
	}
	num = cal(root.Left, num)
	num = cal(root.Right, num)
	return num
}

// @lc code=end

