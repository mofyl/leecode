/*
 * @Author: lirui
 * @Date: 2022-01-10 16:45:01
 * @LastEditTime: 2022-01-10 17:09:52
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\slide_window\v3\713_numSubarrayProductLessThanK_test.go
 */

package v3

import (
	"fmt"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {

	nums := []int{10, 5, 2, 6}
	k := 100
	// 情况1
	// nums := []int{1, 2, 3}
	// k := 0

	res := numSubarrayProductLessThanK(nums, k)

	fmt.Println(res)
}

func numSubarrayProductLessThanK(nums []int, k int) int {

	res := 0
	cur := 1
	l := 0
	r := 0
	for r = 0; r < len(nums); r++ {

		cur *= nums[r]

		// 这里有 等于的原因是情况1，k=0 的时候，这里 其实 l == r 了，但是 在计算res 的时候 会多加1，就会多算，
		for cur >= k && l <= r {
			cur /= nums[l]
			l++
		}
		// 若 窗口中的数组 乘积都 比k小，那么 窗口中的每个数字 一定比 k小
		res += r - l + 1

	}

	return res
}
