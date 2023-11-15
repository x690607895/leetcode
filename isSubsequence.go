package main

import "fmt"

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

// 用双指针不要用穷举

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	lenS, lenT := len(s), len(t)
	sIndex, tIndex := 0, 0
	for sIndex < lenS && tIndex < lenT {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		tIndex++
	}
	return sIndex == lenS
}

func main() {
	fmt.Println(isSubsequence("xcf", "acbedcfg"))
}
