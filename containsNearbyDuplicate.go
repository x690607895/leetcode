package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	mapNums := map[int]int{}
	for k2, v := range nums {
		if _, ok := mapNums[v]; ok && k2-mapNums[v] <= k {
			return true
		}
		mapNums[v] = k2
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
