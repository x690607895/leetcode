package main

import "log"

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
	a := createTree([]int{1, 2, 2, 3, 4, 4, 3})
	log.Fatal(levelOrder(a))
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

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	var getData func(root *TreeNode, depath int)
	getData = func(root *TreeNode, depath int) {
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
