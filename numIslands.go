package main

import "log"

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

// 此外，你可以假设该网格的四条边均被水包围。

// 示例 1：

// 输入：grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// 输出：1
// 示例 2：

// 输入：grid = [
//   ["1","1","0","0","0"],
//   ["1","1","0","0","0"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// 输出：3

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvtsnm/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func main() {
	grid := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	log.Println(numIslands(grid))
}
func numIslands(grid [][]byte) int {
	result := 0

	for x, v := range grid {
		for y, v2 := range v {
			if v2 == '1' {
				result++
				dfs(grid, x, y)
			}

		}
	}
	return result
}

func dfs(grid [][]byte, x, y int) {
	if x < 0 || y < 0 || x > len(grid)-1 || y > len(grid[x])-1 {
		return
	}
	if grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	dfs(grid, x-1, y)
	dfs(grid, x+1, y)
	dfs(grid, x, y+1)
	dfs(grid, x, y-1)
}
