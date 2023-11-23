package main

import (
	"fmt"
)

// 给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。

//  s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。

// 例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
// 返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。

// 思路
// 取出单个word的长度，后面获取字符串时用，记做：lPerWords
// 取出words的个数,记做：lAllWords
// 将words转换成map数组，key是当个word，value是出现的次数，记做：mapWords
// 取出words总的字符串长度，后面用来判断是否相同，记做：lWords
// 循环lPerWords（因为这里一跳就是lPerWords，所以必须lPerWords每个都检查到），开启滑动窗格，左边是从i开始，右边从i+lPerWords开始
// 循环字符串s，每次+lPerWords。如果在mapWords里，则本地备份数量+1，如果和mapWords数量相同，则本地count++
//当截取的字符串数量超过了总长度，进行修剪，开始动左边，取出l+lPerWords字符串，本地备份数量--，如果本地备份记录进过count的那count--，l++
// 如果本地count==lAllWords并且r-l=lWords，则标志找到了我们需要的l，进行记录

func findSubstring(s string, words []string) []int {
	lPerWords, lWords := len(words[0]), len(words)*len(words[0])
	lAllWords := 0
	result := []int{}
	mapWords := make(map[string]int)
	for _, v := range words {
		if _, ok := mapWords[v]; !ok {
			lAllWords++
		}
		mapWords[v]++
	}

	for i := 0; i < lPerWords; i++ {
		tmpWords := 0                       //记录有多少个单词
		tmpMapWords := make(map[string]int) //单词出现次数map
		l, r := i, i+lPerWords              //滑动窗格的左右两个指针
		for ; r <= len(s); r += lPerWords {
			tmpS := s[r-lPerWords : r]
			if _, ok := mapWords[tmpS]; ok {
				tmpMapWords[tmpS]++
				if tmpMapWords[tmpS] == mapWords[tmpS] {
					tmpWords++
				}
			}

			for r-l > lWords {
				tmpS = s[l : l+lPerWords]
				if _, ok := tmpMapWords[tmpS]; ok {
					tmpMapWords[tmpS]--
				}
				if tmpMapWords[tmpS]+1 == mapWords[tmpS] {
					tmpWords--
				}
				l += lPerWords
			}

			if tmpWords == lAllWords && r-l == lWords {
				result = append(result, l)
			}
		}
	}

	return result
}

func main() {
	fmt.Println(findSubstring("aaaaaaaaaaaaaa", []string{"aa", "aa"}))
}
