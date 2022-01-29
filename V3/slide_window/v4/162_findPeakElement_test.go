package v4

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T) {

	nums := []int{1, 2, 1, 3, 5, 6, 4}

	//nums := []int{1}

	res := findPeakElementV2(nums)

	fmt.Println(res)
}

func findPeakElement(nums []int) int {

	l, r := 0, len(nums)-1

	for l < r {

		mid := l + (r-l)>>1

		if mid-1 >= 0 && mid+1 < len(nums) {

			if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
				return mid
			}

			if nums[mid] < nums[mid-1] {
				l = mid
			} else {
				r = mid
			}
		}

	}

	return -1

}

func findPeakElementV2(nums []int) int {

	l, r := 0, len(nums)-1

	for l < r {

		mid := l + (r-l)>>1

		// 这里表示 沿着递增的方向去二分
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}

	}

	return l

}
