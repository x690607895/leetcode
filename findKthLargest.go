package main

import "log"

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// 示例 1:

// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
// 示例 2:

// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvsehe/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	log.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[k-1]
}

func quickSort(nums []int, l, r int) {
	i, j := l, r
	temp := nums[l]
	p := l

	for i <= j {
		for j >= p && nums[j] <= temp {
			j--
		}
		if j >= p {
			nums[p] = nums[j]
			p = j
		}
		for i <= p && nums[i] >= temp {
			i++
		}
		if i <= p {
			nums[p] = nums[i]
			p = i
		}
	}
	nums[p] = temp
	if p-l > 1 {
		quickSort(nums, l, p-1)
	}
	if r-p > 1 {
		quickSort(nums, p+1, r)
	}
}
