package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{3, 1}
	target := 1
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	if nums == nil {
		return -1
	}
	if len(nums) == 0 {
		return -1
	}
	if nums[0] == target {
		return 0
	}

	l := 0
	r := len(nums) - 1
	mid := 0
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[l] {
			// 说明左侧是连续递增的
			// 表示 nums[l] <= target <= nums[mid]
			if nums[l] <= target && target <= nums[mid] {
				// 说明左侧的范围取小了 这时该收缩左侧
				r = mid - 1
			} else {
				// 说明右侧的范围取大了 这时该收缩右侧
				l = mid + 1
			}
		} else {
			//  说明右侧是连续递增的
			// nums[mid] <= target <= nums[right]
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		}
	}
	return -1
}
