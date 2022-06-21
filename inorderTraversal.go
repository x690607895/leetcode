package main

import (
	"TreeNode"
	"log"
)

// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

// 示例 1：

// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
// 示例 2：

// 输入：root = []
// 输出：[]
// 示例 3：

// 输入：root = [1]
// 输出：[1]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xv7pir/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func main() {
	root := TreeNode.CreateTree([]int{1, -1, 2, -1, -1, 3})
	log.Println(inorderTraversal(root))
	inorderTraversal2(root)
	// TreeNode.PrintListNode(root)
}

func inorderTraversal(root *TreeNode.TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	temp := []*TreeNode.TreeNode{}

	for root != nil || len(temp) > 0 {
		for root != nil {
			temp = append(temp, root)
			root = root.Left
		}
		root = temp[len(temp)-1]
		temp = temp[:len(temp)-1]
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}
func inorderTraversal2(root *TreeNode.TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	var getData func(root *TreeNode.TreeNode)
	getData = func(root *TreeNode.TreeNode) {
		if root.Left != nil {
			getData(root.Left)
		}
		result = append(result, root.Val)
		if root.Right != nil {
			getData(root.Right)
		}
	}
	getData(root)
	return result
}
