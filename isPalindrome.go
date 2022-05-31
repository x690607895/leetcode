package main

import "log"

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

// 说明：本题中，我们将空字符串定义为有效的回文串。

// 示例 1:

// 输入: "A man, a plan, a canal: Panama"
// 输出: true
// 解释："amanaplanacanalpanama" 是回文串
// 示例 2:

// 输入: "race a car"
// 输出: false
// 解释："raceacar" 不是回文串

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xne8id/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	s := "A man, a plan, a canal: Panama"
	s = "OP"
	// s = "ababa"
	log.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	tmpByte := make([]rune, 0)
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			tmpByte = append(tmpByte, v)
		}
		if v >= 'A' && v <= 'Z' {
			v += 32
			tmpByte = append(tmpByte, v)
		}
	}

	r := len(tmpByte) - 1
	for i := 0; i < r; i++ {
		if tmpByte[i] != tmpByte[r] {
			return false
		}
		r--
	}
	return true
}
