package main

import (
	"fmt"
	"sort"
)

// 有一些球形气球贴在一堵用 XY 平面表示的墙面上。墙面上的气球记录在整数数组 points ，其中points[i] = [xstart, xend] 表示水平直径在 xstart 和 xend之间的气球。你不知道气球的确切 y 坐标。

// 一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被 引爆 。可以射出的弓箭的数量 没有限制 。 弓箭一旦被射出之后，可以无限地前进。

// 给你一个数组 points ，返回引爆所有气球所必须射出的 最小 弓箭数 。

// 思路
// 先排序
// 然后记录第一个数组的最大值
// 如果这个最大值小于其他数组的起始值，
// 最大值更新为这个数组的最大值
// 箭数量+1

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	} else if len(points) == 1 {
		return 1
	}

	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	maxRight := points[0][1]
	result := 1
	for _, v := range points {
		if v[0] > maxRight {
			maxRight = v[1]
			result++
		}
	}
	return result
}

func main() {
	a := [][]int{
		[]int{10, 16},
		[]int{2, 8},
		[]int{1, 6},
		[]int{7, 12},
	}
	fmt.Println(findMinArrowShots(a))
}
