/*
 * @Author: lirui
 * @Date: 2021-12-30 12:47:10
 * @LastEditTime: 2021-12-30 14:53:04
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\array\V3\54_spiralOrder_test.go
 */
package V3

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

	// matrix := [][]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// }

	res := spiralOrder(matrix)

	fmt.Println(res)
}

func spiralOrder(matrix [][]int) []int {

	lLimit := 0
	rLimit := len(matrix[0]) - 1
	upLimit := 0
	dnLimit := len(matrix) - 1

	res := make([]int, 0)
	for {

		// 从 左向右
		for i := lLimit; i <= rLimit; i++ {
			res = append(res, matrix[upLimit][i])
		}
		upLimit++
		if upLimit > dnLimit {
			break
		}

		// 从上向下
		for i := upLimit; i <= dnLimit; i++ {
			res = append(res, matrix[i][rLimit])
		}
		rLimit--
		if rLimit < lLimit {
			break
		}

		for i := rLimit; i >= lLimit; i-- {
			res = append(res, matrix[dnLimit][i])
		}
		dnLimit--
		if dnLimit < upLimit {
			break
		}

		for i := dnLimit; i >= upLimit; i-- {
			res = append(res, matrix[i][lLimit])
		}
		lLimit++
		if lLimit > rLimit {
			break
		}
	}

	return res
}
