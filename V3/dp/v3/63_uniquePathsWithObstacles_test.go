package v3

import (
	"fmt"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {

	// obstacleGrid := [][]int{
	// 	{0, 0, 0},
	// 	{0, 1, 0},
	// 	{0, 0, 0},
	// }

	obstacleGrid := [][]int{
		{0, 1},
		{0, 0},
	}

	res := uniquePathsWithObstacles(obstacleGrid)

	fmt.Println(res)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	// return uniquePathsWithObstaclesHelper(obstacleGrid, 0, 0)
	return uniquePathsWithObstaclesDP(obstacleGrid)
}

func uniquePathsWithObstaclesHelper(obstacleGrid [][]int, row, col int) int {

	if row == len(obstacleGrid)-1 && col == len(obstacleGrid[0])-1 {
		return 1
	}

	if row >= len(obstacleGrid) || col >= len(obstacleGrid[0]) {
		return 0
	}
	num := 0
	if col+1 < len(obstacleGrid[0]) && obstacleGrid[row][col+1] == 0 {
		// 向右走
		num += uniquePathsWithObstaclesHelper(obstacleGrid, row, col+1)
	}
	// 向下走
	if row+1 < len(obstacleGrid) && obstacleGrid[row+1][col] == 0 {
		num += uniquePathsWithObstaclesHelper(obstacleGrid, row+1, col)
	}

	return num

}

func uniquePathsWithObstaclesDP(obstacleGrid [][]int) int {

	dp := make([][]int, len(obstacleGrid))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}

	if obstacleGrid[len(dp)-1][len(dp[0])-1] == 1 {
		// 目的位置为 障碍物
		return 0
	}
	dp[len(dp)-1][len(dp[0])-1] = 1
	// 这里是倒着走的 也就是 从终点 向起点走
	for i := len(dp) - 1; i >= 0; i-- {

		for j := len(dp[0]) - 1; j >= 0; j-- {

			// 该位置为障碍物
			if obstacleGrid[i][j] == 1 {
				continue
			}

			// 向 && 地图里面 目标位置不是 障碍物
			if i-1 >= 0 && obstacleGrid[i-1][j] == 0 {
				dp[i-1][j] += dp[i][j]
			}

			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				dp[i][j-1] += dp[i][j]
			}

		}

	}

	return dp[0][0]

}
