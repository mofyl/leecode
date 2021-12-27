/*
 * @Author: lirui38
 * @Date: 2021-12-26 11:22:23
 * @LastEditTime: 2021-12-26 18:11:42
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\bfs\v1\130_solve_test.go
 */

package v1

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {

	// board := [][]byte{

	// 	{'X', 'X', 'X', 'X'},
	// 	{'X', 'O', 'O', 'X'},
	// 	{'X', 'X', 'O', 'X'},
	// 	{'X', 'O', 'X', 'X'},
	// }

	// solve(board)

	board := [][]byte{
		{'X', 'O', 'X'},
		{'O', 'X', 'O'},
		{'X', 'O', 'X'},
	}

	solve(board)

	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}

}

var (
	solveCol = 0
	solveRow = 0
)

func solve(board [][]byte) {

	solveRow = len(board)

	// 小于 3 行 是没办法用 'X' 来包围 'O' 的
	if solveRow < 3 {
		return
	}

	solveCol = len(board[0])
	// 先将 边界上的 'O' 和  边界上 'O' 直接相连的 'O' 标记为 '1'

	// 修改头尾两列
	for i := 0; i < solveRow; i++ {
		solveHelper(board, i, 0)
		solveHelper(board, i, solveCol-1)
	}
	// 修改 除 头尾列的 中间列
	for i := 1; i < solveCol; i++ {
		solveHelper(board, 0, i)
		solveHelper(board, solveRow-1, i)
	}

	// 将 所有 被 标记为 '1' 的还原为 'O'
	// 将所有 未被标记为 '1' 的 'O' 改为 'X'

	for i := 0; i < solveRow; i++ {
		for j := 0; j < solveCol; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}

}

func solveHelper(board [][]byte, i, j int) {
	if i < 0 || i >= solveRow ||
		j < 0 || j >= solveCol {
		return
	}

	if board[i][j] == '1' || board[i][j] == 'X' {
		return
	}

	board[i][j] = '1'

	solveHelper(board, i, j-1)
	solveHelper(board, i, j+1)
	solveHelper(board, i-1, j)
	solveHelper(board, i+1, j)

}
