package main

import "log"

// 给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。

// 不占用额外内存空间能否做到？

//

// 示例 1:

// 给定 matrix =
// [
//   [1,2,3],
//   [4,5,6],
//   [7,8,9]
// ],

// 原地旋转输入矩阵，使其变为:
// [
//   [7,4,1],
//   [8,5,2],
//   [9,6,3]
// ]
// 示例 2:

// 给定 matrix =
// [
//   [ 5, 1, 9,11],
//   [ 2, 4, 8,10],
//   [13, 3, 6, 7],
//   [15,14,12,16]
// ],

// 原地旋转输入矩阵，使其变为:
// [
//   [15,13, 2, 5],
//   [14, 3, 4, 1],
//   [12, 6, 8, 9],
//   [16, 7,10,11]
// ]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/clpgd/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	martix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(martix)
	log.Println(martix)
}

func rotate(matrix [][]int) {
	// 先上下翻转
	// 再对角线翻转

	matrixLen := len(matrix)
	if matrixLen <= 1 {
		return
	}

	for i := 0; i < matrixLen/2; i++ {
		if i == matrixLen/2 && matrixLen/2%2 == 1 {
			break
		}
		matrix[i], matrix[matrixLen-1-i] = matrix[matrixLen-1-i], matrix[i]
	}

	for i := 0; i < matrixLen; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
