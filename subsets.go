package main

import (
	"log"
)

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// 示例 1：

// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 示例 2：

// 输入：nums = [0]
// 输出：[[],[0]]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xv67o6/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func main() {
	nums := []int{1, 2, 3}
	log.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	var temp func(data []int, last []int)
	temp = func(data []int, last []int) {
		result = append(result, data)
		if len(last) == 0 {
			return
		}

		for k, v := range last {
			data = append(data, v)
			temp(data, last[k+1:])
		}
	}
	result = append(result, []int{})
	for k, v := range nums {
		temp([]int{v}, nums[k+1:])
	}
	return result
}

func subsets(nums []int) [][]int {
	res := [][]int{}

	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		if i == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		list = append(list, nums[i])
		dfs(i+1, list)
		list = list[:len(list)-1]
		dfs(i+1, list)
	}
	dfs(0, []int{})

	return res
}
