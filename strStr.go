package main

import "log"

// 实现 strStr() 函数。

// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

// 说明：

// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

// 示例 1：

// 输入：haystack = "hello", needle = "ll"
// 输出：2
// 示例 2：

// 输入：haystack = "aaaaa", needle = "bba"
// 输出：-1

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/cm5e2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	haystack := "ABABBBED_ABABACCD"
	needle := "ABCDABD"
	needle = "ABABACCD"
	log.Println(strStr(haystack, needle))
}

//  A B A B A C C D
// [0 1 2 3 4 5 6 7]
// [0 1 1 2 3 4 1 1]
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	// kmp
	// 制作next数组
	// 0-n中从j开始截取字符串
	nextArr := make([]int, len(needle))
	for i := 0; i <= len(needle)-1; i++ {
		if i == 0 {
			nextArr[i] = 0
			continue
		}
		mapMaxStr := make(map[string]int)
		tmp := needle[:i+1]
		for j := 0; j < i+1; j++ {
			mapMaxStr[tmp[:j]] += 1
			mapMaxStr[tmp[j:]] += 1
		}
		nextArr[i] = getMaxStr(mapMaxStr)
	}

	// 双指针
	for up, down := 0, 0; up < len(haystack); up++ {
		for down > 0 && haystack[up] != needle[down] {
			down = nextArr[down-1]
		}
		if haystack[up] == needle[down] {
			down++
		}
		if up == len(haystack)-1 {
			return up - len(needle) + 1
		}
	}
	return -1
}

func getMaxStr(source map[string]int) (result int) {
	max := 0
	for str, times := range source {
		if times == 1 {
			continue
		}
		if times == max && len(str) > result {
			result = len(str)
		}
		if times > max {
			result = len(str)
			max = times
		}
	}
	return
}
