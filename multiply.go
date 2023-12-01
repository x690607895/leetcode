package main

import (
	"fmt"
)

// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

// 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。

func multiply(num1 string, num2 string) string {
	var m, n = len(num1), len(num2)
	var a1, a2 = []byte(num1), []byte(num2)
	for i := range a1 {
		a1[i] -= '0'
	}
	for i := range a2 {
		a2[i] -= '0'
	}
	var ret = make([]byte, m+n)
	for i := 0; i < m; i++ {
		d1 := a1[m-1-i]
		for j := 0; j < n; j++ {
			d2 := a2[n-1-j]
			ret[m+n-1-i-j] += d1 * d2
			for x := m + n - 1 - i - j; ret[x] > 9; x-- { //2)最难写的部分
				ret[x-1] += ret[x] / 10
				ret[x] %= 10
			}
		}
	}
	var start int
	for i := range ret { //1)标记前置0
		if start != m+n-1 && start == i && ret[i] == 0 {
			start++
		}
		ret[i] += '0'
	}
	return string(ret[start:])
}

func main() {
	fmt.Println(multiply("498828660196", "840477629533"))
}
