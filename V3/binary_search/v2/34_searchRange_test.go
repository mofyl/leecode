package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestSearchRange(t *testing.T) {

	//nums := []int{5, 7, 7, 8, 8, 10}
	//target := 8

	//nums := []int{5, 7, 7, 8, 8, 10}
	//target := 6

	nums := []int{}
	target := 0

	res := searchRange(nums, target)

	fmt.Println(res)
}

func searchRange(nums []int, target int) []int {

	if len(nums) < 1 {
		return []int{-1, -1}
	}

	// 先找左边
	left := findRangLeft(nums, target)

	// 在找右边
	right := findRangeRight(nums, target)

	return []int{left, right}
}

func findRangLeft(nums []int, target int) int {

	l, r := 0, len(nums)-1
	res := len(nums) + 1
	for l <= r {

		mid := l + (r-l)>>1

		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}

		if nums[mid] == target {
			res = tools.Min(res, mid)
		}

	}

	if res == len(nums)+1 {
		return -1
	}

	return res

}

func findRangeRight(nums []int, target int) int {

	l, r := 0, len(nums)-1

	res := -1

	for l <= r {

		mid := l + (r-l)>>1

		if nums[mid] > target {
			r = mid - 1
		} else {

			l = mid + 1

		}

		if nums[mid] == target {
			res = tools.Max(res, mid)
		}

	}

	return res

}
