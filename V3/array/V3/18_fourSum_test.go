/*
 * @Author: lirui
 * @Date: 2021-12-29 15:07:50
 * @LastEditTime: 2021-12-30 12:13:57
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\array\V3\18_fourSum_test.go
 */

package V3

import (
	"fmt"
	"sort"
	"testing"
)

func TestFourSum(t *testing.T) {
	// nums := []int{1, 0, -1, 0, -2, 2}
	// target := 0

	// nums := []int{2, 2, 2, 2, 2}
	// target := 8

	nums := []int{-2, -1, -1, 1, 1, 2, 2}
	target := 0

	res := fourSumBatter(nums, target)

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func fourSum(nums []int, target int) [][]int {

	if len(nums) < 4 {
		return nil
	}

	res := make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 使用滑动窗口

			l := j + 1
			r := len(nums) - 1

			for l < r {

				tmp := nums[i] + nums[j] + nums[l] + nums[r]

				if tmp == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})

					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}

					l++
					r--
				} else if tmp < target {
					l++
				} else {
					r--
				}

			}

		}

	}

	return res

}

func fourSumBatter(nums []int, target int) [][]int {

	if len(nums) < 4 {
		return nil
	}

	res := make([][]int, 0)

	sort.Ints(nums)

	n := len(nums)

	for i := 0; i < n-3; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			// 这里就可以直接return 了
			return res
		}

		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}

		// 不可以有 是错误的分支
		// if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] > target {
		// 	// 无论是什么操作
		// }

		for j := i + 1; j < n-2; j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				// 这里不能直接return. 可能本次j循环 已经循环到了后面 ，也就是到了比较大的数字，导致 > target
				// 但是 可能下一个i，j在从小的开始 就合适了
				// return res
				break
			}

			if nums[i]+nums[j]+nums[n-1]+nums[n-1] < target {
				continue
			}
			// 不可以有 是错误的分支
			// if nums[i]+nums[j]+nums[n-1]+nums[n-1] > target {
			// 	// 无论是什么操作
			// }

			// 使用滑动窗口

			l := j + 1
			r := len(nums) - 1

			for l < r {

				tmp := nums[i] + nums[j] + nums[l] + nums[r]

				if tmp == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})

					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}

					l++
					r--
				} else if tmp < target {
					l++
				} else {
					r--
				}

			}

		}

	}

	return res

}
