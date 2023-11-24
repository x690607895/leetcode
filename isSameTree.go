package main

import (
	"TreeNode"
	"fmt"
)

func main() {
	fmt.Println(isSameTree(TreeNode.CreateTree([]int{1, 2, 3}), TreeNode.CreateTree([]int{1, 2, 3})))
	fmt.Println(isSameTree(TreeNode.CreateTree([]int{1, 2, 1}), TreeNode.CreateTree([]int{1, 1, 2})))
}

func isSameTree(p *TreeNode.TreeNode, q *TreeNode.TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
