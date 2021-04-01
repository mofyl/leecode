package v1

import (
	"fmt"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	//
	//ob := [][]int{
	//	{0, 0, 0},
	//	{0, 1, 0},
	//	{0, 0, 0},
	//}

	ob := [][]int{
		{0, 0},
		{1, 1},
		{0, 0},
		//{0, 0},
	}
	res := uniquePathsWithObstacles(ob)

	fmt.Println(res)

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	n := len(obstacleGrid)
	if n < 1 {
		return 0
	}
	m := len(obstacleGrid[0])
	if len(obstacleGrid[0]) < 1 {
		return 0
	}

	dp := make([][]int, 0, n)
	// 初始化dp
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		if i == 0 {
			for j := 0; j < len(tmp); j++ {
				if obstacleGrid[0][j] != 1 {
					// 这里表示没有障碍
					if j > 0 {
						// 没有障碍 需要根据 前面的格子 来做出判断
						// 若前面的格子有障碍物。那么是到不了这个格子的
						tmp[j] = tmp[j-1]
					} else {
						tmp[j] = 1
					}
				} else {
					// 这里表示有障碍
					tmp[j] = 0
				}
			}
		} else {
			if obstacleGrid[i][0] != 1 {
				// 这里表示 没有障碍 。但是需要根据前面的格子做出判断
				// 若前面的格子有障碍物。那么是到不了这个格子的
				tmp[0] = dp[i-1][0]
			} else {
				tmp[0] = 0
			}
		}

		dp = append(dp, tmp)
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}

	return dp[n-1][m-1]
}
