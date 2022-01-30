package v2

import (
	"fmt"
	"testing"
)

func TestSearch_81(t *testing.T) {

	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 3

	//nums := []int{1, 0, 1, 1, 1}
	//target := 0

	res := search_81(nums, target)

	fmt.Println(res)
}

func search_81(nums []int, target int) bool {

	// 这里有 重复的数字
	// 重复的数字 就会使 二分 失去二段性
	// 二段性的意思是  数字 a 要么在  [b , c] 区间上，要么在 [d,e] 区间上，不能又在 [b , c] 上，又在 [d,e] 上
	// 若有了重复的数字 那么就有可能 又在[b , c] 区间上 ，又在[d,e] 上
	// 那么我们 就要当 mid 位置所在元素 和 边界 位置所在元素相同时，将相同的元素跳过

	l, r := 0, len(nums)-1

	for l <= r {

		mid := l + (r-l)>>1

		if nums[mid] == target {
			return true
		}

		// 表示 在  [l , mid] 上有序
		if nums[l] < nums[mid] {

			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}

		} else if nums[l] > nums[mid] {

			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		} else {
			// 这里就出现了 相等的情况， nums[mid] == nums[l] || nums[mid] == nums[r]
			// 这里就要将 相同元素跳过

			for l <= r && nums[mid] == nums[l] {
				l++
			}

			for l <= r && nums[mid] == nums[r] {
				r--
			}
		}

	}

	return false

}
