package main

import (
	"TreeNode"
	"log"
)

// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

//

// 示例 1：

// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
// 示例 2：

// 输入：root = [1]
// 输出：[[1]]
// 示例 3：

// 输入：root = []
// 输出：[]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnldjj/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func main() {
	a := TreeNode.CreateTree([]int{1, 2, 2, 3, 4, 4, 3})
	log.Fatal(levelOrder(a))
}

func levelOrder(root *TreeNode.TreeNode) [][]int {
	result := make([][]int, 0)
	var getData func(root *TreeNode.TreeNode, depath int)
	getData = func(root *TreeNode.TreeNode, depath int) {
		if root == nil {
			return
		}

		if len(result) < depath+1 {
			result = append(result, []int{})
		}
		result[depath] = append(result[depath], root.Val)
		depath++
		getData(root.Left, depath)
		getData(root.Right, depath)
	}
	getData(root, 0)
	return result
}
