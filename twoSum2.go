package main

import (
	"log"
)

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

// 你可以按任意顺序返回答案。

// 示例 1：

// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
// 示例 2：

// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
// 示例 3：

// 输入：nums = [3,3], target = 6
// 输出：[0,1]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2jrse/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	numbers := []int{1, 2, 3, 4, 4, 9, 56, 90}
	target := 8
	log.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers-1; i++ {
		for j := i + 1; j < lenNumbers-1; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
