package v1

import (
	"fmt"
	"leecode/tools"
	"math"
	"testing"
)

func TestSearchRange(t *testing.T) {

	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8

	res := searchRange(nums, target)

	fmt.Println(res)

}

func searchRange(nums []int, target int) []int {

	if len(nums) < 1 {
		return []int{-1, -1}
	}

	// 这里数组是有序的 所以 相同的数字会相邻，这里先找左边界
	// 可以使用二分法
	minL := math.MaxInt64
	searchRangeLeftHelper(nums, 0, len(nums)-1, target, &minL)

	if minL == math.MaxInt64 {
		return []int{-1, -1}
	}

	maxR := math.MinInt64
	searchRangeRightHelper(nums, 0, len(nums)-1, target, &maxR)

	if maxR == math.MinInt64 {
		return []int{-1, -1}
	}

	return []int{minL, maxR}
}

// 返回左侧的位置
func searchRangeRightHelper(nums []int, l, r int, target int, maxR *int) {

	if l > r {
		return
	}

	mid := l + (r-l)>>1

	if nums[mid] < target {
		// 小于收缩 左边界
		searchRangeRightHelper(nums, mid+1, r, target, maxR)
	} else if nums[mid] > target {
		// 大于收缩 右边界
		searchRangeRightHelper(nums, l, mid-1, target, maxR)
	} else {
		// 等于要 移动左边界， 因为要找最大的right
		if nums[mid] == target {
			*maxR = tools.Max(*maxR, mid)
		}

		searchRangeRightHelper(nums, mid+1, r, target, maxR)
	}

}

// 返回左侧的位置
func searchRangeLeftHelper(nums []int, l, r int, target int, minL *int) {

	if l > r {
		return
	}

	mid := l + (r-l)>>1

	if nums[mid] < target {
		searchRangeLeftHelper(nums, mid+1, r, target, minL)
	} else {
		// 大于 和 等于都要收缩 边界

		if nums[mid] == target {
			*minL = tools.Min(*minL, mid)
		}

		searchRangeLeftHelper(nums, l, mid-1, target, minL)
	}

}
