package main

import (
	"github.com/go-acme/lego/log"
)

// 整数数组 nums 按升序排列，数组中的值 互不相同 。

// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

// 示例 1：

// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4
// 示例 2：

// 输入：nums = [4,5,6,7,0,1,2], target = 3
// 输出：-1
// 示例 3：

// 输入：nums = [1], target = 0
// 输出：-1

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvyz1t/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	nums := []int{1, 2, 4, 5, 6, 7, 0}
	target := 0
	log.Println(search(nums, target))
}

/**
二分查找法
每次先找中间，
相等返回
不想等
当中间比最右边小的时候，并且中间到最右边是一个递增队列的话，左边移到中间后，否则右边移到中间前
**/
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r + l) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] <= nums[r] {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if target >= nums[l] && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
