package main

import "log"

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

// 示例 1：

// 输入：head = [1,2,2,1]
// 输出：true
// 示例 2：

// 输入：head = [1,2]
// 输出：false

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnv1oc/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	a := createList([]int{1})
	log.Println(isPalindrome(a))
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

func isPalindrome(head *ListNode) bool {
	halfList := getList(head)
	halfReverseList := reverseList(halfList)
	for ; halfReverseList != nil; halfReverseList, head = halfReverseList.Next, head.Next {
		if halfReverseList.Val != head.Val {
			return false
		}
	}
	return true
}
func reverseList(head *ListNode) *ListNode {
	result := &ListNode{}

	for ; head != nil; head = head.Next {
		result.Val = head.Val
		temp := &ListNode{0, result}
		result = temp
	}
	return result.Next
}

func getList(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
