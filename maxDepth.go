package main

import "log"

// 给定一个二叉树，找出其最大深度。

// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

// 说明: 叶子节点是指没有子节点的节点。

// 示例：
// 给定二叉树 [3,9,20,null,null,15,7]，

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回它的最大深度 3

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnd69e/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	a := createTree([]int{3, 9, 20, -1, -1, 15, 7})
	log.Fatal(maxDepth(a))
}

func createTree(src []int) *TreeNode {
	tempArr := make([]*TreeNode, len(src))
	for k, v := range src {
		if v == -1 {
			tempArr[k] = nil
		} else {
			tempArr[k] = &TreeNode{v, nil, nil}
		}
	}
	for i := len(src) - 1; i > 0; i-- {
		if i%2 == 1 {
			k := i / 2
			tempArr[k].Left = tempArr[i]

		} else {
			k := i/2 - 1
			tempArr[k].Right = tempArr[i]
		}

	}
	log.Println(tempArr)
	return tempArr[0]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
