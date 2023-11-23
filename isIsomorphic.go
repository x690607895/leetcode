package main

import "fmt"

// 给定两个字符串 s 和 t ，判断它们是否是同构的。

// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

// 思路
// 建立两个hash表，循环遍历S，交叉比对，即一个映射表里的k对应的v和另一个字符串不一样则推出

func isIsomorphic(s, t string) bool {
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}

func main() {
	// fmt.Println(isIsomorphic("egg", "add"))
	// fmt.Println(isIsomorphic("foo", "bar"))
	// fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
}
