package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {

	board := [][]byte{

		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word := "ABCCED"
	//
	//board := [][]byte{
	//	{'a', 'b'},
	//	{'c', 'd'},
	//}
	//
	//word := "abcd"

	res := exist(board, word)

	fmt.Println(res)
}

func exist(board [][]byte, word string) bool {

	for i := 0; i < len(board); i++ {

		for j := 0; j < len(board[i]); j++ {

			if board[i][j] == word[0] {
				// 开始 深搜
				tmp := board[i][j]
				board[i][j] = '1'
				if existHelper(board, word, i, j, 1) {
					return true
				}
				board[i][j] = tmp
			}

		}

	}
	return false

}

func existHelper(board [][]byte, word string, row, col int, idx int) bool {

	if idx >= len(word) {
		return true
	}

	dir := [][]int{
		{0, -1}, // 左边
		{0, 1},  // 右边
		{-1, 0}, // 上边
		{1, 0},  // 下边
	}

	for i := 0; i < len(dir); i++ {

		x := row + dir[i][0]
		y := col + dir[i][1]
		if x >= 0 && x < len(board) &&
			y >= 0 && y < len(board[0]) {

			if board[x][y] != '1' && board[x][y] == word[idx] {
				tmp := board[x][y]
				board[x][y] = '1'
				if existHelper(board, word, x, y, idx+1) {
					return true
				}
				board[x][y] = tmp
			}

		}

	}

	return false
}
