package v1

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T) {

	//nums := []int{1, 2, 3, 1}
	//nums := []int{1, 2, 1, 3, 5, 6, 4}
	nums := []int{3, 4, 3, 2, 1}
	res := findPeakElement(nums)

	fmt.Println(res)
}

func findPeakElement(nums []int) int {

	if len(nums) < 2 {
		return 0
	}

	// 先对边界进行判断，若边界满足峰值条件则直接返回
	// 这里要返回索引
	if nums[0] > nums[1] {
		return 0
	}

	lenNum := len(nums)
	if nums[lenNum-1] > nums[lenNum-2] {
		return lenNum - 1
	}

	l, r := 0, lenNum-1

	for l < r {

		mid := l + (r-l)>>1

		// mid 在中间了 判断越界的情况
		if mid-1 >= 0 && mid+1 <= lenNum-1 {

			if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
				return mid
			}
			// 这里的mid 不能跳过，对应的就是  3,4,3,2,1这样的情况
			// 第一次mid 指向idx=2的位置 ，然后r收缩，r必须还是=2
			if nums[mid] < nums[mid+1] {
				//l = mid + 1
				l = mid
			} else {
				//r = mid - 1
				r = mid
			}

		} else {
			return -1
		}
	}
	return -1
}
