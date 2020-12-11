package array

import (
	"fmt"
	"strings"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	n := 8

	res := solveNQueens(n)

	fmt.Println(len(res))

	for i := 0; i < len(res); i++ {

		fmt.Println(res[i])

		fmt.Println("================")

	}

	//fmt.Println(res)
}

func solveNQueens(n int) [][]string {

	res := make([][]string, 0, n)

	solveNQueensLoop(0, n, &res, []string{})

	return res
}

func solveQueenPos(col, row int, res []string, n int) bool {

	upCol := col
	upRow := row
	// 判断上面是否有Q

	for upRow-1 >= 0 {
		if res[upRow-1][upCol] == 'Q' {
			return false
		}
		upRow--
	}

	// 判断左斜有没有Q
	lCol := col
	lRow := row
	for lCol-1 >= 0 && lRow-1 >= 0 {
		if res[lRow-1][lCol-1] == 'Q' {
			return false
		}
		lCol--
		lRow--
	}

	// 判断右斜 有没有Q
	rCol := col
	rRow := row
	for rCol+1 < n && rRow-1 >= 0 {
		if res[rRow-1][rCol+1] == 'Q' {
			return false
		}
		rCol++
		rRow--
	}

	return true
}

func solveNQueensLoop(start int, n int, res *[][]string, cur []string) {

	if start == n {
		// 出现了 结果 直接走
		tmp := make([]string, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < n; i++ {

		// 判断当前位置是否正确
		if !solveQueenPos(i, start, cur, n) {
			continue
		}
		curStr := strings.Builder{}

		for j := 0; j < n; j++ {
			if j == i {
				curStr.WriteString("Q")
			} else {
				curStr.WriteString(".")
			}
		}
		cur = append(cur, curStr.String())
		solveNQueensLoop(start+1, n, res, cur)
		cur = cur[:len(cur)-1]
	}

}
