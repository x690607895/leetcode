package main

import (
	"log"
)

// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

//

// 示例 1：

// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：

// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/ceda1/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	strs := []string{"flower", "flow", "flight"}
	log.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 1 || len(strs[0]) < 1 {
		return strs[0]
	}
	minLength := 9999
	for _, v := range strs {
		minLength = min(minLength, len(v))
	}
	result := make([]byte, 0)
	for i := 0; i < minLength; i++ {
		str1 := strs[0][i]
		flag := true

		for k, v := range strs {
			if k == 0 {
				continue
			}
			if v[i] != str1 {
				flag = false
			}

		}
		if flag {
			result = append(result, str1)
		} else {
			break
		}
	}
	return string(result)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
