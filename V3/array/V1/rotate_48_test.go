package V1

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	//
	//matrix := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}

	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {

	row := len(matrix)
	if row == 0 {
		return
	}
	col := len(matrix[0])
	if col == 0 {
		return
	}
	var tmp int
	for i := 0; i < row/2; i++ {

		for j := i; j < col-1-i; j++ {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[row-1-j][i]
			matrix[row-1-j][i] = matrix[row-1-i][row-1-j]
			matrix[row-1-i][row-1-j] = matrix[j][col-1-i]
			matrix[j][col-1-i] = tmp

			//tmp1 := matrix[j][col-1-i]
			//matrix[j][col-1-i] = matrix[i][j]
			//
			//tmp2 := matrix[row-1-i][col-1-j]
			//matrix[row-1-i][col-1-j] = tmp1
			//
			//tmp3 := matrix[row-1-j][i]
			//matrix[row-1-j][i] = tmp2
			//
			//matrix[i][j] = tmp3

		}

	}

}
