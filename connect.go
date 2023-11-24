package main

// 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

// struct Node {
//   int val;
//   Node *left;
//   Node *right;
//   Node *next;
// }
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

// 初始状态下，所有 next 指针都被设置为 NULL。

// 示例 1：

// 输入：root = [1,2,3,4,5,6,7]
// 输出：[1,#,2,3,#,4,5,6,7,#]
// 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。
// 示例 2:

// 输入：root = []
// 输出：[]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvijdh/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	root := CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	connect(root)
}

func CreateTree(src []int) *Node {
	tempArr := make([]*Node, len(src))
	for k, v := range src {
		tempArr[k] = &Node{v, nil, nil, nil}
	}
	for i := len(src) - 1; i > 0; i-- {
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

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		tmp := queue
		queue = nil

		for i, v := range tmp {
			if i < len(tmp)-1 {
				//打通同一层的节点
				v.Next = tmp[i+1]
			}
			// 左右数全部加入列表
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
	}
	return root
}
