package main

import "log"

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// 示例 1：

// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true
// 示例 2：

// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// 输出：true
// 示例 3：

// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
// 输出：false

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvkwe2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	board := [][]byte{{'C', 'A', 'A', 'E'}, {'A', 'A', 'A', 'S'}, {'B', 'D', 'E', 'E'}}
	word := "AAB"

	log.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	m, n, o := len(board), len(board[0]), len(word)
	word2 := []byte(word)
	var posIsExists func(x, y, index int) bool
	posIsExists = func(x, y, index int) bool {
		if index == o {
			return true
		}

		if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == word2[index] {
			board[x][y] += 100
			index++
			f1 := posIsExists(x-1, y, index) || posIsExists(x+1, y, index) || posIsExists(x, y+1, index) || posIsExists(x, y-1, index)
			board[x][y] -= 100
			return f1
		}
		return false
	}
	for x, v := range board {
		for y := range v {
			if posIsExists(x, y, 0) {
				return true
			}
		}
	}
	return false
}
