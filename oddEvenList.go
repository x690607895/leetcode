package main

import "ListNode"

// 给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。

// 第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。

// 请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。

// 你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。

// 示例 1:

// 输入: head = [1,2,3,4,5]
// 输出: [1,3,5,2,4]
// 示例 2:

// 输入: head = [2,1,3,5,6,4,7]
// 输出: [2,3,6,7,1,5,4]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvdwtj/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	head := ListNode.CreateList([]int{2, 1, 3, 5, 6, 4, 7})
	ListNode.PrintListNode(oddEvenList(head))
}

func oddEvenList(head *ListNode.ListNode) *ListNode.ListNode {
	odd, evenHead := head, head.Next
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
