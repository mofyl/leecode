package binary_search

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {

	matrix := [][]int{

		//{1, 3, 5, 7},
		//{10, 11, 16, 20},
		//{23, 30, 34, 60},
		{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
	}

	//fmt.Println(searchMatrix(matrix, 3))
	fmt.Println(searchMatrix(matrix, 13))

}

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) < 1 {
		return false
	}

	for i := 0; i < len(matrix); i++ {

		row := matrix[i]

		if row[0] <= target && row[len(row)-1] >= target {
			// 开始二分查询
			return binarySearch(row, 0, len(row)-1, target)
		}

	}
	return false
}

func binarySearch(data []int, start, end, target int) bool {

	if start > end {
		return false
	}

	mid := (start + end) >> 1

	if data[mid] == target {
		return true
	} else if data[mid] > target {
		return binarySearch(data, start, mid-1, target)
	} else {
		return binarySearch(data, mid+1, end, target)
	}

}
