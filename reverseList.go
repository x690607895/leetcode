package main

import (
	"ListNode"
	"log"
)

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//

// 示例 1：

// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
// 示例 2：

// 输入：head = [1,2]
// 输出：[2,1]
// 示例 3：

// 输入：head = []
// 输出：[]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnnhm6/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	a := ListNode.CreateList([]int{1, 2, 3, 4, 5})
	b := reverseList(a)
	for ; b != nil; b = b.Next {
		log.Println(b)
	}
}

func reverseList(head *ListNode.ListNode) *ListNode.ListNode {
	result := &ListNode.ListNode{}

	for ; head != nil; head = head.Next {
		result.Val = head.Val
		temp := &ListNode.ListNode{0, result}
		result = temp
	}
	return result.Next
}
