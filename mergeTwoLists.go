package main

import "log"

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

//

// 示例 1：

// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// 示例 2：

// 输入：l1 = [], l2 = []
// 输出：[]
// 示例 3：

// 输入：l1 = [], l2 = [0]
// 输出：[0]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnnbp2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	a := createList([]int{1, 2, 3, 4, 5})
	c := createList([]int{1, 2, 6, 7})
	b := mergeTwoLists(a, c)
	for ; b != nil; b = b.Next {
		log.Println(b)
	}
}

func createList(src []int) *ListNode {
	result := &ListNode{src[0], nil}
	temp := result
	for i := 1; i < len(src); i++ {
		temp.Next = &ListNode{src[i], nil}
		temp = temp.Next
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			temp.Next = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
		}
		temp = temp.Next
	}
	if list1 == nil {
		temp.Next = list2
	}
	if list2 == nil {
		temp.Next = list1
	}
	return result.Next
}
