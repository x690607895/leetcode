package main

import "log"

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

// 示例 1:

// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 示例 2:

// 输入: nums = [0]
// 输出: [0]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/c6u02/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	log.Println(nums)
}

func moveZeroes(nums []int) {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[slow] = nums[i]
			slow++
		}
	}

	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
}
