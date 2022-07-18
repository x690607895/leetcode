package main

import "log"

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 示例 1：

// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
// 示例 2：

// 输入：n = 1
// 输出：["()"]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xv33m7/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func main() {
	log.Println(generateParenthesis(3))
}
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	var temp func(l, r int, str string)
	temp = func(l, r int, str string) {
		if l == 0 && r == 0 {
			result = append(result, str)
			return
		}
		if r < l || l < 0 {
			return
		}
		temp(l-1, r, str+"(")
		temp(l, r-1, str+")")
	}

	temp(n, n, "")
	return result
}
