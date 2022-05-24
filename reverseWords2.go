package main

import (
	"log"
	"strings"
)

// 给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

// 示例 1：

// 输入：s = "Let's take LeetCode contest"
// 输出："s'teL ekat edoCteeL tsetnoc"
// 示例 2:

// 输入： s = "God Ding"
// 输出："doG gniD"

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/c8su7/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	s := "Let's take LeetCode contest"
	log.Println(reverseWords(s))
}

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	result := make([]string, 0)
	for _, v := range arr {
		lenV := len(v)
		tmp := make([]rune, lenV)
		lenV--
		for k2, v2 := range v {
			tmp[lenV-k2] = v2
		}
		result = append(result, string(tmp))
	}
	return strings.Join(result, " ")
}
