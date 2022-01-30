package v2

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T) {

	//nums := []int{1, 2, 3, 1}

	//nums := []int{1, 2, 1, 3, 5, 6, 4}

	nums := []int{2, 1}

	res := findPeakElement(nums)

	fmt.Println(res)
}

func findPeakElement(nums []int) int {

	if len(nums) == 1 {
		return 0
	}

	// 第一个元素就是峰值
	if len(nums) > 1 && nums[0] > nums[1] {
		return 0
	}
	// 最后一个元素是峰值
	if len(nums) > 1 && nums[len(nums)-2] < nums[len(nums)-1] {
		return len(nums) - 1
	}

	l, r := 0, len(nums)-1

	for l <= r {

		mid := l + (r-l)>>1

		if mid-1 >= 0 && mid+1 < len(nums) {

			if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
				return mid
			}

			if nums[mid-1] < nums[mid] {
				// 说明 处于 上升的区间
				l = mid + 1
			} else {
				// 说明处于 下降的区间
				r = mid
			}

		} else {
			return -1
		}

	}

	return -1

}
