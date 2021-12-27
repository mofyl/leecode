/*
 * @Author: lirui38
 * @Date: 2021-12-25 18:12:23
 * @LastEditTime: 2021-12-26 14:18:59
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\bfs\v1\1254_closedIsland_test.go
 */

package v1

import (
	"fmt"
	"testing"
)

func TestClosedIsland(t *testing.T) {

	// grid := [][]int{
	// 	{1, 1, 1, 1, 1, 1, 1, 0},
	// 	{1, 0, 0, 0, 0, 1, 1, 0},
	// 	{1, 0, 1, 0, 1, 1, 1, 0},
	// 	{1, 0, 0, 0, 0, 1, 0, 1},
	// 	{1, 1, 1, 1, 1, 1, 1, 0},
	// }

	// res := closedIsland(grid)

	// fmt.Println(res)

	// grid := [][]int{
	// 	{1, 1, 1, 1, 1, 1, 1},
	// 	{1, 0, 0, 0, 0, 0, 1},
	// 	{1, 0, 1, 1, 1, 0, 1},
	// 	{1, 0, 1, 0, 1, 0, 1},
	// 	{1, 0, 1, 1, 1, 0, 1},
	// 	{1, 0, 0, 0, 0, 0, 1},
	// 	{1, 1, 1, 1, 1, 1, 1},
	// }

	// res := closedIsland(grid)

	// fmt.Println(res)

	grid := [][]int{
		{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
	}

	res := closedIsland(grid)

	fmt.Println(res)

}

var (
	row = 0
	col = 0
)

func closedIsland(grid [][]int) int {

	row = len(grid)

	if row < 1 {
		return 0
	}

	col = len(grid[0])

	// 遇到岛屿 就去向四周递归
	res := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			// 表示是岛屿,就去判断是否满足条件
			if grid[i][j] == 0 && closedIslandHelper(grid, i, j) {
				res++
			}

		}
	}

	return res

}

func closedIslandHelper(grid [][]int, i, j int) bool {

	// 这里不能有 等于
	if i < 0 || i >= row ||
		j < 0 || j >= col {
		return false
	}

	// 这里可能是 碰到了 海水 也就是 grid[i][j] == 1
	// 也有可能碰到了一片已经被 遍历过的陆地 grid[i][j] == 2
	if grid[i][j] > 0 {
		return true
	}
	// 若没有被遍历过的陆地，一定会进行递归
	// 标记为已经遍历过 防止下次递归 回头
	grid[i][j] = 2

	ilRes := closedIslandHelper(grid, i, j-1)
	iRRes := closedIslandHelper(grid, i, j+1)
	jURes := closedIslandHelper(grid, i-1, j)
	jDRes := closedIslandHelper(grid, i+1, j)
	// 上面得出结果的步骤一定不能放到 return 中写，因为 与的机制，会漏掉一部分
	return ilRes && iRRes && jURes && jDRes

	// return closedIslandHelper(grid, i, j-1) &&
	// 	closedIslandHelper(grid, i, j+1) &&
	// 	closedIslandHelper(grid, i-1, j) &&
	// 	closedIslandHelper(grid, i+1, j)
}
