package main

import (
	"TreeNode"
	"log"
)

// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

// 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

//

// 示例 1：

// 输入：nums = [-10,-3,0,5,9]
// 输出：[0,-3,9,-10,null,5]
// 解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

// 示例 2：

// 输入：nums = [1,3]
// 输出：[3,1]
// 解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xninbt/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func main() {
	a := sortedArrayToBST([]int{1, 2, 3, 4, 5})
	log.Fatal(a)
}

func sortedArrayToBST(nums []int) *TreeNode.TreeNode {
	return insert(nums)
}

func insert(nums []int) *TreeNode.TreeNode {
	mid := len(nums) / 2
	if len(nums) == 1 {
		return &TreeNode.TreeNode{
			Val: nums[0],
		}
	} else if len(nums) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   nums[mid],
		Left:  insert(nums[:mid]),
		Right: insert(nums[mid+1:]),
	}
}
