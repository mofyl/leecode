package v2

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3

	res := search(nums, target)

	fmt.Println(res)
}

func search(nums []int, target int) int {

	if len(nums) < 1 {
		return -1
	}

	// 注意 这里的元素是 没有重复的
	// 这就是一个局部有序的问题
	// 对于这道题 如何判断 哪个区间是有序的。 依据  nums[mid] , nums[l] 的关系
	// 若 nums[l] > nums[mid] 说明  nums[mid] ~ nums[r] 是有序的
	// 若 nums[l] < nums[mid] 说明  nums[l] ~ nums[mid] 是有序的

	l, r := 0, len(nums)-1

	for l <= r {

		mid := l + (r-l)>>1

		if nums[mid] == target {
			return mid
		}

		// 判断哪个区间有序
		if nums[l] <= nums[mid] {

			// 有序后 那么我们的 判断区间就变成了  [l , mid]
			// 表示 在 [l , mid] 区间中， 那么就该收缩 r
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				// 表示不在 [l , mid] 区间中，那么就去找 [mid , r]
				l = mid + 1
			}

		} else {
			// 这里的区间 表示 [mid , r] 上有序

			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		}

	}

	return -1
}
