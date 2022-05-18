package main

import "log"

// 给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。

// 示例 1：

// 输入：mat = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,4,7,5,3,6,8,9]
// 示例 2：

// 输入：mat = [[1,2],[3,4]]
// 输出：[1,2,3,4]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/cuxq3/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	log.Println(findDiagonalOrder(mat))
}

func findDiagonalOrder(mat [][]int) []int {
	M := len(mat)
	N := len(mat[0])
	temp := make([][]int, M+N)
	for k1, v1 := range mat {
		for k2, v2 := range v1 {
			temp[k1+k2] = append(temp[k1+k2], v2)
		}
	}

	result := make([]int, 0)
	for i := 0; i < M+N; i++ {
		if i%2 == 0 && len(temp[i]) > 1 {
			tempLen := len(temp[i])
			for j := 0; j < tempLen/2; j++ {
				temp[i][j], temp[i][tempLen-1-j] = temp[i][tempLen-1-j], temp[i][j]
			}
		}
		result = append(result, temp[i]...)
	}
	return result
}
