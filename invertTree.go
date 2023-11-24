package main

import (
	"TreeNode"
)

func main() {
	TreeNode.PrintListNode(invertTree(TreeNode.CreateTree([]int{4, 2, 7, 1, 3, 6, 9})))
}

func invertTree(root *TreeNode.TreeNode) *TreeNode.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
