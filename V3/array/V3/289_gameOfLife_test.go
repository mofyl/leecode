/*
 * @Author: lirui
 * @Date: 2022-01-03 14:25:54
 * @LastEditTime: 2022-01-03 16:13:22
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\array\V3\289_gameOfLife_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestGameOfLife(t *testing.T) {

	// board := [][]int{

	// 	{0, 1, 0},
	// 	{0, 0, 1},
	// 	{1, 1, 1},
	// 	{0, 0, 0},
	// }

	board := [][]int{

		{1, 1},
		{1, 0},
	}

	row := len(board)

	gameOfLife(board)

	for i := 0; i < row; i++ {

		fmt.Println(board[i])
	}
}

// 标记为 -1 表示需要改为 1  由于标记完成后 该格子 还是要参与后面格子的计算，所以格子的值 和 标记值需要有联系
//  标记为2 表示需要改为0
func gameOfLife(board [][]int) {

	row := len(board)
	col := len(board[0])

	for i := 0; i < row; i++ {

		for j := 0; j < col; j++ {
			count := liveCount(board, i, j)

			// 规则1, 规则3
			if board[i][j] >= 1 && (count < 2 || count > 3) {
				board[i][j] = 2
			}
			if board[i][j] == 1 && (count == 2 || count == 3) {
				// 规则2 不变
			}

			if board[i][j] <= 0 && count == 3 {
				board[i][j] = -1
			}

		}

	}

	for i := 0; i < row; i++ {

		for j := 0; j < col; j++ {

			if board[i][j] == -1 {
				board[i][j] = 1
			}

			if board[i][j] == 2 {
				board[i][j] = 0
			}
		}
	}

}

func liveCount(nums [][]int, i, j int) int {

	dirs := [][]int{
		{-1, 0},  // 上边
		{1, 0},   // 下边
		{0, -1},  // 左边
		{0, 1},   // 右边
		{-1, -1}, // 左上
		{-1, 1},  // 右上
		{1, -1},  // 左下
		{1, 1},   // 右下
	}
	row := len(nums)
	col := len(nums[0])

	count := 0
	for d := 0; d < len(dirs); d++ {

		x := i + dirs[d][0]
		y := j + dirs[d][1]

		if x < 0 || x >= row || y < 0 || y >= col {
			continue
		}

		if nums[x][y] >= 1 {
			count++
		}

	}

	return count

}
