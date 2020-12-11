package dp

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	// 这里m是列
	m := 23
	// n 是行
	n := 12
	res := uniquePaths(m, n)
	fmt.Println(res)
}

func uniquePaths(m, n int) int {

	dp := make([][]int, 0, m)
	// 初始化dp
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		if i == 0 {
			for k := 0; k < len(tmp); k++ {
				tmp[k] = 1
			}
		} else {
			tmp[0] = 1
		}

		dp = append(dp, tmp)
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			// 当前格子的步数 需要依赖 上方 和 左方的步数
			// 因为 上方和左方都可以到当前格子
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[n-1][m-1]
}

func uniquePaths_v1(m int, n int) int {

	var res int
	uniquePathsLoop(1, 1, m, n, &res)
	return res
}

func uniquePathsLoop(i, j, m, n int, res *int) {

	if i == m && j == n {
		*res++
	}

	if i > m {
		return
	}
	if j > n {
		return
	}
	uniquePathsLoop(i+1, j, m, n, res)
	uniquePathsLoop(i, j+1, m, n, res)
}

func uniquePathsLoop_(i, j, m, n int, res *int, cur string, resSlice *[]string) {

	if i == m && j == n {
		*res++
		*resSlice = append(*resSlice, cur+",")
	}

	if i > m {
		return
	}
	if j > m {
		return
	}
	cur += fmt.Sprintf("+%d ", i)
	uniquePathsLoop_(i+1, j, m, n, res, cur, resSlice)
	cur = cur[:len(cur)-2]

	cur += fmt.Sprintf("-%d ", j)
	uniquePathsLoop_(i, j+1, m, n, res, cur, resSlice)
	cur = cur[:len(cur)-2]
}
