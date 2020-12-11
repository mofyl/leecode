package array

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {

	nums := []int{3, 5, 1, 6, 2, 4}
	k := 2

	j := findKthLargest(nums, k)
	fmt.Println(j)
	fmt.Println(nums)
}

func findKthLargest(nums []int, k int) int {

	left := 0
	right := len(nums) - 1
	target := len(nums) - k
	for {
		mid := fastQuick(nums, left, right)
		if mid == target {
			return nums[mid]
		} else if mid < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

func fastQuick(nums []int, left, right int) int {

	mid := nums[left]
	j := left

	for i := left + 1; i <= right; i++ {
		if nums[i] < mid {
			j++
			swap(nums, i, j)
		}
	}

	// 这里一定要做一次交换
	swap(nums, left, j)
	return j
}
