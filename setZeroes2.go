package main

import "log"

// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

// 示例 1：

// 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
// 输出：[[1,0,1],[0,0,0],[1,0,1]]
// 示例 2：

// 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvmy42/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	log.Println(matrix)
}

func setZeroes(matrix [][]int) {
	// 列不动，行+=列数
	zeroArr := make([]int, 0)
	maxCol := len(matrix[0])
	for i, v := range matrix {
		for j, v2 := range v {
			if v2 == 0 {
				zeroArr = append(zeroArr, j)
				zeroArr = append(zeroArr, i+maxCol)
			}
		}
	}

	for _, v := range zeroArr {
		if v >= maxCol {
			v -= maxCol
			for k2 := range matrix[v] {
				matrix[v][k2] = 0
			}
		} else {
			for k2 := range matrix {
				matrix[k2][v] = 0
			}
		}
	}

}
