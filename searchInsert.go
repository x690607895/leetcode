package main

import (
	"log"
	"math"
)

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

// 请必须使用时间复杂度为 O(log n) 的算法。

// 示例 1:

// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
// 示例 2:

// 输入: nums = [1,3,5,6], target = 2
// 输出: 1
// 示例 3:

// 输入: nums = [1,3,5,6], target = 7
// 输出: 4

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/cxqdh/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	a := []int{2, 3, 5, 6, 9}
	b := 7
	log.Println(searchInsert(a, b), 10>>1)

}
func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}
	numsLen := len(nums)

	if target == nums[numsLen-1] {
		return numsLen - 1
	}
	if target > nums[numsLen-1] {
		return numsLen
	}

	harfNumsLen := numsLen / 2
	a := harfNumsLen
	middle := a
	for i := 0; i < harfNumsLen+1; i++ {
		a = int(math.Max(float64(a)/2, 1))
		if target < nums[middle-1] {
			middle = a
			continue
		} else if target == nums[middle-1] {
			return middle - 1
		} else if target > nums[middle-1] && target <= nums[middle] {
			return middle
		} else {
			middle += a
			continue
		}
	}
	return 0
}
