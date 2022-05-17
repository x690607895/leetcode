package main

import "log"

// 寻找数组的中心索引

// 给你一个整数数组 nums ，请计算数组的 中心下标 。

// 数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。

// 如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。

// 如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/cxqdh/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	a := []int{1, 7, 3, 6, 5, 6}
	a = []int{-1, -1, -1, -1, -1, -1}
	log.Println(pivotIndex(a))

}
func pivotIndex(nums []int) int {
	if len(nums) <= 1 {
		if len(nums) == 1 && nums[0] == 0 {
			return 0
		}
		return -1
	}

	total := 0
	for _, v := range nums {
		total += v
	}
	l := 0
	for k, v := range nums {
		if 2*l+v == total {
			return k
		}
		l += v
	}
	return -1
}
