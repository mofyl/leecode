package v2

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T) {

	nums := []int{1, 2, 3, 1}
	//nums := []int{1, 2, 1, 3, 5, 6, 4}
	//nums := []int{3, 4, 3, 2, 1}

	res := findPeakElement(nums)

	fmt.Println(res)

}

func findPeakElement(nums []int) int {
	// 这里是在寻找峰值！！, 返回峰值元素的下标

	if len(nums) < 1 {
		return 0
	}

	if len(nums) == 1 {
		return 0
	}

	l, r := 0, len(nums)-1

	// 首先对起始位置 和  末尾位置 进行判断，若起始位置和末尾位置 是峰值就直接返回

	if nums[l] > nums[l+1] {
		return l
	}

	if nums[r] > nums[r-1] {
		return r
	}

	// 说明 中间一定有一个峰值

	for l < r {

		mid := l + (r-l)>>1

		if mid-1 >= 0 && mid+1 <= len(nums)-1 {

			if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
				return mid
			}

			if nums[mid] < nums[mid-1] {
				r = mid
			} else {
				l = mid
			}

		} else {
			return 0
		}

	}

	return 0

}
