package dp

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMinPathSum(t *testing.T) {

	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	res := minPathSum(grid)
	fmt.Println(res)
}

func minPathSum(grid [][]int) int {

	n := len(grid)
	if n < 1 {
		return 0
	}

	m := len(grid[0])
	if m < 1 {
		return 0
	}

	dp := make([][]int, 0, n)
	// 初始化dp
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		dp = append(dp, tmp)
	}

	//init first col
	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i][0] = grid[i][0] + dp[i-1][0]
		} else {
			dp[i][0] = grid[i][0]
		}
	}

	// init first row
	for i := 0; i < m; i++ {
		if i > 0 {
			dp[0][i] = grid[0][i] + dp[0][i-1]
		} else {
			dp[0][i] = grid[0][i]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = tools.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[n-1][m-1]

}
