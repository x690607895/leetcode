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

func main() {
	target := 6
	nums := []int{1, 2, 3, 4, 5}
	log.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	sum := 0
	for _, v := range nums {
		if v >= target {
			return 1
		}
		sum += v
	}
	if sum < target {
		return 0
	}

	result := 9999
	lNums := len(nums)
	for i := 0; i < lNums; i++ {
		sum := nums[i]
		for j := i + 1; j < lNums; j++ {
			sum += nums[j]
			if sum >= target {
				if (j - i + 1) < result {
					result = j - i + 1
				}
				break
			}
		}
	}

	if result == 9999 {
		result = 0
	}

	return result
}
