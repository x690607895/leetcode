package main

import "fmt"

func canJump(nums []int) bool {
	position := 0
	for k, v := range nums {
		if k > position {
			return false
		}
		tmpPosition := k + v
		if tmpPosition > position {
			position = tmpPosition
		}
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{2, 5, 0, 0}))
}
