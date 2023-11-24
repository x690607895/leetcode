package main

import "fmt"

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

func spiralOrder(matrix [][]int) []int {
	direction := "right"
	used := map[int]struct{}{}
	x, y := 0, 0
	row, col := len(matrix), len(matrix[0])
	result := []int{}
	for col*row != len(used) {
		result = append(result, matrix[x][y])
		mapIndex := x*10 + y
		used[mapIndex] = struct{}{}
		switch direction {
		case "up":
			x--
			break
		case "down":
			x++
			break
		case "left":
			y--
			break
		case "right":
			y++
			break
		}

		mapIndex = x*10 + y
		_, ok := used[mapIndex]
		switch direction {
		case "up":
			if x < 0 || ok {
				x++
				y++
				direction = "right"
			}
			break
		case "down":
			if x >= row || ok {
				x--
				y--
				direction = "left"
			}
			break
		case "left":
			if y < 0 || ok {
				y++
				x--
				direction = "up"
			}
			break
		case "right":
			if y >= col || ok {
				y--
				x++
				direction = "down"
			}
			break
		}
	}
	return result
}

func main() {
	fmt.Println(spiralOrder([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}))
}
