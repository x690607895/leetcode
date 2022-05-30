package main

import (
	"log"
	"sort"
)

// 给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

// 示例 1：

// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2,2]
// 示例 2:

// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[4,9]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2y0c2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	log.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	result := make([]int, 0)
	nums1Index, nums2Index := 0, 0
	for {
		if nums1Index >= len(nums1) || nums2Index >= len(nums2) {
			break
		}
		if nums1[nums1Index] == nums2[nums2Index] {
			result = append(result, nums1[nums1Index])
			nums1Index++
			nums2Index++
		} else if nums1[nums1Index] < nums2[nums2Index] {
			nums1Index++
		} else {
			nums2Index++
		}

	}
	return result
}
