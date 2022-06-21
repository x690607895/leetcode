package main

import "TreeNode"

// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

// 示例 1:

// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
// 示例 2:

// 输入: preorder = [-1], inorder = [-1]
// 输出: [-1]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvix0d/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	TreeNode.PrintListNode(buildTree(preorder, inorder))
}

func buildTree(preorder, inorder []int) *TreeNode.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode.TreeNode{preorder[0], nil, nil}

	index := 0
	for k, v := range inorder {
		if v == preorder[0] {
			index = k
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}
