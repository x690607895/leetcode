package main

import "log"

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

// 示例 1：

// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 示例 2：

// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
// 示例 3：

// 输入：nums = [1]
// 输出：[[1]]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvqup5/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	nums := []int{1, 2, 3}
	log.Println(permute(nums))
}
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	lenNums := len(nums)
	var temp func(start int, data []int)
	temp = func(start int, data []int) {
		if start == lenNums {
			result = append(result, data)
			return
		}
		start++
		for _, v := range nums {
			if len(data) == 0 || !in_array(v, data) {
				temp(start, append(data, v))
			}
		}
	}
	temp(0, []int{})
	return result
}

func in_array(needle int, src []int) bool {
	for _, v := range src {
		if v == needle {
			return true
		}
	}
	return false
}
