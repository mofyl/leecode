package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestIntervalIntersection(t *testing.T) {
	nums := [][]int{
		{0, 2},
		{5, 10},
		{13, 23},
		{24, 25},
	}

	nums2 := [][]int{
		{1, 5},
		{8, 12},
		{15, 24},
		{25, 26},
	}

	res := intervalIntersection(nums, nums2)

	fmt.Println(res)
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {

	if len(firstList) < 1 || len(secondList) < 1 {
		return nil
	}

	i, j := 0, 0

	res := make([][]int, 0)
	for j < len(secondList) && i < len(firstList) {

		/*

			先比较第一个位置的数字,
				若second 第一个位置的数字 小于 first 第一个位置的数字
					并且 second 第二个位置的数字 大于 first第一个位置
					说明左边界有交集， 交集的起始位置就是 first 否则左边界没有交集

				若 second 第一个位置的数据 大于 first 第一个位置的数字
					并且 second 第一位置的数据 小于 fist第二位置的数据 就表示有交集
					起始位置就是 second 第一个位置
		*/

		f := tools.Max(firstList[i][0], secondList[j][0])
		c := tools.Min(firstList[i][1], secondList[j][1])

		if f <= c {
			// 表示有交集
			res = append(res, []int{f, c})
		}
		/*
			first : [0 , 2]	[5 , 10]
			second : [1 , 5] [8 , 12]
			交集就是  [1 , 2] [5 , 5] [8 , 10]
		*/
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}

	}

	return res

}
