package v3

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T) {

	//nums := []int{1, 2, 3, 1}
	nums := []int{3, 4, 3, 2, 1}

	res := findPeakElement(nums)

	fmt.Println(res)
}

func findPeakElement(nums []int) int {

	l := 0
	r := len(nums) - 1

	for l < r {

		mid := l + (r-l)>>1

		if mid > 0 && mid < len(nums) {
			if nums[mid-1] < nums[mid] && nums[mid+1] < nums[mid] {
				return mid
			}
			// 这里随便找一个边 进行比较就行 , 可以是 nums[mid-1] 和 nums[mid] 比较
			// 也可以是 nums[mid] 和 nums[mid+1] 进行比较

			// 这里 是要 r = mid 不是 mid-1 ，case: 3,4,3,2,1 这样的情况。
			if nums[mid-1] > nums[mid] {
				r = mid
			} else {
				l = mid
			}
		}

	}

	return -1
}
