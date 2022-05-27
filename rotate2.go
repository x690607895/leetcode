package main

import "log"

// 给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

// 示例 1:

// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4]
// 示例 2:

// 输入：nums = [-1,-100,3,99], k = 2
// 输出：[3,99,-1,-100]
// 解释:
// 向右轮转 1 步: [99,-1,-100,3]
// 向右轮转 2 步: [3,99,-1,-100]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2skh7/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// k := 3
	nums := []int{1, 2}
	k := 2
	rotate(nums, k)
	log.Println(nums)
}

func rotate(nums []int, k int) {
	num2 := make([]int, len(nums))
	for i, v := range nums {
		num2[(i+k)%len(nums)] = v
	}
	copy(nums, num2)
}
