package main

import "log"

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 示例 1：

// 输入：n = 2
// 输出：2
// 解释：有两种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶
// 2. 2 阶
// 示例 2：

// 输入：n = 3
// 输出：3
// 解释：有三种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶 + 1 阶
// 2. 1 阶 + 2 阶
// 3. 2 阶 + 1 阶

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn854d/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	log.Println(climbStairs(5))
}

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	temp := make([]int, n+1)
	var temp2 func(n int) int
	temp2 = func(n int) int {
		if temp[n] != 0 {
			return temp[n]
		}
		value := 0
		if n <= 1 {
			value = 1
		} else if n == 2 {
			value = 2
		} else {
			value = temp2(n-2) + temp2(n-1)
		}
		temp[n] = value
		return temp[n]
	}

	return temp2(n)
}
