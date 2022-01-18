/*
 * @Author: lirui
 * @Date: 2022-01-09 14:09:59
 * @LastEditTime: 2022-01-09 14:49:00
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\backtracking\v3\51_solveNQueens_test.go
 */

package v3

import (
	"fmt"
	"strings"
	"testing"
)

func TestSolveNQueens(t *testing.T) {

	n := 1
	res := solveNQueens(n)

	fmt.Println(res)
}

func solveNQueens(n int) [][]string {

	res := make([][]string, 0)

	solveNQueensHelper(n, 0, []string{}, &res)

	return res
}

func solveNQueensHelper(n int, row int, cur []string, res *[][]string) {

	if row == n {
		tmp := make([]string, len(cur))
		copy(tmp, cur)

		*res = append(*res, tmp)
		return
	}

	for i := 0; i < n; i++ {

		if !solveNQueensCheck(cur, row, i) {
			continue
		}

		sb := strings.Builder{}

		for j := 0; j < n; j++ {
			if j == i {
				sb.WriteByte('Q')
			} else {
				sb.WriteByte('.')
			}
		}

		cur = append(cur, sb.String())
		solveNQueensHelper(n, row+1, cur, res)
		cur = cur[:len(cur)-1]
	}

}

// 这里虽然给的是  row ,但是却 取不到 row
func solveNQueensCheck(board []string, row, col int) bool {

	if len(board) < 1 {
		return true
	}

	// 检查上面
	for i := 0; i < len(board); i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 检查左斜
	i := row
	j := col
	for i > 0 && j > 0 {

		if board[i-1][j-1] == 'Q' {
			return false
		}
		i--
		j--
	}

	// 检查右斜
	i = row
	j = col

	for i > 0 && j < len(board[0])-1 {

		if board[i-1][j+1] == 'Q' {
			return false
		}

		i--
		j++
	}

	return true
}
