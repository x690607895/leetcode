package main

import "log"

// 编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。

//

// 示例 1：

// 输入：
// [
//   [1,1,1],
//   [1,0,1],
//   [1,1,1]
// ]
// 输出：
// [
//   [1,0,1],
//   [0,0,0],
//   [1,0,1]
// ]
// 示例 2：

// 输入：
// [
//   [0,1,2,0],
//   [3,4,5,2],
//   [1,3,1,5]
// ]
// 输出：
// [
//   [0,0,0,0],
//   [0,4,5,0],
//   [0,3,1,0]
// ]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/ciekh/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	martix := [][]int{
		{0, 1},
	}
	setZeroes(martix)
	log.Println(martix)
}

func setZeroes(matrix [][]int) {
	M := len(matrix)
	N := len(matrix[0])
	zeroArr := make([][]int, 0)
	for k1, v1 := range matrix {
		for k2, v2 := range v1 {
			if v2 == 0 {
				zeroArr = append(zeroArr, []int{k1, k2})
			}
		}
	}

	for _, v := range zeroArr {
		x, y := v[0], v[1]
		for i := 0; i < N; i++ {
			matrix[x][i] = 0
		}
		for i := 0; i < M; i++ {
			matrix[i][y] = 0
		}
	}
}
