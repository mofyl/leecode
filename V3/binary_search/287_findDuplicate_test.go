package binary_search

import (
	"fmt"
	"testing"
)

func TestFindDuplicate(t *testing.T) {

	//nums := []int{1, 3, 4, 2, 2}
	//nums := []int{3, 1, 3, 4, 2}
	nums := []int{1, 1, 2}

	res := findDuplicate(nums)
	fmt.Println(res)
}

// 使用二分法
func findDuplicateV2(nums []int) int {

	/*
		由于本题有限制： 所有数字都在 1~n的范围内。若没有这个条件二分就不得行了

		思路：对于一个数组来说，先找到所有数字的中位数mid，
				这里的中位数 指  1~n的的中位数， 不是数组的中位数,找到中位数后
			然后统计比中位数小的数 有多少个 记为cnt，
			若没有重复的 比中位数小的数字 的数量 cnt <= mid (所有数字都在1~n的范围内)
			若 cnt > mid 那么就表示当前区间内 存在重复的数字
	*/

	if len(nums) < 1 {
		return 0
	}

	l, r := 1, len(nums)-1

	for l < r {
		// 找到当前的中位数
		mid := l + (r-l)>>1

		// 开始寻找 比中位数小的数字的个数
		cnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] < mid {
				cnt++
			}
		}

		if cnt < mid {
			// 说明重复的数字不在这里面 移动left
			l = mid
		} else {
			// cnt 比mid大 说明重复数字就在这个区间中，开始收缩r
			r = mid - 1
		}

	}
	return l
}

// 抽屉问题
func findDuplicate(nums []int) int {

	if len(nums) < 1 {
		return 0
	}

	repeat := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {

		if repeat[nums[i]] != 0 {
			return nums[i]
		} else {
			repeat[nums[i]] = nums[i]
		}

	}

	return 0
}
