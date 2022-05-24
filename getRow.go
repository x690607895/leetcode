package main

import "log"

// 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。

// 示例 1:

// 输入: rowIndex = 3
// 输出: [1,3,3,1]
// 示例 2:

// 输入: rowIndex = 0
// 输出: [1]
// 示例 3:

// 输入: rowIndex = 1
// 输出: [1,1]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/ctyt1/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	rowIndex := 4
	log.Println(getRow(rowIndex))
}

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 0; i <= rowIndex; i++ {
		prev := 1
		for j := 0; j <= i; j++ {
			if j == 0 || j >= i {
				result[j] = 1
			} else {
				prev2 := result[j]
				result[j] += prev
				prev = prev2
			}
		}
	}
	return result
}
