package main

import "fmt"

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

// 返回容器可以储存的最大水量。

// 思路
// 从两边开始计算最大值
// 算完后短的那边往后（前）移
// 算之前入过比前（后）短就直接过

func maxArea(height []int) int {
	if len(height) == 2 {
		return min(height[0], height[1])
	}
	maxResult := 0
	for l, r := 0, len(height)-1; l < r; {
		if l > 0 && height[l] < height[l-1] {
			l++
			continue
		}
		if r < len(height)-1 && height[r] < height[r+1] {
			r--
			continue
		}
		maxResult = max(maxResult, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxResult
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 2, 3}))
}
