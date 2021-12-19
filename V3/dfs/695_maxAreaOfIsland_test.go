package dfs

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	//
	//nums := [][]int{
	//	{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	//	{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
	//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	//}

	nums := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	res := maxAreaOfIsland(nums)
	fmt.Println(res)
}

func maxAreaOfIsland(grid [][]int) int {
	//
	maxA := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {

				curA := maxAreaOfIslandHelper(grid, i, j)
				maxA = tools.Max(maxA, curA)
			}
		}
	}

	return maxA
}

// 返回以 i , j 开始的岛屿的面积
func maxAreaOfIslandHelper(grid [][]int, i, j int) int {

	if i >= len(grid) || j >= len(grid[0]) ||
		i < 0 || j < 0 {
		return 0
	}

	if grid[i][j] == 2 {
		return 0
	}

	if grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 2

	return 1 + maxAreaOfIslandHelper(grid, i+1, j) +
		maxAreaOfIslandHelper(grid, i-1, j) +
		maxAreaOfIslandHelper(grid, i, j+1) +
		maxAreaOfIslandHelper(grid, i, j-1)

}
