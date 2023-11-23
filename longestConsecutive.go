package main

import "fmt"

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 思路：
// 转换成map
// 循环map，如果不存在当前数字前一个，就进入循环判断后面的数字是否存在，如果不存在则跳出，存在的话，数字+1，长度+1

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}

	mapNums := map[int]struct{}{}
	for _, v := range nums {
		mapNums[v] = struct{}{}
	}
	result := 0
	min, max := 0, 0
	for k := range mapNums {
		if result != 0 && k > min && k < max {
			continue
		}
		if _, ok := mapNums[k-1]; !ok {
			currNum, currResult := k+1, 1
			for {
				if _, ok := mapNums[currNum]; ok {
					currNum++
					currResult++
				} else {
					break
				}
			}
			if currResult > result {
				min = k
				max = currNum
				result = currResult
			}
		}
	}
	return result
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}))
	fmt.Println(longestConsecutive([]int{0, 0, -1, 1}))
}
