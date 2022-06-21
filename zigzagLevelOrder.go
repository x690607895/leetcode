package main

import (
	"TreeNode"
	"log"
)

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

// 示例 1：

// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
// 示例 2：

// 输入：root = [1]
// 输出：[[1]]
// 示例 3：a

// 输入：root = []
// 输出：[]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvle7s/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	root := TreeNode.CreateTree([]int{1, 2, 3, 4, -1, -1, 5})
	log.Println(zigzagLevelOrder(root))
}

func zigzagLevelOrder(root *TreeNode.TreeNode) [][]int {
	result := make([][]int, 0)
	deepath := 0
	var temp func(root *TreeNode.TreeNode, deepth int)
	temp = func(root *TreeNode.TreeNode, deepth int) {
		if root == nil {
			return
		}
		if len(result)-1 < deepth {
			result = append(result, []int{})
		}
		if deepth%2 != 1 {
			result[deepth] = append(result[deepth], root.Val)
		} else {
			result[deepth] = append([]int{root.Val}, result[deepth]...)
		}
		deepth++
		temp(root.Left, deepth)
		temp(root.Right, deepth)
	}
	temp(root, deepath)
	return result
}
