package main

import "fmt"

// 「快乐数」 定义为：

// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

func isHappy(n int) bool {
	mapN := map[int]bool{}
	for n != 1 && !mapN[n] {
		mapN[n] = true
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
	}
	return n == 1
}

func main() {
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(19))
}
