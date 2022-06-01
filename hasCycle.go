package main

import "log"

// 给你一个链表的头节点 head ，判断链表中是否有环。

// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

// 如果链表中存在环 ，则返回 true 。 否则，返回 false 。

// 示例 1：

// 输入：head = [3,2,0,-4], pos = 1
// 输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。
// 示例 2：

// 输入：head = [1,2], pos = 0
// 输出：true
// 解释：链表中有一个环，其尾部连接到第一个节点。
// 示例 3：

// 输入：head = [1], pos = -1
// 输出：false
// 解释：链表中没有环。

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnwzei/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	a := createList([]int{1, 2, 3})
	log.Println(hasCycle(a))
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

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	temp := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := temp[head]; ok {
			return true
		}
		temp[head] = struct{}{}
		head = head.Next
	}
	log.Println(temp)
	return false
}
