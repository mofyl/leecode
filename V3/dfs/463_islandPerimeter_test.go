package dfs

import (
	"fmt"
	"testing"
)

func TestIslandPerimeter(t *testing.T) {
	//
	//nums := [][]int{
	//	{0, 1, 0, 0},
	//	{1, 1, 1, 0},
	//	{0, 1, 0, 0},
	//	{1, 1, 0, 0},
	//}

	nums := [][]int{
		{1, 0},
	}

	res := islandPerimeter(nums)

	fmt.Println(res)
}

func islandPerimeter(grid [][]int) int {

	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == 1 {
				// 这里直接返回 是因为题目中说只有一个岛屿
				return islandPerimeterHelper(grid, i, j)
			}

		}

	}
	return 0
}

func islandPerimeterHelper(grid [][]int, i, j int) int {

	// 这里是边界的部分， 边界也是需要算一条边的
	if i >= len(grid) || j >= len(grid[0]) ||
		i < 0 || j < 0 {
		return 1
	}

	// 这里表示是海洋，海洋是不继续遍历的，因为若海洋格子还继续就会多统计
	// 与数字相邻的海洋格子 要计算一条边进去。
	// 这里只是计算一条就行，因为 若有多个陆地格子共用了该海洋格子，那么每次dfs都会来计算一次
	// 所以是不会少的
	if grid[i][j] == 0 {
		return 1
	}

	if grid[i][j] == 2 {
		return 0
	}

	grid[i][j] = 2

	return islandPerimeterHelper(grid, i+1, j) +
		islandPerimeterHelper(grid, i-1, j) +
		islandPerimeterHelper(grid, i, j+1) +
		islandPerimeterHelper(grid, i, j-1)

}
