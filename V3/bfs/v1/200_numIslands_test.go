/*
 * @Author: lirui38
 * @Date: 2021-12-26 18:00:55
 * @LastEditTime: 2021-12-26 18:27:45
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\bfs\v1\200_numIslands_test.go
 */

package v1

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {

	// grid := [][]byte{

	// 	{'1', '1', '1', '1', '0'},
	// 	{'1', '1', '0', '1', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '0', '0', '0'},
	// }

	grid := [][]byte{

		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	res := numIslands(grid)

	fmt.Println(res)

}

var (
	rowIsLands = 0
	colIsLands = 0
)

func numIslands(grid [][]byte) int {

	rowIsLands = len(grid)

	if rowIsLands < 1 {
		return 0
	}

	colIsLands = len(grid[0])
	res := 0
	for i := 0; i < rowIsLands; i++ {

		for j := 0; j < colIsLands; j++ {

			// 这里只要碰到的陆地 就去进行累加
			if grid[i][j] == '1' {
				// 这里是将该陆地相连的 陆地都改为2 ，防止重复统计
				numIslandsHelper(grid, i, j)
				res++
			}
		}

	}

	return res

}

func numIslandsHelper(grid [][]byte, i, j int) {

	if i < 0 || i >= rowIsLands ||
		j < 0 || j >= colIsLands {
		return
	}

	if grid[i][j] == '0' || grid[i][j] == '2' {
		return
	}

	grid[i][j] = '2'

	numIslandsHelper(grid, i, j-1)
	numIslandsHelper(grid, i, j+1)
	numIslandsHelper(grid, i-1, j)
	numIslandsHelper(grid, i+1, j)

}
