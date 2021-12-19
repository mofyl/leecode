package dfs

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {

	nums := [][]byte{

		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	//
	//nums := [][]byte{
	//	{'1', '1', '0', '0', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '1', '0', '0'},
	//	{'0', '0', '0', '1', '1'},
	//}

	res := numIslands(nums)

	fmt.Println(res)
}

/*
	会将 遍历过的 0 和 1 改为2 ，防止重复遍历，
*/
func numIslands(grid [][]byte) int {

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == '1' {
				// 这里表示会出现岛屿 就去搜索整片岛屿，然后将岛屿改为2
				numIslandsHelper(grid, i, j)
				res++
			}

		}
	}

	return res

}

func numIslandsHelper(grid [][]byte, i, j int) {

	if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 || grid[i][j] == '2' {
		return
	}
	// 这里必须是岛屿才会继续遍历
	if grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'

	numIslandsHelper(grid, i-1, j)
	numIslandsHelper(grid, i+1, j)
	numIslandsHelper(grid, i, j-1)
	numIslandsHelper(grid, i, j+1)
}
