package main

import (
	"ListNode"
	"log"
)

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

// 示例 1：

// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
// 示例 2：

// 输入：head = [1], n = 1
// 输出：[]
// 示例 3：

// 输入：head = [1,2], n = 1
// 输出：[1]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn2925/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	a := ListNode.CreateList([]int{1, 2, 3, 4, 5})
	b := removeNthFromEnd(a, 2)
	for ; b != nil; b = b.Next {
		log.Println(b)
	}
}

func removeNthFromEnd(head *ListNode.ListNode, n int) *ListNode.ListNode {
	index := getLength(head) - n
	// 剪头
	if index == 0 {
		return head.Next
	}

	result := head
	for i := 1; i < index; i++ {
		result = result.Next
	}

	result.Next = result.Next.Next

	return head
}
func getLength(head *ListNode.ListNode) int {
	result := 0
	for ; head != nil; head = head.Next {
		result++

	}
	return result
}
