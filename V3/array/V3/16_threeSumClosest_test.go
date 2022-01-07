/*
 * @Author: lirui
 * @Date: 2021-12-29 10:16:43
 * @LastEditTime: 2021-12-29 11:07:40
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\array\V3\16_threeSumClosest_test.go
 */

package V3

import (
	"fmt"
	"leecode/tools"
	"sort"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	// nums := []int{-1, 2, 1, -4}
	// target := 1

	// nums := []int{0, 0, 0}
	// target := 1

	nums := []int{1, 1, 1, 1}
	target := 3

	res := threeSumClosest(nums, target)

	fmt.Println(res)
}

func threeSumClosest(nums []int, target int) int {

	res := -1
	subAbs := 65535
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		if nums[i] > target {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {

			sum := nums[i] + nums[l] + nums[r]

			if sum == target {
				return sum
			} else if sum < target {
				l++
			} else {
				r--
			}
			tmpSub := tools.Abs(int32(target - sum))
			if tmpSub < subAbs {
				res = sum
				subAbs = tmpSub
			}

		}

	}

	return res

}
