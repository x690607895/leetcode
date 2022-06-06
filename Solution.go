package main

import (
	"log"
	"math/rand"
)

// 给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。

// 实现 Solution class:

// Solution(int[] nums) 使用整数数组 nums 初始化对象
// int[] reset() 重设数组到它的初始状态并返回
// int[] shuffle() 返回数组随机打乱后的结果

// 示例 1：

// 输入
// ["Solution", "shuffle", "reset", "shuffle"]
// [[[1, 2, 3]], [], [], []]
// 输出
// [null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]

// 解释
// Solution solution = new Solution([1, 2, 3]);
// solution.shuffle();    // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。例如，返回 [3, 1, 2]
// solution.reset();      // 重设数组到它的初始状态 [1, 2, 3] 。返回 [1, 2, 3]
// solution.shuffle();    // 随机返回数组 [1, 2, 3] 打乱后的结果。例如，返回 [1, 3, 2]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn6gq1/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	a := []int{1, 2, 3, 4, 6}
	b := Constructor(a, []int{2, 3, 4})
	log.Println(b)
	log.Println(b.Shuffle())
	log.Println(b.Reset())
}

type Solution struct {
	src   []int
	lenth int
}

func Constructor(nums ...[]int) Solution {
	log.Println(nums)
	return Solution{nums[0], len(nums)}
}

func (this *Solution) Reset() []int {
	return this.src
}

func (this *Solution) Shuffle() []int {
	result, temp := make([]int, this.lenth), make([]int, this.lenth)
	copy(temp, this.src)
	num := this.lenth
	for v := range result {
		index := rand.Intn(num)
		result[v] = temp[index]
		temp = append(temp[:index], temp[index+1:]...)
		num--
	}
	return result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
