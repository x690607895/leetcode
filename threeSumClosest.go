package main

import (
	"fmt"
	"math"
	"sort"
)

// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

// 返回这三个数的和。

// 假定每组输入只存在恰好一个解。

func threeSumClosest(nums []int, target int) int {
	lNum := len(nums)
	if lNum == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	res := math.MaxInt

	for i := 0; i < lNum; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// 如果当前数字和前面数字想等则跳过
			continue
		}
		// j是i后面那个，k是从后往前
		j, k := i+1, lNum-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			// 如果总和>目标，则需要移动K
			// 否则移动J
			if sum > target {
				temp := nums[k]
				for j < k && nums[k] == temp {
					k--
				}
			} else {
				temp := nums[j]
				for j < k && nums[j] == temp {
					j++
				}
			}
		}
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, 4}, 2))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}
