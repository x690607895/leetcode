package main

import (
	"TreeNode"
	"fmt"
)

// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。

// 叶子节点 是指没有子节点的节点。

// 思路
// 递归遍历到叶子结点求和，如果等于targetSum返回true

func hasPathSum(root *TreeNode.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

func main() {
	fmt.Println(hasPathSum(TreeNode.CreateTree([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1}), 22))
	fmt.Println(hasPathSum(TreeNode.CreateTree([]int{1, 2, 3}), 22))
	fmt.Println(hasPathSum(TreeNode.CreateTree([]int{1, 2, 3}), 2))
	fmt.Println(hasPathSum(TreeNode.CreateTree([]int{-2, 0, -3}), -5))
}
