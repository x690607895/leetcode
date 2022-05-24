package main

import "log"

// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。

// 示例 1:

// 输入: numRows = 5
// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
// 示例 2:

// 输入: numRows = 1
// 输出: [[1]]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/cuj3m/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	numRows := 7
	log.Println(generate(numRows))
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j-1 < 0 || j >= i {
				tmp[j] = 1
			} else {
				tmp[j] = result[i-1][j-1] + result[i-1][j]
			}
		}
		result[i] = tmp
	}
	return result
}
