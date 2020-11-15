package V1

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	//target := 6
	//nums := []int{2, 2}
	//target := 2
	//res := searchRange(nums, target)
	res := searchRange(nums, target)
	for _, v := range res {
		fmt.Println(v)
	}

}

func searchRange(nums []int, target int) []int {
	left := leftBound(nums, target)
	right := rightBound(nums, target)
	return []int{left, right}
}

func binarySearch(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {

		mid = (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}

	}

	return -1

}

func leftBound(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 收缩 右侧
			right = mid - 1
		}
	}

	// 检查是否越界
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func rightBound(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 收缩左侧
			left = mid + 1
		}

	}

	if right < 0 || nums[right] != target {
		return -1
	}

	return right
}
