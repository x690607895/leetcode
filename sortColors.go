package main

import "github.com/go-acme/lego/log"

// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

// 必须在不使用库的sort函数的情况下解决这个问题。

// 示例 1：

// 输入：nums = [2,0,2,1,1,0]
// 输出：[0,0,1,1,2,2]
// 示例 2：

// 输入：nums = [2,0,1]
// 输出：[0,1,2]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvg25c/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	nums := []int{1, 2, 0}
	sortColors(nums)
	log.Println(nums)
}
func sortColors(nums []int) {
	p1, p2 := 0, len(nums)-1
	for k := range nums {
		for ; k <= p2 && nums[k] == 2; p2-- {
			nums[k], nums[p2] = nums[p2], nums[k]
		}
		if nums[k] == 0 {
			nums[k], nums[p1] = nums[p1], nums[k]
			p1++
		}
	}
}
