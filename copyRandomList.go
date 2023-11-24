package main

import "fmt"

// 给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。

// 构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。

// 例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。

// 返回复制链表的头节点。

// 用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

// val：一个表示 Node.val 的整数。
// random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
// 你的代码 只 接受原链表的头节点 head 作为传入参数。

// 思路
// 先用一个数组记录每个元素
// 遍历数组确定random关系，同时生成一个新的数组并赋值和next关系
// 链接新数组random关系
// 返回数组第一个

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	arrNode := []*Node{}
	for head != nil {
		arrNode = append(arrNode, head)
		head = head.Next
	}
	arrLink := []int{}
	arrNode2 := []*Node{}
	times := 0
	for _, v := range arrNode {
		if v.Random == nil {
			arrLink = append(arrLink, -1)
		} else {
			for k2, v2 := range arrNode {
				if v.Random == v2 {
					arrLink = append(arrLink, k2)
				}
			}
		}

		tempNode := &Node{v.Val, nil, nil}
		if times >= 1 {
			arrNode2[times-1].Next = tempNode
		}
		arrNode2 = append(arrNode2, tempNode)
		times++
	}

	for k, v := range arrNode2 {
		if arrLink[k] != -1 {
			v.Random = arrNode2[arrLink[k]]
		}
	}
	return arrNode2[0]
}

func main() {
	a := []*Node{
		&Node{7, nil, nil},
		&Node{13, nil, nil},
		&Node{11, nil, nil},
		&Node{10, nil, nil},
		&Node{1, nil, nil},
	}
	for k, v := range a {
		if k+1 < len(a) {
			v.Next = a[k+1]
		}
	}
	a[1].Random = a[0]
	a[2].Random = a[4]
	a[3].Random = a[2]
	a[4].Random = a[0]
	fmt.Println(copyRandomList(a[0]))
}
