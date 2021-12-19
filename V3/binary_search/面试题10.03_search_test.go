package binary_search

import (
	"fmt"
	"testing"
)

func Test1003Search(t *testing.T) {

	//arr := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	//target := 5
	//target := 11
	arr := []int{5, 5, 5, 1, 2, 3, 4, 5}
	target := 5

	res := search_1003(arr, target)
	fmt.Println(res)
}

func search_1003(arr []int, target int) int {

	l, r := 0, len(arr)-1
	res := -1
	for l <= r {

		mid := l + (r-l)>>1

		if arr[mid] == target {
			res = mid
			// 索引值最小的， 那么就去收缩右边界
			r = mid - 1
		}

		if arr[l] < arr[mid] {
			if arr[l] <= target && target < arr[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if arr[mid] < arr[r] {

			if arr[mid] < target && target <= arr[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		} else {
			for l <= r && arr[mid] == arr[l] {
				l++
			}
			for l <= r && arr[mid] == arr[r] {
				r--
			}
		}

	}

	return res
}
