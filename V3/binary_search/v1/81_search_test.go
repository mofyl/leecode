package v1

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	//nums := []int{2, 5, 6, 0, 0, 1, 2}
	//target := 0 // 3
	nums := []int{1, 0, 1, 1, 1}
	target := 0 // 3

	res := search(nums, target)
	fmt.Println(res)
}

/*

	这里还是要二分，但是主要是判断 target所在的区间，
		若target是在左半段，那么就去收缩右边，若target在右半段，那么就去收缩左边

	若不满足二分，那么一定出现了 相同的情况，那么收缩边界直到不相同

*/
func search(nums []int, target int) bool {

	l, r := 0, len(nums)-1
	// 这里一定要有等于 不然会漏数字
	for l <= r {

		mid := l + (r-l)>>1

		if nums[mid] == target {
			return true
		}

		if nums[l] < nums[mid] {
			// 判断 target 是否在 [l...mid) 之间 由于上面判断了 nums[mid]和target的关系
			// 所以是左闭右开
			// 这里要有相等是因为 nums中会有重复的数字
			if nums[l] <= target && target < nums[mid] {
				// 说明target 在左边
				// 这里已经判断过 nums[mid] 和target 是不相等的关系了，所以可以跳过mid这个位置
				r = mid - 1
			} else {
				// 同理 可以跳过
				l = mid + 1
			}

		} else if nums[mid] < nums[r] {
			if nums[mid] < target && target <= nums[r] {
				// 同理也可以跳过
				l = mid + 1
			} else {
				// 同理也可以跳过
				r = mid - 1
			}
		} else {
			// 这里表示 nums[l]=nums[mid] || nums[mid] == nums[r]
			// 那么就要跳过相等的部分，继续二分，因为有相等的部分会破坏二段性
			// [2,2,3,4,0,1,2] 这里若target=2 那么 0~3 和 4~5是无法判断出来的 这里就失去了二段性
			for l <= r && nums[l] == nums[mid] {
				l++
			}

			for l <= r && nums[mid] == nums[r] {
				r--
			}
		}

	}
	return false
}
