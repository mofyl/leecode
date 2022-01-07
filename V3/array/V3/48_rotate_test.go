/*
 * @Author: lirui
 * @Date: 2021-12-30 14:55:28
 * @LastEditTime: 2021-12-30 17:26:50
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\array\V3\48_rotate_test.go
 */
package V3

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// matrix := [][]int{
	// 	{5, 1, 9, 11},
	// 	{2, 4, 8, 10},
	// 	{13, 3, 6, 7},
	// 	{15, 14, 12, 16},
	// }

	// matrix := [][]int{
	// 	{1},
	// }

	// matrix := [][]int{
	// 	{1, 2},
	// 	{3, 4},
	// }

	rotate(matrix)

	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}

func rotate(matrix [][]int) {

	l := 0
	r := len(matrix[0]) - 1
	u := 0
	d := len(matrix) - 1

	for u < d && l < r {

		// 最外圈进行旋转 90度
		for i := 0; i < r-l; i++ {

			// 这里可以倒着存放

			// (u , l+i)
			tmp := matrix[u][l+i]
			// (d-i , l) -> (u , l+i)
			matrix[u][l+i] = matrix[d-i][l]
			// (d , r-i) -> (d-i , l)
			matrix[d-i][l] = matrix[d][r-i]
			// (u+i ,r) -> (d , r-i)
			matrix[d][r-i] = matrix[u+i][r]
			// (u , l+i) -> (u+i , r)
			matrix[u+i][r] = tmp

			// tmp := matrix[u+i][r]
			// // (u , l+i) -> (u+i , r)
			// matrix[u+i][r] = matrix[u][l+i]

			// // (u+i,r) -> (d , r-i)
			// tmp1 := matrix[d][r-i]
			// matrix[d][r-i] = tmp
			// tmp = tmp1

			// // (d, r-i) -> (d-i,l)
			// tmp1 = matrix[d-i][l]
			// matrix[d-i][l] = tmp
			// tmp = tmp1

			// // (d-i ,l) -> (u , l+i)
			// matrix[u][l+i] = tmp
		}

		u++
		l++
		r--
		d--

	}

}
