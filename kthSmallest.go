package main

import (
	"TreeNode"
	"log"
)

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

// 示例 1：

// 输入：root = [3,1,4,null,2], k = 1
// 输出：1
// 示例 2：

// 输入：root = [5,3,6,2,4,null,null,1], k = 3
// 输出：3

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvuyv3/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	root := TreeNode.CreateTree([]int{5, 3, 6, 2, 4, -1, -1, 1})
	log.Println(kthSmallest(root, 3))
}

func kthSmallest(root *TreeNode.TreeNode, k int) int {
	result := 0
	times := 0
	var getMin func(root *TreeNode.TreeNode)
	getMin = func(root *TreeNode.TreeNode) {
		if k == 0 {
			return
		}
		times++

		if root.Left != nil {
			getMin(root.Left)
		}
		k--
		if k == 0 {
			result = root.Val
			return
		}
		if root.Right != nil {
			getMin(root.Right)
		}
	}
	getMin(root)
	log.Println(times)
	return result
}
