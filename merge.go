package main

import (
	"log"
	"math"
	"sort"
)

// 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

//

// 示例 1：

// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 示例 2：

// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/array-and-string/c5tv3/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	log.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start, end := intervals[0][0], intervals[0][1]
	result := make([][]int, 0)
	arrLen := len(intervals)
	for i := 1; i < arrLen; i++ {
		if intervals[i][0] >= start && intervals[i][0] <= end {
			end = int(math.Max(float64(end), float64(intervals[i][1])))
		} else {
			result = append(result, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	result = append(result, []int{start, end})
	return result
}
