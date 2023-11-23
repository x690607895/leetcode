package main

import "fmt"

// 给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。

func wordPattern(pattern string, s string) bool {
	lPattern, lS := len(pattern), len(s)
	if lPattern >= lS {
		if lPattern == 1 && lS == 1 {
			if s[0] == ' ' {
				return false
			}
			return true
		}
		return false
	}
	mapWords, mapWords2 := map[byte]string{}, map[string]byte{}
	i := 0
	per := 0
	for k, v := range s {
		if v == ' ' || k == lS-1 {
			tmpS := s[per:k]
			if k == lS-1 {
				tmpS = s[per:]
			}
			per = k + 1
			if byt, ok := mapWords2[tmpS]; !ok {
				mapWords2[tmpS] = pattern[i]
			} else if byt != pattern[i] {
				return false
			}
			if value, ok := mapWords[pattern[i]]; !ok {
				mapWords[pattern[i]] = tmpS
			} else if value != tmpS {
				return false
			}
			i++
			if i == lPattern && k != lS-1 {
				return false
			}
		}
	}
	return i == lPattern
}

func main() {
	// fmt.Println(wordPattern("abba", "abcda"))
	// fmt.Println(wordPattern("a", "a"))
	// fmt.Println(wordPattern("a", " "))
	// fmt.Println(wordPattern("abba", "dog cat cat dog asd"))
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
	fmt.Println(wordPattern("abba", "dog cat cat asd"))
}
