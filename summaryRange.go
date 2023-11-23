package main

import "fmt"

// 给定一个  无重复元素 的 有序 整数数组 nums 。

// 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

// 列表中的每个区间范围 [a,b] 应该按如下格式输出：

// "a->b" ，如果 a != b
// "a" ，如果 a == b

// 思路
// 遍历nums，如果前一个和后一个差不等于1，并且两个不相等就输出

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}
	result := []string{}
	temp2 := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 && nums[i] != nums[i-1] {
			if len(temp2) == 1 {
				result = append(result, fmt.Sprintf("%d", temp2[0]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", temp2[0], temp2[len(temp2)-1]))
			}
			temp2 = []int{}
		}

		if nums[i] == nums[i-1] {
			continue
		}
		temp2 = append(temp2, nums[i])
	}
	if len(temp2) == 1 {
		result = append(result, fmt.Sprintf("%d", temp2[0]))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", temp2[0], temp2[len(temp2)-1]))
	}
	return result
}
func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(summaryRanges([]int{0, 2, 2, 4, 4, 6, 8, 9}))
	fmt.Println(summaryRanges([]int{0, 0, 0, 0, 0}))
}
