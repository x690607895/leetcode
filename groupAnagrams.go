package main

import (
	"log"
	"sort"
)

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

// 示例 1:

// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
// 示例 2:

// 输入: strs = [""]
// 输出: [[""]]
// 示例 3:

// 输入: strs = ["a"]
// 输出: [["a"]]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvaszc/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	strs := []string{"", "b"}
	log.Println(groupAnagrams(strs), 'a', 'z')
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}

	temp := make(map[string][]string)
	for _, v := range strs {
		temp2 := []byte(v)
		sort.Slice(temp2, func(i, j int) bool {
			return temp2[i] < temp2[j]
		})
		s := string(temp2)
		if temp[s] == nil {
			temp[s] = make([]string, 0)
		}
		temp[s] = append(temp[s], v)
	}

	result := make([][]string, 0)
	for _, v := range temp {
		result = append(result, v)
	}
	return result
}
