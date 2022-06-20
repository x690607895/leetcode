package main

import (
	"ListNode"
)

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

// 请你将两个数相加，并以相同形式返回一个表示和的链表。

// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例 1：

// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 示例 2：

// 输入：l1 = [0], l2 = [0]
// 输出：[0]
// 示例 3：

// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvw73v/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	l1 := ListNode.CreateList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := ListNode.CreateList([]int{9, 9, 9, 9})
	ListNode.PrintListNode(addTwoNumbers(l1, l2))
}
func addTwoNumbers(l1 *ListNode.ListNode, l2 *ListNode.ListNode) *ListNode.ListNode {
	result := &ListNode.ListNode{}
	cur := result
	isJinWei := false
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if isJinWei {
			sum++
		}

		sum, isJinWei = sum%10, sum >= 10
		cur.Next = &ListNode.ListNode{sum, nil}
		cur = cur.Next
	}

	if isJinWei {
		cur.Next = &ListNode.ListNode{1, nil}
	}
	return result.Next
}
