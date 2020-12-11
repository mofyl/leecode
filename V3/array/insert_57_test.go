package array

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestInsert(t *testing.T) {
	//interval := [][]int{
	//	{1, 3},
	//	{6, 9},
	//}
	//newInterval := []int{2, 5}
	//
	//interval := [][]int{
	//	{1, 2},
	//	{3, 5},
	//	{6, 7},
	//	{8, 10},
	//	{12, 16},
	//}
	//newInterval := []int{4, 8}

	interval := [][]int{
		{1, 5},
	}
	newInterval := []int{2, 3}

	res := insert(interval, newInterval)
	fmt.Println(res)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	// 这里由于 原来是无重复 无重叠的 所以可以直接进行 遍历
	res := make([][]int, 0, len(intervals))

	i := 0
	// 这里是不重复的部分
	for i < len(intervals) && newInterval[0] > intervals[i][1] {
		res = append(res, intervals[i])
		i++
	}
	// 这里是重复的部分
	/*
		这里的部分表示要合并，合并是往 newInterval上合，合完了最后在加入到res中
		这里由于i要移动，所以不能往 interval[i]上合
	*/
	for i < len(intervals) && newInterval[1] >= intervals[i][0] {
		newInterval[0] = tools.Min(newInterval[0], intervals[i][0])
		newInterval[1] = tools.Max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)

	// 这里是大于的部分
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res

}
