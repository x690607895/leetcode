package main

import "fmt"

// n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

// 你需要按照以下要求，给这些孩子分发糖果：

// 每个孩子至少分配到 1 个糖果。
// 相邻两个孩子评分更高的孩子会获得更多的糖果。
// 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。

// 思路
// 先计算出从左到右的每个孩子需要的糖果
// 再计算出从右到左的每个孩子需要的糖果
// 比较返回每个孩子需要的最小糖果数

func candy(ratings []int) int {
	result := make([]int, 0)
	for k, v := range ratings {
		if k > 0 && v > ratings[k-1] {
			data := result[k-1] + 1
			result = append(result, data)
		} else {
			result = append(result, 1)
		}
	}

	r := 1
	res := 0
	for i := len(ratings) - 1; i >= 0; i-- {
		if i != len(ratings)-1 && ratings[i] > ratings[i+1] {
			r++
		} else {
			r = 1
		}
		res += max(result[i], r)
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
}
