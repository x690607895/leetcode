package main

import "fmt"

/*
这题可以分为两个步骤
先计算左边每个值的乘机
再计算右边每个值的乘机
最后把结果相乘的到结果
*/
func productExceptSelf(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	lenN := len(nums)
	result := []int{1}
	for i := 1; i < lenN; i++ {
		tmp := result[i-1] * nums[i-1]
		result = append(result, tmp)
	}
	tmp := nums[lenN-1]
	for i := lenN - 2; i >= 0; i-- {
		result[i] *= tmp
		tmp *= nums[i]
	}
	return result
}
func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
