package V1

import (
	"fmt"
	"leecode/tools"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	//intervals := [][]int{
	//	{1, 3},
	//	{2, 6},
	//	{8, 10},
	//	{15, 18},
	//}
	//
	intervals := [][]int{
		{1, 4},
		{0, 0},
	}
	//
	//intervals := [][]int{
	//	{1, 4},
	//	{0, 2},
	//	{3, 5},
	//}

	//intervals := [][]int{
	//	{2, 3},
	//	{4, 5},
	//	{6, 7},
	//	{8, 9},
	//	{1, 10},
	//}

	res := merge(intervals)
	fmt.Println(res)
}

func merge(intervals [][]int) [][]int {
	// 先排序 在判断
	// 这里传入的 func 用于比较的
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] > intervals[j][0]
	})
	//
	res := make([][]int, 0, len(intervals))
	//
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] >= cur[0] {
			prev[1] = tools.Max(prev[1], cur[1])
			prev[0] = tools.Min(prev[0], cur[0])
		} else {
			res = append(res, prev)
			prev = cur
		}
	}

	res = append(res, prev)

	return res
}
