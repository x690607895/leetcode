package ListNode

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(src []int) *ListNode {
	result := &ListNode{src[0], nil}
	temp := result
	for i := 1; i < len(src); i++ {
		temp.Next = &ListNode{src[i], nil}
		temp = temp.Next
	}
	return result
}

func PrintListNode(src *ListNode) {
	for ; src != nil; src = src.Next {
		log.Println(src.Val)
	}
}
