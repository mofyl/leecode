package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestMovingCount(t *testing.T) {

	//m := 2
	//n := 3
	//k := 1

	m := 3
	n := 2
	k := 17

	res := movingCount(m, n, k)

	fmt.Println(res)
}

// 该题是要求 机器人 一共可以走多少个不重复的格子！！！
func movingCount(m int, n int, k int) int {

	if m < 1 || n < 1 {
		return 0
	}

	helper := make([][]bool, m)

	for i := 0; i < len(helper); i++ {
		helper[i] = make([]bool, n)
	}

	helper[0][0] = true
	// (0,0) 算一个格子
	return movingCountHelper(m, n, k, helper, 0, 0) + 1
}

func movingCountHelper(m int, n int, k int, helper [][]bool, row, col int) int {

	rowSum := movingCountGetSum(row)
	colSum := movingCountGetSum(col)

	if rowSum+colSum > k {
		return -1
	}

	if row < 0 || row >= m ||
		col < 0 || col >= n {
		return -1
	}

	dir := [][]int{
		{0, -1}, // 向左
		{0, 1},  // 向右
		{-1, 0}, // 向上
		{1, 0},  // 向下
	}

	cur := 0

	for i := 0; i < len(dir); i++ {

		x := row + dir[i][0]
		y := col + dir[i][1]
		//
		if x >= 0 && x < m &&
			y >= 0 && y < n &&
			!helper[x][y] {
			helper[x][y] = true
			cur += movingCountHelper(m, n, k, helper, x, y) + 1
			// 既然是不重复的格子 那么就没有必要回溯了
			//helper[x][y] = false
		}
	}
	return cur
}

func movingCountGetSum(m int) int {

	sum := 0
	for m > 0 {

		sum += m % 10
		m /= 10
	}

	return sum

}
