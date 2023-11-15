package main

import (
	"log"
)

// 给定一个含有 n 个正整数的数组和一个正整数 target 。

// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

// 示例 1：

// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 示例 2：

// 输入：target = 4, nums = [1,4,4]
// 输出：1
// 示例 3：

// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/c0w4r/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// 思路
// 双指针 l，r
// 每次移动R，并且加入总和
// 如果总和>=目标值
// 循环-L，直至小于目标，获取最小个数

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	log.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	result := -1
	lNums := len(nums)
	l, r := 0, 0
	sum := 0
	for r < lNums {
		sum += nums[r]
		minL := 0
		for sum >= target {
			minL = r - l + 1
			sum -= nums[l]
			l++
		}
		if minL > 0 && (result == -1 || minL < result) {
			result = minL
		}
		if minL == 1 {
			return 1
		}
		r++
	}

	if result == -1 {
		result = 0
	}

	return result
}
