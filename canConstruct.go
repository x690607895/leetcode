package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	mapWords := map[rune]int{}
	wordsN := 0
	for _, v := range ransomNote {
		if _, ok := mapWords[v]; !ok {
			wordsN++
		}
		mapWords[v]++
	}
	for _, v := range magazine {
		if _, ok := mapWords[v]; ok {
			mapWords[v]--
			if mapWords[v] == 0 {
				wordsN--
			}
		}
	}
	return wordsN == 0
}

func main() {
	fmt.Println(canConstruct("ab", "aab"))
	fmt.Println(canConstruct("ac", "aab"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aab", "baa"))
}
