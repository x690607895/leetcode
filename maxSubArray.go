package main

import (
	"log"
	"sort"
)

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

// 子数组 是数组中的一个连续部分。

// 示例 1：

// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
// 示例 2：

// 输入：nums = [1]
// 输出：1
// 示例 3：

// 输入：nums = [5,4,-1,7,8]
// 输出：23

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn3cg3/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	log.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	sum := make([]int, len(nums))
	for k, v := range nums {
		if k == 0 {
			sum[k] = v
		} else {
			sum[k] = max(sum[k-1], 0) + v
		}
	}
	sort.Ints(sum)
	return sum[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
