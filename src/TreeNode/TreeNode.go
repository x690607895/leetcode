package TreeNode

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(src []int) *TreeNode {
	tempArr := make([]*TreeNode, len(src))
	for k, v := range src {
		if v == -1 {
			tempArr[k] = nil
		} else {
			tempArr[k] = &TreeNode{v, nil, nil}
		}
	}
	for i := len(src) - 1; i > 0; i-- {
		if tempArr[i] == nil {
			continue
		}
		if i%2 == 1 {
			k := i / 2
			tempArr[k].Left = tempArr[i]

		} else {
			k := i/2 - 1
			tempArr[k].Right = tempArr[i]
		}

	}
	return tempArr[0]
}

func PrintListNode(src *TreeNode) {
	result := []*TreeNode{src}
	for len(result) > 0 {
		log.Println(result[0].Val)
		if result[0].Left != nil {
			result = append(result, result[0].Left)
		}
		if result[0].Right != nil {
			result = append(result, result[0].Right)
		}
		result = result[1:]
	}
}
