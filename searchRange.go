package main

import "log"

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

// 如果数组中不存在目标值 target，返回 [-1, -1]。

// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

// 示例 1：

// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]
// 示例 2：

// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]
// 示例 3：

// 输入：nums = [], target = 0
// 输出：[-1,-1]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xv4bbv/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	nums := []int{-1}
	target := 0
	log.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	mid := len(nums) / 2
	start, end := -1, -1
	if target < nums[mid] {
		for i := mid - 1; i >= 0; i-- {
			if nums[i] == target {
				if end == -1 {
					end = i
				}
				start = i
			}
		}
	} else if target == nums[mid] {
		l, r := true, true
		start, end = mid, mid
		for i := 0; i <= mid; i++ {
			if mid-i >= 0 {
				if target == nums[mid-i] && l {
					start = mid - i
				} else {
					l = false
				}
			}
			if mid+i < len(nums) {
				if target == nums[mid+i] && r {
					end = mid + i
				} else {
					r = false
				}
			}
			if !l && !r {
				break
			}
		}
	} else {
		for i := mid - 1; i < len(nums) && i >= 0; i++ {
			if nums[i] == target {
				if start == -1 {
					start = i
				}
				end = i
			}
		}
	}
	return []int{start, end}
}
