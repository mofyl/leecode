package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestFindNumberIn2DArray(t *testing.T) {

	//matrix := [][]int{
	//
	//	{1, 4, 7, 11, 15},
	//	{2, 5, 8, 12, 19},
	//	{3, 6, 9, 16, 22},
	//	{10, 13, 14, 17, 24},
	//	{18, 21, 23, 26, 30},
	//}

	matrix := [][]int{}
	//target := 5
	target := 20

	res := findNumberIn2DArray(matrix, target)

	fmt.Println(res)
}

func findNumberIn2DArray(matrix [][]int, target int) bool {

	// 由于 每一行都是递增的。每一列 也是递增的，那么就看 target 是否在 当前行 的范围内，若在范围内则 开始二分。
	// 若不在范围内 则返回 不在

	row := len(matrix)

	if row == 0 {
		return false
	}

	col := len(matrix[0])

	if col == 0 {
		return false
	}

	for i := 0; i < row; i++ {

		if matrix[i][0] <= target && target <= matrix[i][col-1] {
			// 在范围内  则开始二分
			if searchBinary(matrix[i], target) {
				return true
			}
		}

		if matrix[i][0] > target {
			// 说明不在 可以返回了
			return false
		}

	}

	return false

}

func searchBinary(nums []int, target int) bool {

	l := 0
	r := len(nums) - 1

	for l <= r {

		mid := l + (r-l)>>1

		if nums[mid] == target {
			return true
		}

		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return false
}

func findNumberIn2DArray_v2(matrix [][]int, target int) bool {

	// 由于是 行列 都是有序的 所以可以基于 行列来进行查询
	if len(matrix) == 0 {
		return false
	}
	n := len(matrix) - 1
	m := len(matrix[0]) - 1

	i, j := n, 0

	for i >= 0 && j <= m {

		// 当前行的 j位置 比target 大 那么就到上一行去找，
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			// 当前行的 j 位置 比 target 小 那么到当前行的后面去搜索
			j++
		} else {
			return true
		}

	}

	return false

}
