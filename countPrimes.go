package main

import "log"

// 给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。

// 示例 1：

// 输入：n = 10
// 输出：4
// 解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
// 示例 2：

// 输入：n = 0
// 输出：0
// 示例 3：

// 输入：n = 1
// 输出：0

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnzlu6/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	n := 10
	log.Println(countPrimes(n))
}

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}

	result := make([]bool, n)
	num := 0
	for i := 2; i < n; i++ {
		if result[i] == true {
			continue
		}
		num++
		for j := 2 * i; j < n; j += i {
			result[j] = true
		}
	}
	return num
}
