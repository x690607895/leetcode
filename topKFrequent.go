package main

import (
	"github.com/go-acme/lego/log"
)

// 前 K 个高频元素
// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

// 示例 1:

// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
// 示例 2:

// 输入: nums = [1], k = 1
// 输出: [1]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvzpxi/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	nums := []int{1, 1, 1, 2, 2, 3333}
	k := 2
	log.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	arrayValues := make(map[int]int)
	for _, v := range nums {
		arrayValues[v]++
	}

	result := []int{}
	for value, times := range arrayValues {
		if len(result) == 0 || times == 1 {
			result = append(result, value)
		} else {
			lenResult := len(result)
			if times >= arrayValues[result[0]] {
				result = append([]int{value}, result...)
			} else if times <= arrayValues[result[len(result)-1]] {
				result = append(result, value)
			} else {
				for i := 0; i < lenResult; i++ {
					if times >= arrayValues[result[i]] {
						result = append(result[:i], append([]int{value}, result[i:]...)...)
						break
					} else if i == lenResult-1 {
						result = append(result, value)
					}
				}
			}
		}
	}
	return result
}
