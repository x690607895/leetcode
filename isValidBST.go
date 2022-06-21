package main

import (
	"TreeNode"
	"log"
	"math"
)

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

// 有效 二叉搜索树定义如下：

// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。

// 示例 1：

// 输入：root = [2,1,3]
// 输出：true
// 示例 2：

// 输入：root = [5,1,4,null,null,3,6]
// 输出：false
// 解释：根节点的值是 5 ，但是右子节点的值是 4 。

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn08xg/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func main() {
	a := TreeNode.CreateTree([]int{5, 4, 6, -1, -1, 3, 7})
	log.Fatal(isValidBST(a))
}

func isValidBST(root *TreeNode.TreeNode) bool {
	return temp2(root, math.MinInt64, math.MaxInt64)
}

func temp2(root *TreeNode.TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return temp2(root.Left, min, root.Val) && temp2(root.Right, root.Val, max)
}
