package v3

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {

	// m := 3
	// n := 7

	m := 3
	n := 2

	res := uniquePaths(m, n)

	fmt.Println(res)
}

func uniquePaths(m int, n int) int {

	// return uniquePathsHelper(m, n, 0, 0)
	return uniquePathsDP(m, n)
}

func uniquePathsHelper(m, n int, row, col int) int {

	if row == m-1 && col == n-1 {
		return 1
	}

	if row > m || col > n {
		return 0
	}

	return uniquePathsHelper(m, n, row+1, col) + uniquePathsHelper(m, n, row, col+1)
}

func uniquePathsDP(m, n int) int {

	dp := make([][]int, m)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	dp[m-1][n-1] = 1

	for i := m - 1; i >= 0; i-- {

		for j := n - 1; j >= 0; j-- {

			if i-1 >= 0 {
				dp[i-1][j] += dp[i][j]
			}

			if j-1 >= 0 {
				dp[i][j-1] += dp[i][j]
			}

		}

	}

	return dp[0][0]

}
