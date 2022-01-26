package v3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMinPathSum(t *testing.T) {

	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	// grid := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// }

	res := minPathSum(grid)

	fmt.Println(res)
}

func minPathSum(grid [][]int) int {
	// return minPathSumHelper(grid, 0, 0)
	return minPathSumDP(grid)
}

func minPathSumHelper(grid [][]int, row, col int) int {

	if row >= len(grid) || col >= len(grid[0]) {
		return -1
	}

	right := minPathSumHelper(grid, row, col+1)
	down := minPathSumHelper(grid, row+1, col)

	if right == -1 && down == -1 {
		return grid[row][col]
	}

	if right == -1 {
		return down + grid[row][col]
	}

	if down == -1 {
		return right + grid[row][col]
	}

	return tools.Min(right, down) + grid[row][col]
}

func minPathSumDP(grid [][]int) int {

	row := len(grid)
	col := len(grid[0])

	dp := make([][]int, row)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}

	dp[0][0] = grid[0][0]

	for i := 0; i < row; i++ {

		for j := 0; j < col; j++ {

			if i-1 < 0 && j-1 < 0 {
				continue
			}

			if j-1 < 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i-1 < 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else {
				dp[i][j] = tools.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}

		}

	}

	return dp[row-1][col-1]

}
