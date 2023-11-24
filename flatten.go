package main

import (
	"TreeNode"
)

// 给你二叉树的根结点 root ，请你将它展开为一个单链表：

// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。

// 思路
// dfs遍历整个二叉树，用前序遍历，加入数组，然后修改left和right节点

func flatten(root *TreeNode.TreeNode) {
	if root == nil {
		return
	}
	arrTreeNode := []*TreeNode.TreeNode{}
	var dfs func(root2 *TreeNode.TreeNode)
	dfs = func(root2 *TreeNode.TreeNode) {
		if root2 == nil {
			return
		}
		arrTreeNode = append(arrTreeNode, root2)
		if root2.Left != nil {
			dfs(root2.Left)
		}
		if root2.Right != nil {
			dfs(root2.Right)
		}
	}
	dfs(root)
	for k, v := range arrTreeNode {
		v.Left = nil
		if k < len(arrTreeNode)-1 {
			v.Right = arrTreeNode[k+1]
		}
	}
}
func main() {
	a := TreeNode.CreateTree([]int{1, 2, 5, 3, 4, 6})
	flatten(a)
	TreeNode.PrintListNode(a)
}
