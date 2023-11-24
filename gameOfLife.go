package main

import "fmt"

// 根据 百度百科 ， 生命游戏 ，简称为 生命 ，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

// 给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

// 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
// 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
// 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
// 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
// 下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。给你 m x n 网格面板 board 的当前状态，返回下一个状态。

type direction struct {
	x, y int
	calc bool
}

func gameOfLife(board [][]int) {
	dir := []direction{
		direction{-1, -1, false},
		direction{-1, 0, false},
		direction{-1, 1, false},
		direction{0, -1, false},
		direction{0, 1, true},
		direction{1, -1, true},
		direction{1, 0, true},
		direction{1, 1, true},
	}
	row, col := len(board), len(board[0])
	// 只有活的会丢到列表里
	used := map[int]struct{}{}
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			alive := 0
			for _, v := range dir {
				curX, curY := x+v.x, y+v.y
				mapIndex := curX*100 + curY
				_, ok := used[mapIndex]
				if v.calc {
					if !ok && curX >= 0 && curX < row && curY >= 0 && curY < col && board[curX][curY] == 1 {
						used[mapIndex] = struct{}{}
						ok = true
					}
				}
				if ok {
					alive++
				}
				if alive > 3 {
					break
				}
			}
			if board[x][y] == 1 {
				used[x*100+y] = struct{}{}
				// 处理活着
				if alive != 2 && alive != 3 {
					board[x][y] = 0
				} else {
					board[x][y] = 1
				}
			} else {
				// 处理死亡
				if alive == 3 {
					board[x][y] = 1
				}
			}
		}
	}
}

func main() {
	a := [][]int{[]int{0, 1, 0}, []int{0, 0, 1}, []int{1, 1, 1}, []int{0, 0, 0}}
	gameOfLife(a)
	fmt.Println(a)
	a = [][]int{[]int{1, 1}, []int{1, 0}}
	gameOfLife(a)
	fmt.Println(a)
}
