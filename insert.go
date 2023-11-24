package main

import "fmt"

// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

// 思路
// 循环遍历数组
// 如果比区间小则直接插入返回数组
// 不然记录区间的开始和结束（区间开始：当前数组开头和区间开头最小值，区间结束：当前数组结束和区间开头的最大值）
// 开始记录区间后
// 如果遇到比当前数组最大值小的直接跳过该数组
// 不然记录区间最大值
// 两种情况
// 区间最大值<数组最小值
// 区间数组和当前数组一起写入
// 区间最大值>=数组最小值
// 区间数组最大值=当前数组最大值
// 区间数组写入
// 区间数组写入完毕，后面数组依次写入

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	start, end := newInterval[0], newInterval[1]
	isStart, isMerge := false, false
	result := [][]int{}
	merge := []int{start, end}

	for _, v := range intervals {
		if !isStart {
			// 还没开始
			if v[1] < start {
				// start比当前数组最大值小的直接塞入
				result = append(result, v)
			} else {
				isStart = true
				if v[0] > end {
					// 如果区间的最大值比newInterval[1]都小，那就证明整个区间都比当前区间小
					// 直接查到了最前面
					result = append(result, newInterval)
					result = append(result, v)
					isMerge = true
				}

				if merge[0] > v[0] {
					merge[0] = v[0]
				}
				// 比较end和当前数组最大值的关系，选择最大的那个
				if end < v[1] {
					end = v[1]
				}
				merge[1] = end

			}
		} else {
			// 已经开始了
			if !isMerge {
				// 还在merge中
				if end >= v[1] {
					// end比区间最大值大，直接跳过这个区间
					continue
				} else {
					// end比区间最小值小，merge写入，并且把当前数组写入
					// 否则，merge的最大值同步当前数组最大值，写入
					if end < v[0] {
						result = append(result, merge)
						result = append(result, v)
					} else {
						end = v[1]
						merge[1] = end
						result = append(result, merge)
					}
					isMerge = true
				}
			} else {
				// start和merge都做完了，直接插入后续数组
				result = append(result, v)
			}
		}
	}
	if !isMerge {
		result = append(result, merge)
	}
	return result
}

func main() {
	a := [][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
	}
	fmt.Println(insert(a, []int{4, 8}))
	fmt.Println(insert([][]int{[]int{1, 5}}, []int{0, 3}))
	fmt.Println(insert([][]int{[]int{1, 5}}, []int{0, 0}))
	fmt.Println(insert([][]int{[]int{1, 3}, []int{6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{}, []int{2, 3}))
}
