package main

import (
	"log"
)

// 给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。

//

// 示例 1：

// 输入: s = "leetcode"
// 输出: 0
// 示例 2:

// 输入: s = "loveleetcode"
// 输出: 2
// 示例 3:

// 输入: s = "aabb"
// 输出: -1

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn5z8r/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	s := "leetcode"
	log.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	tmpStr := make(map[rune]int, len(s))
	for _, v := range s {
		tmpStr[v]++
	}

	for k, v := range s {
		if tmpStr[v] == 1 {
			return k
		}
	}
	return -1
}
