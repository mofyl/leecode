package V1

import (
	"fmt"
	"testing"
)

func TestSolveSudoku(t *testing.T) {

	nums := [][]byte{
		[]byte{
			byte('5'), byte('3'), byte('.'), byte('.'), byte('7'), byte('.'), byte('.'), byte('.'), byte('.'),
		},
		[]byte{
			byte('6'), byte('.'), byte('.'), byte('1'), byte('9'), byte('5'), byte('.'), byte('.'), byte('.'),
		},
		[]byte{
			byte('.'), byte('9'), byte('8'), byte('.'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.'),
		},
		[]byte{
			byte('8'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('.'),
		},
		[]byte{
			byte('4'), byte('.'), byte('.'), byte('8'), byte('.'), byte('3'), byte('.'), byte('.'), byte('1'),
		},
		[]byte{
			byte('7'), byte('.'), byte('.'), byte('.'), byte('2'), byte('.'), byte('.'), byte('.'), byte('6'),
		},
		[]byte{
			byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('.'), byte('2'), byte('8'), byte('.'),
		},
		[]byte{
			byte('.'), byte('.'), byte('.'), byte('4'), byte('1'), byte('9'), byte('.'), byte('.'), byte('5'),
		},
		[]byte{
			byte('.'), byte('.'), byte('.'), byte('.'), byte('8'), byte('.'), byte('.'), byte('7'), byte('9'),
		},
	}

	solveSudoku(nums)

	for i := 0; i < len(nums); i++ {
		fmt.Println(string(nums[i]))
	}

}

func solveSudoku(board [][]byte) {

	// 保存的所有列 的数字是否使用过
	// 第一维是 第几列  第二维是 当前列的 9个数字的使用情况
	// 这里是 9*10 由于我们填写的时候是从 1到9的 所以这里给了10个位置
	colUsed := [9][10]bool{}
	// 保存 的所有行 的数字是否使用过
	// 第一维是 第几行  第二维是 当前行的 9个数字的使用情况
	// 这里是 9*10   由于我们填写的时候是从 1到9的 所以这里给了10个位置
	rowUsed := [9][10]bool{}
	// 保存 所有格子 的数字 是否使用过
	// box 定位方式：row/3 col/3 就定位到是属于第几个box
	// 第几个box 中的9个元素的使用情况 是 3*3*10 的  给10个的原因同上
	boxUsed := [3][3][10]bool{}

	// 这里将 现在有数字的地方 该成true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			b := board[i][j]
			if b != '.' {
				n := b - '0'
				colUsed[j][n] = true
				rowUsed[i][n] = true
				boxUsed[i/3][j/3][n] = true
			}
		}
	}

	solveSudokuLoop(board, colUsed, rowUsed, boxUsed, 0, 0)
}

func solveSudokuLoop(nums [][]byte, colUsed [9][10]bool, rowUsed [9][10]bool, boxUsed [3][3][10]bool,
	row int, col int) bool {

	if col == len(nums) {
		row++
		col = 0
		if row == len(nums) {
			return true
		}
	}

	for i := 1; i <= 9; i++ {

		if string(nums[row][col]) == "." {
			// 开始剪枝
			if !colUsed[col][i] && !rowUsed[row][i] && !boxUsed[row/3][col/3][i] {

				colUsed[col][i] = true
				rowUsed[row][i] = true
				boxUsed[row/3][col/3][i] = true
				nums[row][col] = byte(i + '0')
				if solveSudokuLoop(nums, colUsed, rowUsed, boxUsed, row, col+1) {
					return true
				}

				colUsed[col][i] = false
				rowUsed[row][i] = false
				boxUsed[row/3][col/3][i] = false
				nums[row][col] = byte('.')
			}
		} else {
			// 当前行已经被填写
			return solveSudokuLoop(nums, colUsed, rowUsed, boxUsed, row, col+1)
		}

	}

	return false

}
