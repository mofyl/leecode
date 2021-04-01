package V1

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	//matrix := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}

	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	//matrix := [][]int{
	//	{1},
	//	{4},
	//}

	res := spiralOrder(matrix)
	fmt.Println(res)
}

func spiralOrder(matrix [][]int) []int {

	row := len(matrix)
	if row == 0 {
		return nil
	}

	col := len(matrix[0])

	if col == 0 {
		return nil
	}
	res := make([]int, 0, row*col)

	u := 0
	d := row - 1
	l := 0
	r := col - 1
	for {
		// 从左向 右
		for i := l; i <= r; i++ {
			res = append(res, matrix[u][i])
		}
		u++
		if u > d {
			break
		}
		// 从上往下
		for i := u; i <= d; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		if r < l {
			break
		}
		// 从右往左
		for i := r; i >= l; i-- {
			res = append(res, matrix[d][i])
		}
		d--
		if d < u {
			break
		}
		// 从下往上
		for i := d; i >= u; i-- {
			res = append(res, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}

	return res

}
