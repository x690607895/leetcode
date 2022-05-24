package main

import "log"

func main() {
	nums := []int{2, 4, 6, 8, 10, 1}
	log.Println(findMin(nums))
}

func findMin(nums []int) int {
	first, last := 0, len(nums)-1
	for first < last {
		mid := first + (last-first)/2
		if nums[mid] < nums[last] {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return nums[first]
}
