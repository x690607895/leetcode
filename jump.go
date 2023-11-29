package main

import "fmt"

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

// 思路
// 每次取当前可跳区间里的最大区间去当下次的终点，只要终点超过了数组就意味到达

func jump(nums []int) int {
	end := 0
	maxPosition := 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		tmpPosition := i + nums[i]
		if tmpPosition > maxPosition {
			maxPosition = tmpPosition
		}
		if i == end {
			end = maxPosition
			step++
			if end > len(nums) {
				break
			}
		}
	}
	return step
}

func main() {
	fmt.Println(jump([]int{2, 12, 1, 1, 4}))
}
