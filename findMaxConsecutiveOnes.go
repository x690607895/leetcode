package main

import "log"

// 给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。

// 示例 1：

// 输入：nums = [1,1,0,1,1,1]
// 输出：3
// 解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
// 示例 2:

// 输入：nums = [1,0,1,1,0,1]
// 输出：2

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/cd71t/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	log.Println(findMaxConsecutiveOnes(nums))
}
func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) <= 1 {
		if nums[0] == 0 {
			return 0
		}
		return 1
	}
	prev := nums[0]
	nums2 := 0
	maxNums := 0
	for i := 0; i < len(nums); i++ {
		if prev != nums[i] {
			nums2 = 0
		}
		if nums[i] == 1 {
			nums2++
		}
		prev = nums[i]
		if nums2 > maxNums {
			maxNums = nums2
		}
	}
	return maxNums
}
