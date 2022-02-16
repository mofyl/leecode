package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	res := spiralOrder(matrix)
	//
	//matrix := [][]int{
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{9, 10, 11, 12},
	//}
	//
	//res := spiralOrder(matrix)

	fmt.Println(res)
}

func spiralOrder(matrix [][]int) []int {
	up := 0
	down := len(matrix) - 1

	if len(matrix) == 0 {
		return []int{}
	}
	left := 0
	right := len(matrix[0]) - 1

	if len(matrix[0]) == 0 {
		return []int{}
	}

	res := make([]int, 0)

	for i := 0; i <= down; i++ {

		// 从 左 向右 打印
		for j := left; j <= right; j++ {
			res = append(res, matrix[i][j])
		}
		up++
		if up > down {
			break
		}
		// 从上到下
		for j := up; j <= down; j++ {
			res = append(res, matrix[j][right])
		}
		right--
		if left > right {
			break
		}
		// 从 右向左
		for j := right; j >= left; j-- {
			res = append(res, matrix[down][j])
		}
		down--
		if up > down {
			break
		}
		// 从下向上
		for j := down; j >= up; j-- {
			res = append(res, matrix[j][left])
		}
		left++
		if left > right {
			break
		}
	}

	return res

}
