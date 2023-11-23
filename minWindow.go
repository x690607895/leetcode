package main

import "fmt"

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

// 注意：

// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	if len(t) == 1 {
		for _, v := range s {
			if v == rune(t[0]) {
				return t
			}
		}
	}
	mapWords := make(map[rune]int)
	wordsCount := 0
	for _, v := range t {
		if _, ok := mapWords[v]; !ok {
			wordsCount++
		}
		mapWords[v]++
	}
	l := 0
	tmp := map[rune]int{}
	tmpCount := 0
	res := ""
	for r, v := range s {
		if _, ok := mapWords[v]; ok {
			tmp[v]++
			if tmp[v] == mapWords[v] {
				tmpCount++
			}
		}

		if tmpCount == wordsCount {
			if r-l+1 < len(res) || res == "" {
				res = s[l : r+1]
			}
			for ; l < r && tmpCount == wordsCount; l++ {
				tmpS := rune(s[l])
				if _, ok := tmp[tmpS]; ok {
					tmp[tmpS]--
					if tmp[tmpS]+1 == mapWords[tmpS] {
						tmpCount--
					}
				}
				if (r-l < len(res) || res == "") && tmpCount == wordsCount {
					res = s[l+1 : r+1]
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(minWindow("abc", "bc"))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("bba", "ab"))
}
